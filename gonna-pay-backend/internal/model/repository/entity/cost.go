package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity/enums"
)

type CostEntity struct {
	ID              uuid.UUID          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          uuid.UUID          `gorm:"type:uuid;not null;index"`
	GroupID         *uuid.UUID         `gorm:"type:uuid;index"`
	CostName        string             `gorm:"type:varchar(100);not null"`
	TotalValue      float64            `gorm:"type:numeric(10,2);not null"`
	OwnerPercentage float64            `gorm:"type:numeric(5,2);not null"`
	Category        enums.CostCategory `gorm:"type:varchar(20);not null"`
	Splits          []byte             `gorm:"type:jsonb"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	User  UsersEntity  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Group *GroupEntity `gorm:"foreignKey:GroupID;constraint:OnDelete:SET NULL;"`
}
