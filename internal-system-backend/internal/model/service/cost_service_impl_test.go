package service_test

import (
	"errors"
	"testing"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/mocks"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/service"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func float64Ptr(v float64) *float64 { return &v }

func TestCostService_Create(t *testing.T) {
	t.Run("success without group - owner gets 100%", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		cost := domains.NewCost("user-1", "", "Jantar", "Dinner", 100.0, 0)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Nil()).Return(testutil.NewCostFixture(), nil)

		result, err := svc.Create(cost, nil)

		assert.NoError(t, err)
		assert.Equal(t, "cost-1", result.ID)
		assert.Equal(t, float64(100), cost.OwnerPercentage)
	})

	t.Run("success with group and explicit ownerPercentage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 100.0, 0)
		members := []domains.Member{{ID: "contact-1", Name: "Ana"}, {ID: "contact-2", Name: "Pedro"}}

		mockRepo.EXPECT().GetGroupMembers("group-1").Return(members, nil)
		mockRepo.EXPECT().Create(gomock.Any(), members).Return(testutil.NewCostFixture(), nil)

		result, err := svc.Create(cost, float64Ptr(40.0))

		assert.NoError(t, err)
		assert.Equal(t, "cost-1", result.ID)
		assert.Equal(t, 40.0, cost.OwnerPercentage)
	})

	t.Run("success with group and auto ownerPercentage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 100.0, 0)
		members := []domains.Member{{ID: "contact-1", Name: "Ana"}, {ID: "contact-2", Name: "Pedro"}, {ID: "contact-3", Name: "João"}}

		mockRepo.EXPECT().GetGroupMembers("group-1").Return(members, nil)
		mockRepo.EXPECT().Create(gomock.Any(), members).Return(testutil.NewCostFixture(), nil)

		_, err := svc.Create(cost, nil)

		assert.NoError(t, err)
		assert.Equal(t, 25.0, cost.OwnerPercentage)
	})

	t.Run("invalid ownerPercentage - zero", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 100.0, 0)
		mockRepo.EXPECT().GetGroupMembers("group-1").Return([]domains.Member{{ID: "contact-1", Name: "Ana"}}, nil)

		_, err := svc.Create(cost, float64Ptr(0))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("invalid ownerPercentage - 100", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 100.0, 0)
		mockRepo.EXPECT().GetGroupMembers("group-1").Return([]domains.Member{{ID: "contact-1", Name: "Ana"}}, nil)

		_, err := svc.Create(cost, float64Ptr(100))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("repo error on GetGroupMembers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 100.0, 0)
		mockRepo.EXPECT().GetGroupMembers("group-1").Return(nil, errors.New("db error"))

		_, err := svc.Create(cost, nil)
		assert.Error(t, err)
	})
}

func TestCostService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		c1 := testutil.NewCostFixture()
		c2 := testutil.NewCostFixture()
		c2.ID = "cost-2"
		mockRepo.EXPECT().FindAll("user-1").Return([]*domains.Cost{c1, c2}, nil)

		result, err := svc.FindAll("user-1")

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})
}

func TestCostService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)

		result, err := svc.FindByID("cost-1", "user-1")

		assert.NoError(t, err)
		assert.Equal(t, "cost-1", result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(nil, domains.NewNotFoundError("cost not found"))

		_, err := svc.FindByID("cost-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)

		_, err := svc.FindByID("cost-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestCostService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		existing := testutil.NewCostFixture()
		incoming := domains.NewCost("user-1", "", "Jantar Atualizado", "Dinner", 150.0, 0)

		mockRepo.EXPECT().FindByID("cost-1").Return(existing, nil)
		mockRepo.EXPECT().Update("cost-1", gomock.Any()).Return(testutil.NewCostFixture(), nil)

		result, err := svc.Update("cost-1", "user-1", incoming, nil)

		assert.NoError(t, err)
		assert.Equal(t, "cost-1", result.ID)
	})

	t.Run("success - ownerPercentage mantido quando nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		existing := testutil.NewCostFixture()
		incoming := domains.NewCost("user-1", "", "Jantar Atualizado", "Dinner", 150.0, 0)

		mockRepo.EXPECT().FindByID("cost-1").Return(existing, nil)
		mockRepo.EXPECT().Update("cost-1", gomock.Any()).DoAndReturn(func(_ string, cost *domains.Cost) (*domains.Cost, error) {
			assert.Equal(t, existing.OwnerPercentage, cost.OwnerPercentage)
			return testutil.NewCostFixture(), nil
		})

		_, err := svc.Update("cost-1", "user-1", incoming, nil)

		assert.NoError(t, err)
	})

	t.Run("invalid ownerPercentage - zero", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)

		_, err := svc.Update("cost-1", "user-1", domains.NewCost("user-1", "", "X", "Dinner", 100, 0), float64Ptr(0))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("invalid ownerPercentage - 100", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)

		_, err := svc.Update("cost-1", "user-1", domains.NewCost("user-1", "", "X", "Dinner", 100, 0), float64Ptr(100))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(nil, domains.NewNotFoundError("cost not found"))

		_, err := svc.Update("cost-1", "user-1", domains.NewCost("user-1", "", "X", "Dinner", 100, 0), nil)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)

		_, err := svc.Update("cost-1", "user-2", domains.NewCost("user-2", "", "X", "Dinner", 100, 0), nil)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestCostService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)
		mockRepo.EXPECT().Delete("cost-1").Return(nil)

		err := svc.Delete("cost-1", "user-1")
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(nil, domains.NewNotFoundError("cost not found"))

		err := svc.Delete("cost-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID("cost-1").Return(testutil.NewCostFixture(), nil)

		err := svc.Delete("cost-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}
