package converter_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/converter"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity/enums"
	"github.com/stretchr/testify/assert"
)

var (
	validUserID    = uuid.New().String()
	validGroupID   = uuid.New().String()
	validContactID = uuid.New().String()
)

func TestConvertCostDomainToEntity(t *testing.T) {
	t.Run("with GroupID - pointer is non-nil", func(t *testing.T) {
		cost := domains.NewCost(validUserID, validGroupID, "Jantar", "Dinner", 100.0, 50.0)

		e := converter.ConvertCostDomainToEntity(cost)

		assert.NotNil(t, e.GroupID)
		assert.Equal(t, validGroupID, e.GroupID.String())
	})

	t.Run("without GroupID - pointer is nil", func(t *testing.T) {
		cost := domains.NewCost(validUserID, "", "Jantar", "Dinner", 100.0, 100.0)

		e := converter.ConvertCostDomainToEntity(cost)

		assert.Nil(t, e.GroupID)
	})

	t.Run("category is cast to enum", func(t *testing.T) {
		cost := domains.NewCost(validUserID, "", "Jantar", "Dinner", 100.0, 100.0)

		e := converter.ConvertCostDomainToEntity(cost)

		assert.Equal(t, enums.CostCategory("Dinner"), e.Category)
	})
}

func TestConvertContactDomainToEntity(t *testing.T) {
	t.Run("with ID - UUID is parsed", func(t *testing.T) {
		contact := domains.NewContactWithID(validContactID, validUserID, "Ana", "ana@email.com", "11999999999", "Friend")

		e := converter.ConvertContactDomainToEntity(contact)

		assert.Equal(t, validContactID, e.ID.String())
		assert.Equal(t, validUserID, e.OwnerID.String())
	})

	t.Run("without ID - zero UUID", func(t *testing.T) {
		contact := domains.NewContact(validUserID, "Ana", "ana@email.com", "11999999999", "Friend")

		e := converter.ConvertContactDomainToEntity(contact)

		assert.Equal(t, uuid.UUID{}, e.ID)
	})

	t.Run("category is cast to enum", func(t *testing.T) {
		contact := domains.NewContact(validUserID, "Ana", "ana@email.com", "11999999999", "Friend")

		e := converter.ConvertContactDomainToEntity(contact)

		assert.Equal(t, enums.ContactCategory("Friend"), e.Category)
	})
}

func TestConvertGroupDomainToEntity(t *testing.T) {
	t.Run("with ID - UUID is parsed", func(t *testing.T) {
		group := domains.NewGroupWithID(validGroupID, validUserID, "Viagem", "Travel", time.Now(), time.Now(), nil)

		e := converter.ConvertGroupDomainToEntity(group)

		assert.Equal(t, validGroupID, e.ID.String())
		assert.Equal(t, validUserID, e.OwnerID.String())
	})

	t.Run("category is cast to enum", func(t *testing.T) {
		group := domains.NewGroup("", "Viagem", "Travel")

		e := converter.ConvertGroupDomainToEntity(group)

		assert.Equal(t, enums.GroupCategory("Travel"), e.Category)
	})
}

func TestConvertCostEntityToDomain(t *testing.T) {
	t.Run("with GroupID - string is populated", func(t *testing.T) {
		groupID := uuid.New()
		e := entity.CostEntity{
			ID:              uuid.New(),
			UserID:          uuid.New(),
			GroupID:         &groupID,
			CostName:        "Jantar",
			TotalValue:      100.0,
			OwnerPercentage: 50.0,
			Category:        enums.CostCategory("Dinner"),
		}

		cost := converter.ConvertCostEntityToDomain(e)

		assert.Equal(t, groupID.String(), cost.GroupID)
	})

	t.Run("without GroupID - empty string", func(t *testing.T) {
		e := entity.CostEntity{
			ID:              uuid.New(),
			UserID:          uuid.New(),
			GroupID:         nil,
			CostName:        "Jantar",
			TotalValue:      100.0,
			OwnerPercentage: 100.0,
			Category:        enums.CostCategory("Dinner"),
		}

		cost := converter.ConvertCostEntityToDomain(e)

		assert.Empty(t, cost.GroupID)
	})
}

func TestConvertContactEntityToDomain(t *testing.T) {
	t.Run("category enum is cast to string", func(t *testing.T) {
		e := entity.ContactEntity{
			ID:       uuid.New(),
			OwnerID:  uuid.New(),
			Name:     "Ana",
			Email:    "ana@email.com",
			Phone:    "11999999999",
			Category: enums.ContactCategory("Friend"),
		}

		contact := converter.ConvertContactEntityToDomain(e)

		assert.Equal(t, "Friend", contact.Category)
		assert.Equal(t, e.ID.String(), contact.ID)
		assert.Equal(t, e.OwnerID.String(), contact.OwnerID)
	})
}
