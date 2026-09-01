package converter

import (
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

func ConvertEntityToDomain(e entity.UsersEntity) *domains.User {
	return domains.NewUserWithID(
		e.ID.String(),
		e.Name,
		e.Email,
		e.Password,
		e.Phone,
	)
}

func ConvertContactEntityToDomain(e entity.ContactEntity) *domains.Contact {
	return domains.NewContactWithID(
		e.ID.String(),
		e.OwnerID.String(),
		e.Name,
		e.Email,
		e.Phone,
		string(e.Category),
	)
}

func ConvertCostEntityToDomain(e entity.CostEntity) *domains.Cost {
	groupID := ""
	if e.GroupID != nil {
		groupID = e.GroupID.String()
	}

	return domains.NewCostWithID(
		e.ID.String(),
		e.UserID.String(),
		groupID,
		"",
		e.CostName,
		string(e.Category),
		e.TotalValue,
		e.OwnerPercentage,
		e.CreatedAt,
		e.UpdatedAt,
		nil,
	)
}

func ConvertGroupEntityToDomain(e entity.GroupEntity) *domains.Group {
	return domains.NewGroupWithID(
		e.ID.String(),
		e.OwnerID.String(),
		e.Name,
		string(e.Category),
		e.CreatedAt,
		e.UpdatedAt,
		nil,
	)
}
