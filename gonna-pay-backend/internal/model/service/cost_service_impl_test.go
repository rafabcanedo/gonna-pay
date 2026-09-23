package service_test

import (
	"context"
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
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		cost := domains.NewCost(testutil.UserID, "", m.Name, m.Category, m.TotalValue, 0)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.Cost, nil)

		result, err := svc.Create(context.Background(), cost, nil)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
		assert.Equal(t, float64(100), cost.OwnerPercentage)
	})

	t.Run("success with group and explicit ownerPercentage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, m.Name, m.Category, m.TotalValue, 0)
		members := []domains.Member{{ID: testutil.ContactID, Name: testutil.ContactName}, {ID: "contact-2", Name: "Pedro"}}

		mockGroupRepo.EXPECT().FindByID(gomock.Any(), testutil.GroupID).Return(g.Group, nil)
		mockGroupRepo.EXPECT().GetMembers(gomock.Any(), testutil.GroupID).Return(members, nil)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.Cost, nil)

		result, err := svc.Create(context.Background(), cost, float64Ptr(40.0))

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
		assert.Equal(t, 40.0, cost.OwnerPercentage)
	})

	t.Run("success with group and auto ownerPercentage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, m.Name, m.Category, m.TotalValue, 0)
		members := []domains.Member{
			{ID: testutil.ContactID, Name: testutil.ContactName},
			{ID: "contact-2", Name: "Pedro"},
			{ID: "contact-3", Name: "João"},
		}

		mockGroupRepo.EXPECT().FindByID(gomock.Any(), testutil.GroupID).Return(g.Group, nil)
		mockGroupRepo.EXPECT().GetMembers(gomock.Any(), testutil.GroupID).Return(members, nil)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.Cost, nil)

		_, err := svc.Create(context.Background(), cost, nil)

		assert.NoError(t, err)
		assert.Equal(t, 25.0, cost.OwnerPercentage)
	})

	t.Run("invalid ownerPercentage - zero", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, testutil.CostName, testutil.CostCategory, 100.0, 0)
		mockGroupRepo.EXPECT().FindByID(gomock.Any(), testutil.GroupID).Return(g.Group, nil)
		mockGroupRepo.EXPECT().GetMembers(gomock.Any(), testutil.GroupID).Return([]domains.Member{{ID: testutil.ContactID, Name: testutil.ContactName}}, nil)

		_, err := svc.Create(context.Background(), cost, float64Ptr(0))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("invalid ownerPercentage - 100", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, testutil.CostName, testutil.CostCategory, 100.0, 0)
		mockGroupRepo.EXPECT().FindByID(gomock.Any(), testutil.GroupID).Return(g.Group, nil)
		mockGroupRepo.EXPECT().GetMembers(gomock.Any(), testutil.GroupID).Return([]domains.Member{{ID: testutil.ContactID, Name: testutil.ContactName}}, nil)

		_, err := svc.Create(context.Background(), cost, float64Ptr(100))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("repo error on GetGroupMembers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		g := testutil.NewGroupMock()
		cost := domains.NewCost(testutil.UserID, testutil.GroupID, testutil.CostName, testutil.CostCategory, 100.0, 0)
		mockGroupRepo.EXPECT().FindByID(gomock.Any(), testutil.GroupID).Return(g.Group, nil)
		mockGroupRepo.EXPECT().GetMembers(gomock.Any(), testutil.GroupID).Return(nil, errors.New("db error"))

		_, err := svc.Create(context.Background(), cost, nil)
		assert.Error(t, err)
	})
}

func TestCostService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m1 := testutil.NewCostMock()
		m2 := testutil.NewCostMock()
		m2.Cost.ID = "cost-2"
		mockRepo.EXPECT().FindAll(gomock.Any(), testutil.UserID, 20, 0, domains.CostFilters{}).Return([]*domains.Cost{m1.Cost, m2.Cost}, int64(2), nil)

		result, total, err := svc.FindAll(context.Background(), testutil.UserID, 1, 20, domains.CostFilters{})

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, int64(2), total)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		mockRepo.EXPECT().FindAll(gomock.Any(), testutil.UserID, 20, 0, domains.CostFilters{}).Return(nil, int64(0), errors.New("db error"))

		_, _, err := svc.FindAll(context.Background(), testutil.UserID, 1, 20, domains.CostFilters{})
		assert.Error(t, err)
	})
}

func TestCostService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)

		result, err := svc.FindByID(context.Background(), m.ID, testutil.UserID)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		mockRepo.EXPECT().FindByID(gomock.Any(), testutil.CostID).Return(nil, domains.NewNotFoundError("cost not found"))

		_, err := svc.FindByID(context.Background(), testutil.CostID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)

		_, err := svc.FindByID(context.Background(), m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestCostService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		incoming := domains.NewCost(testutil.UserID, "", "Jantar Atualizado", m.Category, 150.0, 0)

		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)
		mockRepo.EXPECT().Update(gomock.Any(), m.ID, gomock.Any()).Return(m.Cost, nil)

		result, err := svc.Update(context.Background(), m.ID, testutil.UserID, incoming, nil)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("success - ownerPercentage mantido quando nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		incoming := domains.NewCost(testutil.UserID, "", "Jantar Atualizado", m.Category, 150.0, 0)

		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)
		mockRepo.EXPECT().Update(gomock.Any(), m.ID, gomock.Any()).DoAndReturn(func(_ context.Context, _ string, cost *domains.Cost) (*domains.Cost, error) {
			assert.Equal(t, m.OwnerPercentage, cost.OwnerPercentage)
			return m.Cost, nil
		})

		_, err := svc.Update(context.Background(), m.ID, testutil.UserID, incoming, nil)

		assert.NoError(t, err)
	})

	t.Run("invalid ownerPercentage - zero", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)

		_, err := svc.Update(context.Background(), m.ID, testutil.UserID, domains.NewCost(testutil.UserID, "", "X", testutil.CostCategory, 100, 0), float64Ptr(0))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("invalid ownerPercentage - 100", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)

		_, err := svc.Update(context.Background(), m.ID, testutil.UserID, domains.NewCost(testutil.UserID, "", "X", testutil.CostCategory, 100, 0), float64Ptr(100))

		assert.ErrorIs(t, err, domains.ErrInvalidInput)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		mockRepo.EXPECT().FindByID(gomock.Any(), testutil.CostID).Return(nil, domains.NewNotFoundError("cost not found"))

		_, err := svc.Update(context.Background(), testutil.CostID, testutil.UserID, domains.NewCost(testutil.UserID, "", "X", testutil.CostCategory, 100, 0), nil)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)

		_, err := svc.Update(context.Background(), m.ID, testutil.UserID2, domains.NewCost(testutil.UserID2, "", "X", testutil.CostCategory, 100, 0), nil)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestCostService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)
		mockRepo.EXPECT().Delete(gomock.Any(), m.ID).Return(nil)

		err := svc.Delete(context.Background(), m.ID, testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		mockRepo.EXPECT().FindByID(gomock.Any(), testutil.CostID).Return(nil, domains.NewNotFoundError("cost not found"))

		err := svc.Delete(context.Background(), testutil.CostID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockCostRepository(ctrl)
		mockGroupRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewCostService(mockRepo, mockGroupRepo)

		m := testutil.NewCostMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.Cost, nil)

		err := svc.Delete(context.Background(), m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}
