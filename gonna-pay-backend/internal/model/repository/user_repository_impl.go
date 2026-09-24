package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/converter"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *domains.User) (*domains.User, error)
	FindAll(ctx context.Context, limit, offset int) ([]*domains.User, int64, error)
	FindByID(ctx context.Context, id string) (*domains.User, error)
	FindByEmail(ctx context.Context, email string) (*domains.User, error)
	Update(ctx context.Context, user *domains.User) (*domains.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domains.User) (*domains.User, error) {
	e := converter.ConvertDomainToEntity(user)
	e.ID = uuid.New()

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users_entities (id, name, email, password, phone, email_verified) VALUES ($1, $2, $3, $4, $5, $6)`,
		e.ID, e.Name, e.Email, e.Password, e.Phone, e.EmailVerified,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertEntityToDomain(*e), nil
}

func (r *userRepository) FindAll(ctx context.Context, limit, offset int) ([]*domains.User, int64, error) {
	var total int64

	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users_entities`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(ctx, `SELECT id, name, email, password, phone, email_verified FROM users_entities LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*domains.User
	for rows.Next() {
		var e entity.UsersEntity
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone, &e.EmailVerified); err != nil {
			return nil, 0, err
		}
		users = append(users, converter.ConvertEntityToDomain(e))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*domains.User, error) {
	var e entity.UsersEntity

	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, email, password, phone, email_verified FROM users_entities WHERE id = $1`,
		id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone, &e.EmailVerified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("user not found")
		}
		return nil, err
	}

	return converter.ConvertEntityToDomain(e), nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domains.User, error) {
	var e entity.UsersEntity

	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, email, password, phone, email_verified FROM users_entities WHERE email = $1`,
		email,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Phone, &e.EmailVerified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domains.NewNotFoundError("user not found")
		}
		return nil, err
	}

	return converter.ConvertEntityToDomain(e), nil
}

func (r *userRepository) Update(ctx context.Context, user *domains.User) (*domains.User, error) {
	e := converter.ConvertDomainToEntity(user)

	_, err := r.db.ExecContext(ctx,
		`UPDATE users_entities SET name = $1, email = $2, password = $3, phone = $4, email_verified = $5 WHERE id = $6`,
		e.Name, e.Email, e.Password, e.Phone, e.EmailVerified, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return converter.ConvertEntityToDomain(*e), nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM users_entities WHERE id = $1`, id)
	return err
}
