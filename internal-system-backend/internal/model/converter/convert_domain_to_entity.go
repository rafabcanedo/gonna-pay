package converter

import (
	"github.com/google/uuid"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository/entity"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository/entity/enums"
)

func ConvertDomainToEntity(user *domains.User) *entity.UsersEntity {
	var userID uuid.UUID
	if user.ID != "" {
		userID = uuid.MustParse(user.ID)
	}

	return &entity.UsersEntity{
		ID:       userID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
}

func ConvertContactDomainToEntity(contact *domains.Contact) *entity.ContactEntity {
	var contactID, ownerUUID uuid.UUID
	if contact.ID != "" {
		contactID = uuid.MustParse(contact.ID)
	}
	if contact.OwnerID != "" {
		ownerUUID = uuid.MustParse(contact.OwnerID)
	}

	return &entity.ContactEntity{
		ID:       contactID,
		OwnerID:  ownerUUID,
		Name:     contact.Name,
		Email:    contact.Email,
		Phone:    contact.Phone,
		Category: enums.ContactCategory(contact.Category),
	}
}

func ConvertCostDomainToEntity(cost *domains.Cost) *entity.CostEntity {
	var costID, userUUID uuid.UUID
	if cost.ID != "" {
		costID = uuid.MustParse(cost.ID)
	}
	if cost.UserID != "" {
		userUUID = uuid.MustParse(cost.UserID)
	}

	var groupUUID *uuid.UUID
	if cost.GroupID != "" {
		parsed := uuid.MustParse(cost.GroupID)
		groupUUID = &parsed
	}

	return &entity.CostEntity{
		ID:              costID,
		UserID:          userUUID,
		GroupID:         groupUUID,
		CostName:        cost.CostName,
		TotalValue:      cost.TotalValue,
		OwnerPercentage: cost.OwnerPercentage,
		Category:        enums.CostCategory(cost.Category),
	}
}

func ConvertGroupDomainToEntity(group *domains.Group) *entity.GroupEntity {
	var groupID, ownerUUID uuid.UUID
	if group.ID != "" {
		groupID = uuid.MustParse(group.ID)
	}
	if group.OwnerID != "" {
		ownerUUID = uuid.MustParse(group.OwnerID)
	}

	return &entity.GroupEntity{
		ID:       groupID,
		OwnerID:  ownerUUID,
		Name:     group.Name,
		Category: enums.GroupCategory(group.Category),
	}
}
