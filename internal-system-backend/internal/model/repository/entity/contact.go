package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository/entity/enums"
)

type ContactEntity struct {
	ID        uuid.UUID             `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID   uuid.UUID             `gorm:"type:uuid;not null"`
	Name      string                `gorm:"type:varchar(50);not null"`
	Email     string                `gorm:"type:varchar(50)"`
	Phone     string                `gorm:"type:varchar(20)"`
	Category  enums.ContactCategory `gorm:"type:varchar(20);not null"`
	DeletedAt *time.Time            `gorm:"index"`

	Owner UsersEntity `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`
}
