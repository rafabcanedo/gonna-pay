package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/converter"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository/entity"
)

type ContactRepository interface {
	Create(contact *domains.Contact) (*domains.Contact, error)
	FindAll(ownerID string) ([]*domains.Contact, error)
	FindByID(id string) (*domains.Contact, error)
	Update(contact *domains.Contact) (*domains.Contact, error)
	Delete(id string) error
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

	_, err := r.db.Exec(
		`INSERT INTO contact_entities (id, owner_id, name, email, phone, category) VALUES ($1, $2, $3, $4, $5, $6)`,
		e.ID, e.OwnerID, e.Name, e.Email, e.Phone, e.Category,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertContactEntityToDomain(*e), nil
}

func (r *contactRepository) FindAll(ownerID string) ([]*domains.Contact, error) {
	rows, err := r.db.Query(
		`SELECT id, owner_id, name, email, phone, category FROM contact_entities WHERE owner_id = $1`,
		ownerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*domains.Contact
	for rows.Next() {
		var e entity.ContactEntity
		if err := rows.Scan(&e.ID, &e.OwnerID, &e.Name, &e.Email, &e.Phone, &e.Category); err != nil {
			return nil, err
		}
		contacts = append(contacts, converter.ConvertContactEntityToDomain(e))
	}

	return contacts, nil
}

func (r *contactRepository) FindByID(id string) (*domains.Contact, error) {
	var e entity.ContactEntity

	row := r.db.QueryRow(
		`SELECT id, owner_id, name, email, phone, category FROM contact_entities WHERE id = $1`,
		id,
	)

	err := row.Scan(&e.ID, &e.OwnerID, &e.Name, &e.Email, &e.Phone, &e.Category)
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
	_, err := r.db.Exec(`DELETE FROM contact_entities WHERE id = $1`, id)
	return err
}
