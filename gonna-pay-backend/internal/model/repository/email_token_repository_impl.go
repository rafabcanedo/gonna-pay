package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity/enums"
)

type EmailTokenRepository interface {
	Save(ctx context.Context, userID, tokenHash string, tokenType enums.EmailTokenType, expiresAt time.Time) error
	FindByHashAndType(ctx context.Context, tokenHash string, tokenType enums.EmailTokenType) (*entity.EmailTokenEntity, error)
	DeleteByHash(ctx context.Context, tokenHash string) error
	DeleteByUserIDAndType(ctx context.Context, userID string, tokenType enums.EmailTokenType) error
}

type emailTokenRepository struct {
	db *sql.DB
}

func NewEmailTokenRepository(db *sql.DB) EmailTokenRepository {
	return &emailTokenRepository{db: db}
}

func (r *emailTokenRepository) Save(ctx context.Context, userID, tokenHash string, tokenType enums.EmailTokenType, expiresAt time.Time) error {
	id := uuid.New()
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO email_token_entities (id, user_id, token_hash, type, expires_at) VALUES ($1, $2, $3, $4, $5)`,
		id, userID, tokenHash, tokenType, expiresAt,
	)
	return err
}

func (r *emailTokenRepository) FindByHashAndType(ctx context.Context, tokenHash string, tokenType enums.EmailTokenType) (*entity.EmailTokenEntity, error) {
	var e entity.EmailTokenEntity
	err := r.db.QueryRowContext(ctx,
		`SELECT id, user_id, token_hash, type, expires_at, created_at FROM email_token_entities WHERE token_hash = $1 AND type = $2`,
		tokenHash, tokenType,
	).Scan(&e.ID, &e.UserID, &e.TokenHash, &e.Type, &e.ExpiresAt, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *emailTokenRepository) DeleteByHash(ctx context.Context, tokenHash string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM email_token_entities WHERE token_hash = $1`, tokenHash)
	return err
}

func (r *emailTokenRepository) DeleteByUserIDAndType(ctx context.Context, userID string, tokenType enums.EmailTokenType) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM email_token_entities WHERE user_id = $1 AND type = $2`, userID, tokenType)
	return err
}
