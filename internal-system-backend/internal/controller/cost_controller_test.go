package controller_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/controller"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/mocks"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/testutil"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/response"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		cost := testutil.NewCostFixture()
		mockService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(cost, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		ownerPct := 50.0
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":        "Jantar",
			"totalValue":      100.0,
			"category":        "Dinner",
			"ownerPercentage": ownerPct,
		})

		cc.CreateCost(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.CostDetailResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Jantar", body.CostName)
		assert.Equal(t, 100.0, body.TotalValue)
		assert.Len(t, body.Splits, 1)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"costName": "Jantar",
		})

		cc.CreateCost(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   "Jantar",
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
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   "Jantar",
			"totalValue": 0,
			"category":   "Dinner",
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

		c1 := testutil.NewCostFixture()
		c2 := testutil.NewCostFixture()
		c2.ID = "cost-2"
		mockService.EXPECT().FindAll("user-1").Return([]*domains.Cost{c1, c2}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, nil, nil)

		cc.FindAllCosts(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body []response.CostResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Len(t, body, 2)
	})
}

func TestFindCostByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		cost := testutil.NewCostFixture()
		mockService.EXPECT().FindByID("cost-1", "user-1").Return(cost, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "cost-1"}}, nil)

		cc.FindCostByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.CostDetailResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "cost-1", body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().FindByID("cost-1", "user-1").Return(nil, domains.NewNotFoundError("cost not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "cost-1"}}, nil)

		cc.FindCostByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().FindByID("cost-1", "user-2").Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-2")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "cost-1"}}, nil)

		cc.FindCostByID(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestUpdateCost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		updated := testutil.NewCostFixture()
		mockService.EXPECT().Update("cost-1", "user-1", gomock.Any(), gomock.Any()).Return(updated, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePut(ctx, gin.Params{{Key: "id", Value: "cost-1"}}, map[string]any{
			"costName":   "Jantar Atualizado",
			"totalValue": 150.0,
			"category":   "Dinner",
		})

		cc.UpdateCost(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePut(ctx, gin.Params{{Key: "id", Value: "cost-1"}}, map[string]any{})

		cc.UpdateCost(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().Update("cost-1", "user-1", gomock.Any(), gomock.Any()).Return(nil, domains.NewNotFoundError("cost not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePut(ctx, gin.Params{{Key: "id", Value: "cost-1"}}, map[string]any{
			"costName":   "Jantar",
			"totalValue": 100.0,
			"category":   "Dinner",
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

		mockService.EXPECT().Delete("cost-1", "user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "cost-1"}})

		cc.DeleteCost(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockCostService(ctrl)
		cc := controller.NewCostController(mockService)

		mockService.EXPECT().Delete("cost-1", "user-1").Return(domains.NewNotFoundError("cost not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "cost-1"}})

		cc.DeleteCost(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
