package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/converter"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

type GroupRepository interface {
	Create(group *domains.Group, memberIDs []string) (*domains.Group, error)
	FindAll(ownerID string, limit, offset int) ([]*domains.Group, int64, error)
	FindByID(id string) (*domains.Group, error)
	Update(group *domains.Group) (*domains.Group, error)
	Delete(id string) error
	AddMember(groupID, contactID string) error
	RemoveMember(groupID, contactID string) error
	IsContactOwnedBy(contactID, ownerID string) (bool, error)
	MemberExists(groupID, contactID string) (bool, error)
}

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(group *domains.Group, memberIDs []string) (*domains.Group, error) {
	e := converter.ConvertGroupDomainToEntity(group)
	e.ID = uuid.New()
	now := time.Now()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`INSERT INTO group_entities (id, owner_id, name, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		e.ID, e.OwnerID, e.Name, e.Category, now, now,
	)
	if err != nil {
		return nil, err
	}

	for _, contactID := range memberIDs {
		_, err = tx.Exec(
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

func (r *groupRepository) FindAll(ownerID string, limit, offset int) ([]*domains.Group, int64, error) {
	var total int64

	err := r.db.QueryRow(`SELECT COUNT(*) FROM group_entities WHERE owner_id = $1`, ownerID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(
		`SELECT id, owner_id, name, category, created_at, updated_at FROM group_entities WHERE owner_id = $1 LIMIT $2 OFFSET $3`,
		ownerID, limit, offset,
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

func (r *groupRepository) FindByID(id string) (*domains.Group, error) {
	rows, err := r.db.Query(`
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

func (r *groupRepository) Update(group *domains.Group) (*domains.Group, error) {
	e := converter.ConvertGroupDomainToEntity(group)
	now := time.Now()

	_, err := r.db.Exec(
		`UPDATE group_entities SET name = $1, category = $2, updated_at = $3 WHERE id = $4`,
		e.Name, e.Category, now, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return domains.NewGroupWithID(group.ID, group.OwnerID, group.Name, group.Category, group.CreatedAt, now, group.Members), nil
}

func (r *groupRepository) Delete(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`UPDATE cost_entities SET owner_percentage = 100, splits = NULL WHERE group_id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM group_entities WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *groupRepository) AddMember(groupID, contactID string) error {
	_, err := r.db.Exec(
		`INSERT INTO group_member_entities (id, group_id, contact_id) VALUES ($1, $2, $3)`,
		uuid.New(), groupID, contactID,
	)
	return err
}

func (r *groupRepository) RemoveMember(groupID, contactID string) error {
	_, err := r.db.Exec(
		`DELETE FROM group_member_entities WHERE group_id = $1 AND contact_id = $2`,
		groupID, contactID,
	)
	return err
}

func (r *groupRepository) IsContactOwnedBy(contactID, ownerID string) (bool, error) {
	var count int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM contact_entities WHERE id = $1 AND owner_id = $2 AND deleted_at IS NULL`,
		contactID, ownerID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *groupRepository) MemberExists(groupID, contactID string) (bool, error) {
	var count int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM group_member_entities WHERE group_id = $1 AND contact_id = $2`,
		groupID, contactID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
