package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/converter"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

type GroupRepository interface {
	Create(ctx context.Context, group *domains.Group, memberIDs []string) (*domains.Group, error)
	FindAll(ctx context.Context, ownerID string, limit, offset int, filters domains.GroupFilters) ([]*domains.Group, int64, error)
	FindByID(ctx context.Context, id string) (*domains.Group, error)
	Update(ctx context.Context, group *domains.Group) (*domains.Group, error)
	Delete(ctx context.Context, id string) error
	AddMember(ctx context.Context, groupID, contactID string) error
	RemoveMember(ctx context.Context, groupID, contactID string) error
	IsContactOwnedBy(ctx context.Context, contactID, ownerID string) (bool, error)
	MemberExists(ctx context.Context, groupID, contactID string) (bool, error)
	GetMembers(ctx context.Context, groupID string) ([]domains.Member, error)
}

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(ctx context.Context, group *domains.Group, memberIDs []string) (*domains.Group, error) {
	e := converter.ConvertGroupDomainToEntity(group)
	e.ID = uuid.New()
	now := time.Now()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`INSERT INTO group_entities (id, owner_id, name, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		e.ID, e.OwnerID, e.Name, e.Category, now, now,
	)
	if err != nil {
		return nil, err
	}

	for _, contactID := range memberIDs {
		_, err = tx.ExecContext(ctx,
			`INSERT INTO group_member_entities (id, group_id, contact_id) VALUES ($1, $2, $3)`,
			uuid.New(), e.ID, contactID,
		)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return domains.NewGroupWithID(e.ID.String(), group.OwnerID, group.Name, group.Category, now, now, nil), nil
}

func (r *groupRepository) FindAll(ctx context.Context, ownerID string, limit, offset int, filters domains.GroupFilters) ([]*domains.Group, int64, error) {
	conditions := []string{"owner_id = $1"}
	args := []any{ownerID}
	idx := 2

	if filters.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", idx))
		args = append(args, filters.Category)
		idx++
	}
	if filters.Search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", idx))
		args = append(args, "%"+filters.Search+"%")
		idx++
	}

	where := strings.Join(conditions, " AND ")

	var total int64
	err := r.db.QueryRowContext(ctx,
		fmt.Sprintf(`SELECT COUNT(*) FROM group_entities WHERE %s`, where),
		args...,
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	queryArgs := append(args, limit, offset)
	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf(`
			SELECT id, owner_id, name, category, created_at, updated_at
			FROM group_entities
			WHERE %s
			ORDER BY created_at DESC
			LIMIT $%d OFFSET $%d
		`, where, idx, idx+1),
		queryArgs...,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var groups []*domains.Group
	for rows.Next() {
		var e entity.GroupEntity
		if err := rows.Scan(&e.ID, &e.OwnerID, &e.Name, &e.Category, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, 0, err
		}
		groups = append(groups, converter.ConvertGroupEntityToDomain(e))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return groups, total, nil
}

func (r *groupRepository) FindByID(ctx context.Context, id string) (*domains.Group, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			g.id, g.owner_id, g.name, g.category, g.created_at, g.updated_at,
			c.id, c.name, c.email
		FROM group_entities g
		LEFT JOIN group_member_entities gm ON gm.group_id = g.id
		LEFT JOIN contact_entities c ON c.id = gm.contact_id AND c.deleted_at IS NULL
		WHERE g.id = $1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		groupID   uuid.UUID
		ownerID   uuid.UUID
		name      string
		category  string
		createdAt time.Time
		updatedAt time.Time
		members   []domains.Member
		found     bool
	)

	for rows.Next() {
		var (
			memberID    sql.NullString
			memberName  sql.NullString
			memberEmail sql.NullString
		)

		if err := rows.Scan(
			&groupID, &ownerID, &name, &category, &createdAt, &updatedAt,
			&memberID, &memberName, &memberEmail,
		); err != nil {
			return nil, err
		}

		found = true

		if memberID.Valid {
			members = append(members, domains.Member{
				ID:    memberID.String,
				Name:  memberName.String,
				Email: memberEmail.String,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if !found {
		return nil, domains.NewNotFoundError("group not found")
	}

	return domains.NewGroupWithID(groupID.String(), ownerID.String(), name, category, createdAt, updatedAt, members), nil
}

func (r *groupRepository) Update(ctx context.Context, group *domains.Group) (*domains.Group, error) {
	e := converter.ConvertGroupDomainToEntity(group)
	now := time.Now()

	_, err := r.db.ExecContext(ctx,
		`UPDATE group_entities SET name = $1, category = $2, updated_at = $3 WHERE id = $4`,
		e.Name, e.Category, now, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return domains.NewGroupWithID(group.ID, group.OwnerID, group.Name, group.Category, group.CreatedAt, now, group.Members), nil
}

func (r *groupRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`UPDATE cost_entities SET owner_percentage = 100, splits = NULL WHERE group_id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM group_entities WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *groupRepository) AddMember(ctx context.Context, groupID, contactID string) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO group_member_entities (id, group_id, contact_id) VALUES ($1, $2, $3)`,
		uuid.New(), groupID, contactID,
	)
	return err
}

func (r *groupRepository) RemoveMember(ctx context.Context, groupID, contactID string) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM group_member_entities WHERE group_id = $1 AND contact_id = $2`,
		groupID, contactID,
	)
	return err
}

func (r *groupRepository) IsContactOwnedBy(ctx context.Context, contactID, ownerID string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM contact_entities WHERE id = $1 AND owner_id = $2 AND deleted_at IS NULL`,
		contactID, ownerID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *groupRepository) MemberExists(ctx context.Context, groupID, contactID string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM group_member_entities WHERE group_id = $1 AND contact_id = $2`,
		groupID, contactID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *groupRepository) GetMembers(ctx context.Context, groupID string) ([]domains.Member, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT gm.contact_id, c.name
		FROM group_member_entities gm
		JOIN contact_entities c ON c.id = gm.contact_id AND c.deleted_at IS NULL
		WHERE gm.group_id = $1
	`, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []domains.Member
	for rows.Next() {
		var m domains.Member
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
