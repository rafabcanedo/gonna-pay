package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/converter"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

type ContactRepository interface {
	Create(ctx context.Context, contact *domains.Contact) (*domains.Contact, error)
	FindAll(ctx context.Context, ownerID string, limit, offset int, filters domains.ContactFilters) ([]*domains.Contact, int64, error)
	FindByID(ctx context.Context, id string) (*domains.Contact, error)
	Update(ctx context.Context, contact *domains.Contact) (*domains.Contact, error)
	Delete(ctx context.Context, id string) error
	ExistsByEmailAndOwner(ctx context.Context, email, ownerID string) (bool, error)
	FindContactsByFrequency(ctx context.Context, userID string, limit int) ([]domains.ContactFrequency, error)
	FindStats(ctx context.Context, ownerID string) (*domains.ContactStats, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(ctx context.Context, contact *domains.Contact) (*domains.Contact, error) {
	e := converter.ConvertContactDomainToEntity(contact)
	e.ID = uuid.New()
	now := time.Now()

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO contact_entities (id, owner_id, name, email, phone, category, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		e.ID, e.OwnerID, e.Name, e.Email, e.Phone, e.Category, now,
	)
	if err != nil {
		return nil, err
	}

	e.CreatedAt = now
	return converter.ConvertContactEntityToDomain(*e), nil
}

func (r *contactRepository) FindAll(ctx context.Context, ownerID string, limit, offset int, filters domains.ContactFilters) ([]*domains.Contact, int64, error) {
	conditions := []string{"owner_id = $1", "deleted_at IS NULL"}
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
		fmt.Sprintf(`SELECT COUNT(*) FROM contact_entities WHERE %s`, where),
		args...,
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	queryArgs := append(args, limit, offset)
	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf(`
			SELECT id, owner_id, name, email, phone, category, created_at
			FROM contact_entities
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

	var contacts []*domains.Contact
	for rows.Next() {
		var e entity.ContactEntity
		if err := rows.Scan(&e.ID, &e.OwnerID, &e.Name, &e.Email, &e.Phone, &e.Category, &e.CreatedAt); err != nil {
			return nil, 0, err
		}
		contacts = append(contacts, converter.ConvertContactEntityToDomain(e))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}

func (r *contactRepository) FindByID(ctx context.Context, id string) (*domains.Contact, error) {
	var e entity.ContactEntity

	err := r.db.QueryRowContext(ctx,
		`SELECT id, owner_id, name, email, phone, category, created_at FROM contact_entities WHERE id = $1 AND deleted_at IS NULL`,
		id,
	).Scan(&e.ID, &e.OwnerID, &e.Name, &e.Email, &e.Phone, &e.Category, &e.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("contact not found")
		}
		return nil, err
	}

	return converter.ConvertContactEntityToDomain(e), nil
}

func (r *contactRepository) Update(ctx context.Context, contact *domains.Contact) (*domains.Contact, error) {
	e := converter.ConvertContactDomainToEntity(contact)

	_, err := r.db.ExecContext(ctx,
		`UPDATE contact_entities SET name = $1, email = $2, phone = $3, category = $4 WHERE id = $5`,
		e.Name, e.Email, e.Phone, e.Category, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertContactEntityToDomain(*e), nil
}

func (r *contactRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE contact_entities SET deleted_at = $1 WHERE id = $2`,
		time.Now(), id,
	)
	return err
}

func (r *contactRepository) ExistsByEmailAndOwner(ctx context.Context, email, ownerID string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM contact_entities WHERE email = $1 AND owner_id = $2 AND deleted_at IS NULL`,
		email, ownerID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *contactRepository) FindStats(ctx context.Context, ownerID string) (*domains.ContactStats, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT category, COUNT(*) FROM contact_entities WHERE owner_id = $1 AND deleted_at IS NULL GROUP BY category`,
		ownerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	byCategory := make(map[string]int64)
	for rows.Next() {
		var category string
		var count int64
		if err := rows.Scan(&category, &count); err != nil {
			return nil, err
		}
		byCategory[category] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &domains.ContactStats{ByCategory: byCategory}, nil
}

func (r *contactRepository) FindContactsByFrequency(ctx context.Context, userID string, limit int) ([]domains.ContactFrequency, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT elem->>'contactId', elem->>'contactName', COUNT(*)
		FROM cost_entities c, jsonb_array_elements(c.splits) AS elem
		WHERE c.user_id = $1
		GROUP BY elem->>'contactId', elem->>'contactName'
		ORDER BY COUNT(*) DESC
		LIMIT $2`,
		userID, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domains.ContactFrequency
	for rows.Next() {
		var t domains.ContactFrequency
		if err := rows.Scan(&t.ContactID, &t.ContactName, &t.SharedCosts); err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
