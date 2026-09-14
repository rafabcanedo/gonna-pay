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

func TestCreateGroup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		m := testutil.NewGroupMock()
		mockService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.Group, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"category": m.Category,
		})

		gc.CreateGroup(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.GroupResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.Name, body.Name)
		assert.Equal(t, m.Category, body.Category)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{})

		gc.CreateGroup(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     testutil.GroupName,
			"category": "Invalid",
		})

		gc.CreateGroup(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("service error - forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     testutil.GroupName,
			"category": testutil.GroupCategory,
		})

		gc.CreateGroup(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestFindAllGroups(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		m1 := testutil.NewGroupMock()
		m2 := testutil.NewGroupMock()
		m2.Group.ID = "group-2"
		mockService.EXPECT().FindAll(testutil.UserID, 1, 20).Return([]*domains.Group{m1.Group, m2.Group}, int64(2), nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, nil)

		gc.FindAllGroups(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.PaginatedResponse[response.GroupResponse]
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, int64(2), body.Total)
		assert.Equal(t, 1, body.TotalPages)
	})

	t.Run("custom page and limit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		m := testutil.NewGroupMock()
		mockService.EXPECT().FindAll(testutil.UserID, 2, 10).Return([]*domains.Group{m.Group}, int64(11), nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, url.Values{"page": {"2"}, "limit": {"10"}})

		gc.FindAllGroups(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.PaginatedResponse[response.GroupResponse]
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, 2, body.Page)
		assert.Equal(t, 10, body.Limit)
		assert.Equal(t, int64(11), body.Total)
		assert.Equal(t, 2, body.TotalPages)
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().FindAll(testutil.UserID, 1, 20).Return(nil, int64(0), domains.NewNotFoundError("not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, nil, nil)

		gc.FindAllGroups(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestFindGroupByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		m := testutil.NewGroupMock()
		mockService.EXPECT().FindByID(m.ID, testutil.UserID).Return(m.Group, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: m.ID}}, nil)

		gc.FindGroupByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.GroupDetailResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.ID, body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().FindByID(testutil.GroupID, testutil.UserID).Return(nil, domains.NewNotFoundError("group not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}}, nil)

		gc.FindGroupByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().FindByID(testutil.GroupID, testutil.UserID2).Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID2)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}}, nil)

		gc.FindGroupByID(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestUpdateGroup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		m := testutil.NewGroupMock()
		m.Group.Name = testutil.GroupUpdatedName
		mockService.EXPECT().Update(gomock.Any(), testutil.UserID).Return(m.Group, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: m.ID}}, map[string]any{
			"name": testutil.GroupUpdatedName,
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.GroupResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, testutil.GroupUpdatedName, body.Name)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}}, map[string]any{
			"category": "Invalid",
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Update(gomock.Any(), testutil.UserID).Return(nil, domains.NewNotFoundError("group not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}}, map[string]any{
			"name": testutil.GroupUpdatedName,
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Update(gomock.Any(), testutil.UserID2).Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID2)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}}, map[string]any{
			"name": testutil.GroupUpdatedName,
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestDeleteGroup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Delete(testutil.GroupID, testutil.UserID).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}})

		gc.DeleteGroup(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Delete(testutil.GroupID, testutil.UserID).Return(domains.NewNotFoundError("group not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.GroupID}})

		gc.DeleteGroup(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestAddMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().AddMember(testutil.GroupID, testutil.ContactID, testutil.UserID).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, gin.Params{
			{Key: "id", Value: testutil.GroupID},
			{Key: "contactId", Value: testutil.ContactID},
		}, nil)

		gc.AddMember(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("conflict - already a member", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().AddMember(testutil.GroupID, testutil.ContactID, testutil.UserID).Return(domains.NewConflictError("already a member"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, gin.Params{
			{Key: "id", Value: testutil.GroupID},
			{Key: "contactId", Value: testutil.ContactID},
		}, nil)

		gc.AddMember(ctx)

		assert.Equal(t, http.StatusConflict, rec.Code)
	})
}

func TestRemoveMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().RemoveMember(testutil.GroupID, testutil.ContactID, testutil.UserID).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{
			{Key: "id", Value: testutil.GroupID},
			{Key: "contactId", Value: testutil.ContactID},
		})

		gc.RemoveMember(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().RemoveMember(testutil.GroupID, testutil.ContactID, testutil.UserID).Return(domains.NewNotFoundError("member not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{
			{Key: "id", Value: testutil.GroupID},
			{Key: "contactId", Value: testutil.ContactID},
		})

		gc.RemoveMember(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
