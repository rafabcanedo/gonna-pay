package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

type AuthRepository interface {
	Save(ctx context.Context, userID string, tokenHash string, expiresAt time.Time) error
	FindByHash(ctx context.Context, tokenHash string) (*entity.RefreshTokenEntity, error)
	DeleteByHash(ctx context.Context, tokenHash string) error
	DeleteByUserID(ctx context.Context, userID string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Save(ctx context.Context, userID string, tokenHash string, expiresAt time.Time) error {
	id := uuid.New()
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO refresh_token_entities (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, $4)`,
		id, userID, tokenHash, expiresAt,
	)
	return err
}

func (r *authRepository) FindByHash(ctx context.Context, tokenHash string) (*entity.RefreshTokenEntity, error) {
	var e entity.RefreshTokenEntity
	err := r.db.QueryRowContext(ctx,
		`SELECT id, user_id, token_hash, expires_at, created_at FROM refresh_token_entities WHERE token_hash = $1`,
		tokenHash,
	).Scan(&e.ID, &e.UserID, &e.TokenHash, &e.ExpiresAt, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *authRepository) DeleteByHash(ctx context.Context, tokenHash string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM refresh_token_entities WHERE token_hash = $1`, tokenHash)
	return err
}

func (r *authRepository) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM refresh_token_entities WHERE user_id = $1`, userID)
	return err
}
