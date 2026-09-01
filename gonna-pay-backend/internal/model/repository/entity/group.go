package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity/enums"
)

type GroupEntity struct {
	ID        uuid.UUID          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID   uuid.UUID          `gorm:"type:uuid;not null;index"`
	Name      string             `gorm:"type:varchar(100);not null"`
	Category  enums.GroupCategory `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Owner   UsersEntity         `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`
	Members []GroupMemberEntity `gorm:"foreignKey:GroupID"`
}
