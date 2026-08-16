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
	Create(user *domains.User) (*domains.User, error)
	FindAll() ([]*domains.User, error)
	FindByID(id string) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
	Update(user *domains.User) (*domains.User, error)
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domains.User) (*domains.User, error) {
	e := converter.ConvertDomainToEntity(user)
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

func (r *userRepository) FindAll() ([]*domains.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email, password, phone FROM users_entities`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domains.User
	for rows.Next() {
		var e entity.UsersEntity
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone); err != nil {
			return nil, err
		}
		users = append(users, converter.ConvertEntityToDomain(e))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(id string) (*domains.User, error) {
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

func (r *userRepository) FindByEmail(email string) (*domains.User, error) {
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

func (r *userRepository) Update(user *domains.User) (*domains.User, error) {
	e := converter.ConvertDomainToEntity(user)

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
