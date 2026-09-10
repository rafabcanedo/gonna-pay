package service_test

import (
	"errors"
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/mocks"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/service"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func float64Ptr(v float64) *float64 { return &v }

func TestCostService_Create(t *testing.T) {
	t.Run("success without group - owner gets 100%", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		cost := domains.NewCost(testutil.UserID, "", m.Name, m.Category, m.TotalValue, 0)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Nil()).Return(m.Cost, nil)

		result, err := svc.Create(cost, nil)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
		assert.Equal(t, float64(100), cost.OwnerPercentage)
	})

	t.Run("success with group and explicit ownerPercentage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, m.Name, m.Category, m.TotalValue, 0)
		members := []domains.Member{{ID: testutil.ContactID, Name: testutil.ContactName}, {ID: "contact-2", Name: "Pedro"}}

		mockRepo.EXPECT().GetGroupByID(testutil.GroupID).Return(g.Group, nil)
		mockRepo.EXPECT().GetGroupMembers(testutil.GroupID).Return(members, nil)
		mockRepo.EXPECT().Create(gomock.Any(), members).Return(m.Cost, nil)

		result, err := svc.Create(cost, float64Ptr(40.0))

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
		assert.Equal(t, 40.0, cost.OwnerPercentage)
	})

	t.Run("success with group and auto ownerPercentage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, m.Name, m.Category, m.TotalValue, 0)
		members := []domains.Member{
			{ID: testutil.ContactID, Name: testutil.ContactName},
			{ID: "contact-2", Name: "Pedro"},
			{ID: "contact-3", Name: "João"},
		}

		mockRepo.EXPECT().GetGroupByID(testutil.GroupID).Return(g.Group, nil)
		mockRepo.EXPECT().GetGroupMembers(testutil.GroupID).Return(members, nil)
		mockRepo.EXPECT().Create(gomock.Any(), members).Return(m.Cost, nil)

		_, err := svc.Create(cost, nil)

		assert.NoError(t, err)
		assert.Equal(t, 25.0, cost.OwnerPercentage)
	})

	t.Run("invalid ownerPercentage - zero", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, testutil.CostName, testutil.CostCategory, 100.0, 0)
		mockRepo.EXPECT().GetGroupByID(testutil.GroupID).Return(g.Group, nil)
		mockRepo.EXPECT().GetGroupMembers(testutil.GroupID).Return([]domains.Member{{ID: testutil.ContactID, Name: testutil.ContactName}}, nil)

		_, err := svc.Create(cost, float64Ptr(0))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("invalid ownerPercentage - 100", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, testutil.CostName, testutil.CostCategory, 100.0, 0)
		mockRepo.EXPECT().GetGroupByID(testutil.GroupID).Return(g.Group, nil)
		mockRepo.EXPECT().GetGroupMembers(testutil.GroupID).Return([]domains.Member{{ID: testutil.ContactID, Name: testutil.ContactName}}, nil)

		_, err := svc.Create(cost, float64Ptr(100))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("repo error on GetGroupMembers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, testutil.CostName, testutil.CostCategory, 100.0, 0)
		mockRepo.EXPECT().GetGroupByID(testutil.GroupID).Return(g.Group, nil)
		mockRepo.EXPECT().GetGroupMembers(testutil.GroupID).Return(nil, errors.New("db error"))

		_, err := svc.Create(cost, nil)
		assert.Error(t, err)
	})
}

func TestCostService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m1 := testutil.NewCostMock()
		m2 := testutil.NewCostMock()
		m2.Cost.ID = "cost-2"
		mockRepo.EXPECT().FindAll(testutil.UserID).Return([]*domains.Cost{m1.Cost, m2.Cost}, nil)

		result, err := svc.FindAll(testutil.UserID)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})
}

func TestCostService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)

		result, err := svc.FindByID(m.ID, testutil.UserID)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.CostID).Return(nil, domains.NewNotFoundError("cost not found"))

		_, err := svc.FindByID(testutil.CostID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)

		_, err := svc.FindByID(m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestCostService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		incoming := domains.NewCost(testutil.UserID, "", "Jantar Atualizado", m.Category, 150.0, 0)

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)
		mockRepo.EXPECT().Update(m.ID, gomock.Any()).Return(m.Cost, nil)

		result, err := svc.Update(m.ID, testutil.UserID, incoming, nil)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("success - ownerPercentage mantido quando nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		incoming := domains.NewCost(testutil.UserID, "", "Jantar Atualizado", m.Category, 150.0, 0)

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)
		mockRepo.EXPECT().Update(m.ID, gomock.Any()).DoAndReturn(func(_ string, cost *domains.Cost) (*domains.Cost, error) {
			assert.Equal(t, m.OwnerPercentage, cost.OwnerPercentage)
			return m.Cost, nil
		})

		_, err := svc.Update(m.ID, testutil.UserID, incoming, nil)

		assert.NoError(t, err)
	})

	t.Run("invalid ownerPercentage - zero", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)

		_, err := svc.Update(m.ID, testutil.UserID, domains.NewCost(testutil.UserID, "", "X", testutil.CostCategory, 100, 0), float64Ptr(0))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("invalid ownerPercentage - 100", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)

		_, err := svc.Update(m.ID, testutil.UserID, domains.NewCost(testutil.UserID, "", "X", testutil.CostCategory, 100, 0), float64Ptr(100))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.CostID).Return(nil, domains.NewNotFoundError("cost not found"))

		_, err := svc.Update(testutil.CostID, testutil.UserID, domains.NewCost(testutil.UserID, "", "X", testutil.CostCategory, 100, 0), nil)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)

		_, err := svc.Update(m.ID, testutil.UserID2, domains.NewCost(testutil.UserID2, "", "X", testutil.CostCategory, 100, 0), nil)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestCostService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)
		mockRepo.EXPECT().Delete(m.ID).Return(nil)

		err := svc.Delete(m.ID, testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.CostID).Return(nil, domains.NewNotFoundError("cost not found"))

		err := svc.Delete(testutil.CostID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		svc := service.NewCostService(mockRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Cost, nil)

		err := svc.Delete(m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}
