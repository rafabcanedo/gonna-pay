package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/converter"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

type ContactRepository interface {
	Create(contact *domains.Contact) (*domains.Contact, error)
	FindAll(ownerID string, limit, offset int) ([]*domains.Contact, int64, error)
	FindByID(id string) (*domains.Contact, error)
	Update(contact *domains.Contact) (*domains.Contact, error)
	Delete(id string) error
	ExistsByEmailAndOwner(email, ownerID string) (bool, error)
	FindContactsByFrequency(userID string, limit int) ([]domains.ContactFrequency, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(contact *domains.Contact) (*domains.Contact, error) {
	e := converter.ConvertContactDomainToEntity(contact)
	e.ID = uuid.New()
	now := time.Now()

	_, err := r.db.Exec(
		`INSERT INTO contact_entities (id, owner_id, name, email, phone, category, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		e.ID, e.OwnerID, e.Name, e.Email, e.Phone, e.Category, now,
	)
	if err != nil {
		return nil, err
	}

	e.CreatedAt = now
	return converter.ConvertContactEntityToDomain(*e), nil
}

func (r *contactRepository) FindAll(ownerID string, limit, offset int) ([]*domains.Contact, int64, error) {
	var total int64

	err := r.db.QueryRow(`SELECT COUNT(*) FROM contact_entities WHERE owner_id = $1 AND deleted_at IS NULL`, ownerID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(
		`SELECT id, owner_id, name, email, phone, category, created_at
		 FROM contact_entities
		 WHERE owner_id = $1 AND deleted_at IS NULL
		 ORDER BY created_at DESC
		 LIMIT $2 OFFSET $3`,
		ownerID, limit, offset,
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

func (r *contactRepository) FindByID(id string) (*domains.Contact, error) {
	var e entity.ContactEntity

	row := r.db.QueryRow(
		`SELECT id, owner_id, name, email, phone, category, created_at FROM contact_entities WHERE id = $1 AND deleted_at IS NULL`,
		id,
	)

	err := row.Scan(&e.ID, &e.OwnerID, &e.Name, &e.Email, &e.Phone, &e.Category, &e.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("contact not found")
		}
		return nil, err
	}

	return converter.ConvertContactEntityToDomain(e), nil
}

func (r *contactRepository) Update(contact *domains.Contact) (*domains.Contact, error) {
	e := converter.ConvertContactDomainToEntity(contact)

	_, err := r.db.Exec(
		`UPDATE contact_entities SET name = $1, email = $2, phone = $3, category = $4 WHERE id = $5`,
		e.Name, e.Email, e.Phone, e.Category, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertContactEntityToDomain(*e), nil
}

func (r *contactRepository) Delete(id string) error {
	_, err := r.db.Exec(
		`UPDATE contact_entities SET deleted_at = $1 WHERE id = $2`,
		time.Now(), id,
	)
	return err
}

func (r *contactRepository) ExistsByEmailAndOwner(email, ownerID string) (bool, error) {
	var count int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM contact_entities WHERE email = $1 AND owner_id = $2 AND deleted_at IS NULL`,
		email, ownerID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *contactRepository) FindContactsByFrequency(userID string, limit int) ([]domains.ContactFrequency, error) {
	rows, err := r.db.Query(
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
