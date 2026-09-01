package entity

import "github.com/google/uuid"

type UsersEntity struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"type:varchar(50);not null"`
	Email string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"type:varchar(60);not null"`
	Phone string `gorm:"type:varchar(20);not null"`

	Contacts []ContactEntity `gorm:"foreignKey:OwnerID"`
	Costs    []CostEntity    `gorm:"foreignKey:UserID"`
}
