package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity/enums"
)

type EmailTokenEntity struct {
	ID        uuid.UUID            `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID            `gorm:"type:uuid;not null;index"`
	TokenHash string               `gorm:"type:varchar(64);not null;uniqueIndex"`
	Type      enums.EmailTokenType `gorm:"type:varchar(30);not null"`
	ExpiresAt time.Time            `gorm:"not null"`
	CreatedAt time.Time            `gorm:"autoCreateTime"`
}
