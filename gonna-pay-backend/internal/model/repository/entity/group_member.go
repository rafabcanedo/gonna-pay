package entity

import "github.com/google/uuid"

type GroupMemberEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GroupID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_group_contact"`
	ContactID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_group_contact"`

	Group   GroupEntity   `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
	Contact ContactEntity `gorm:"foreignKey:ContactID;constraint:OnDelete:CASCADE;"`
}
