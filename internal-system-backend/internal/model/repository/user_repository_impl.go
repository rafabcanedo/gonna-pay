package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/converter"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository/entity"
)

type UserRepository interface {
	Create(domain domains.UserDomainInterface) (domains.UserDomainInterface, error)
	FindAll() ([]domains.UserDomainInterface, error)
	FindByID(id string) (domains.UserDomainInterface, error)
	FindByEmail(email string) (domains.UserDomainInterface, error)
	Update(domain domains.UserDomainInterface) (domains.UserDomainInterface, error)
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(domain domains.UserDomainInterface) (domains.UserDomainInterface, error) {
	e := converter.ConvertDomainToEntity(domain)
	e.ID = uuid.New()

	_, err := r.db.Exec(
		`INSERT INTO users_entities (id, name, email, password, phone) VALUES ($1, $2, $3, $4, $5)`,
		e.ID, e.Name, e.Email, e.Password, e.Phone,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertEntityToDomain(*e), nil
}

func (r *userRepository) FindAll() ([]domains.UserDomainInterface, error) {
	rows, err := r.db.Query(`SELECT id, name, email, password, phone FROM users_entities`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domains.UserDomainInterface
	for rows.Next() {
		var e entity.UsersEntity
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone); err != nil {
			return nil, err
		}
		users = append(users, converter.ConvertEntityToDomain(e))
	}

	return users, nil
}

func (r *userRepository) FindByID(id string) (domains.UserDomainInterface, error) {
	var e entity.UsersEntity

	row := r.db.QueryRow(
		`SELECT id, name, email, password, phone FROM users_entities WHERE id = $1`,
		id,
	)

	err := row.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("user not found")
		}
		return nil, err
	}

	return converter.ConvertEntityToDomain(e), nil
}

func (r *userRepository) FindByEmail(email string) (domains.UserDomainInterface, error) {
	var e entity.UsersEntity

	row := r.db.QueryRow(
		`SELECT id, name, email, password, phone FROM users_entities WHERE email = $1`,
		email,
	)

	err := row.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("user not found")
		}
		return nil, err
	}

	return converter.ConvertEntityToDomain(e), nil
}

func (r *userRepository) Update(domain domains.UserDomainInterface) (domains.UserDomainInterface, error) {
	e := converter.ConvertDomainToEntity(domain)

	_, err := r.db.Exec(
		`UPDATE users_entities SET name = $1, email = $2, password = $3, phone = $4 WHERE id = $5`,
		e.Name, e.Email, e.Password, e.Phone, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertEntityToDomain(*e), nil
}

func (r *userRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM users_entities WHERE id = $1`, id)
	return err
}
