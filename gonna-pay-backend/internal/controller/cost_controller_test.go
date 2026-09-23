package controller_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/mocks"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/response"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		mockService.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(m.Cost, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":        m.Name,
			"totalValue":      m.TotalValue,
			"category":        m.Category,
			"ownerPercentage": m.OwnerPercentage,
		})

		cc.CreateCost(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.CostDetailResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.Name, body.CostName)
		assert.Equal(t, m.TotalValue, body.TotalValue)
		assert.Len(t, body.Splits, 1)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"costName": testutil.CostName,
		})

		cc.CreateCost(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   testutil.CostName,
			"totalValue": 100.0,
			"category":   "Invalid",
		})

		cc.CreateCost(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - totalValue must be gt 0", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   testutil.CostName,
			"totalValue": 0,
			"category":   testutil.CostCategory,
		})

		cc.CreateCost(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestFindAllCosts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m1 := testutil.NewCostMock()
		m2 := testutil.NewCostMock()
		m2.Cost.ID = "cost-2"
		mockService.EXPECT().FindAll(gomock.Any(), testutil.UserID, 1, 20, domains.CostFilters{}).Return([]*domains.Cost{m1.Cost, m2.Cost}, int64(2), nil)
		mockService.EXPECT().FindStats(gomock.Any(), testutil.UserID, domains.CostFilters{}).Return(&domains.CostStats{ByCategory: []domains.CostCategoryBreakdown{}}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, nil)

		cc.FindAllCosts(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.CostsListResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, int64(2), body.Total)
		assert.Equal(t, 1, body.TotalPages)
	})

	t.Run("custom page and limit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		mockService.EXPECT().FindAll(gomock.Any(), testutil.UserID, 2, 10, domains.CostFilters{}).Return([]*domains.Cost{m.Cost}, int64(11), nil)
		mockService.EXPECT().FindStats(gomock.Any(), testutil.UserID, domains.CostFilters{}).Return(&domains.CostStats{ByCategory: []domains.CostCategoryBreakdown{}}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, url.Values{"page": {"2"}, "limit": {"10"}})

		cc.FindAllCosts(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.CostsListResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, 2, body.Page)
		assert.Equal(t, 10, body.Limit)
		assert.Equal(t, int64(11), body.Total)
		assert.Equal(t, 2, body.TotalPages)
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().FindAll(gomock.Any(), testutil.UserID, 1, 20, domains.CostFilters{}).Return(nil, int64(0), domains.NewNotFoundError("not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, nil)

		cc.FindAllCosts(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestFindAllCosts_Filters(t *testing.T) {
	t.Run("filters by category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		expectedFilters := domains.CostFilters{Category: "Dinner"}
		mockService.EXPECT().FindAll(gomock.Any(), testutil.UserID, 1, 20, expectedFilters).Return([]*domains.Cost{m.Cost}, int64(1), nil)
		mockService.EXPECT().FindStats(gomock.Any(), testutil.UserID, expectedFilters).Return(&domains.CostStats{ByCategory: []domains.CostCategoryBreakdown{}}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, url.Values{"category": {"Dinner"}})

		cc.FindAllCosts(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("filters by period and type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		expectedFilters := domains.CostFilters{
			Period: domains.CostPeriodMonth,
			Type:   domains.CostTypeSolo,
		}
		mockService.EXPECT().FindAll(gomock.Any(), testutil.UserID, 1, 20, expectedFilters).Return([]*domains.Cost{m.Cost}, int64(1), nil)
		mockService.EXPECT().FindStats(gomock.Any(), testutil.UserID, expectedFilters).Return(&domains.CostStats{ByCategory: []domains.CostCategoryBreakdown{}}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, url.Values{"period": {"month"}, "type": {"solo"}})

		cc.FindAllCosts(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestFindCostByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		mockService.EXPECT().FindByID(gomock.Any(), m.ID, testutil.UserID).Return(m.Cost, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: m.ID}}, nil)

		cc.FindCostByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.CostDetailResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.ID, body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().FindByID(gomock.Any(), testutil.CostID, testutil.UserID).Return(nil, domains.NewNotFoundError("cost not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.CostID}}, nil)

		cc.FindCostByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().FindByID(gomock.Any(), testutil.CostID, testutil.UserID2).Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID2)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.CostID}}, nil)

		cc.FindCostByID(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestUpdateCost(t *testing.T) {
	t.Run("success - full update", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		mockService.EXPECT().Update(gomock.Any(), m.ID, testutil.UserID, gomock.Any(), gomock.Any()).Return(m.Cost, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: m.ID}}, map[string]any{
			"costName":   "Jantar Atualizado",
			"totalValue": 150.0,
			"category":   m.Category,
		})

		cc.UpdateCost(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("success - partial update (only category)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		m := testutil.NewCostMock()
		mockService.EXPECT().Update(gomock.Any(), m.ID, testutil.UserID, gomock.Any(), gomock.Any()).Return(m.Cost, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: m.ID}}, map[string]any{
			"category": "Lunch",
		})

		cc.UpdateCost(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.CostID}}, map[string]any{
			"category": "Invalid",
		})

		cc.UpdateCost(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().Update(gomock.Any(), testutil.CostID, testutil.UserID, gomock.Any(), gomock.Any()).Return(nil, domains.NewNotFoundError("cost not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.CostID}}, map[string]any{
			"category": testutil.CostCategory,
		})

		cc.UpdateCost(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestDeleteCost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().Delete(gomock.Any(), testutil.CostID, testutil.UserID).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.CostID}})

		cc.DeleteCost(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().Delete(gomock.Any(), testutil.CostID, testutil.UserID).Return(domains.NewNotFoundError("cost not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.CostID}})

		cc.DeleteCost(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
