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

func TestCreateGroup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		group := testutil.NewGroupFixture()
		mockService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(group, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Viagem",
			"category": "Travel",
		})

		gc.CreateGroup(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.GroupResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Viagem", body.Name)
		assert.Equal(t, "Travel", body.Category)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{})

		gc.CreateGroup(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Viagem",
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
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Viagem",
			"category": "Travel",
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

		g1 := testutil.NewGroupFixture()
		g2 := testutil.NewGroupFixture()
		g2.ID = "group-2"
		mockService.EXPECT().FindAll("user-1").Return([]*domains.Group{g1, g2}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, nil, nil)

		gc.FindAllGroups(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body map[string]any
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, float64(2), body["total"])
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().FindAll("user-1").Return(nil, domains.NewNotFoundError("not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
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

		group := testutil.NewGroupFixture()
		mockService.EXPECT().FindByID("group-1", "user-1").Return(group, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "group-1"}}, nil)

		gc.FindGroupByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.GroupDetailResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "group-1", body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().FindByID("group-1", "user-1").Return(nil, domains.NewNotFoundError("group not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "group-1"}}, nil)

		gc.FindGroupByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().FindByID("group-1", "user-2").Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-2")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "group-1"}}, nil)

		gc.FindGroupByID(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestUpdateGroup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		updated := testutil.NewGroupFixture()
		updated.Name = "Novo Nome"
		mockService.EXPECT().Update(gomock.Any(), "user-1").Return(updated, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: "group-1"}}, map[string]any{
			"name": "Novo Nome",
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.GroupResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Novo Nome", body.Name)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: "group-1"}}, map[string]any{
			"category": "Invalid",
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Update(gomock.Any(), "user-1").Return(nil, domains.NewNotFoundError("group not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: "group-1"}}, map[string]any{
			"name": "Novo Nome",
		})

		gc.UpdateGroup(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Update(gomock.Any(), "user-2").Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-2")
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: "group-1"}}, map[string]any{
			"name": "Novo Nome",
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

		mockService.EXPECT().Delete("group-1", "user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "group-1"}})

		gc.DeleteGroup(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().Delete("group-1", "user-1").Return(domains.NewNotFoundError("group not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "group-1"}})

		gc.DeleteGroup(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestAddMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().AddMember("group-1", "contact-1", "user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, gin.Params{
			{Key: "id", Value: "group-1"},
			{Key: "contactId", Value: "contact-1"},
		}, nil)

		gc.AddMember(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("conflict - already a member", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().AddMember("group-1", "contact-1", "user-1").Return(domains.NewConflictError("already a member"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, gin.Params{
			{Key: "id", Value: "group-1"},
			{Key: "contactId", Value: "contact-1"},
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

		mockService.EXPECT().RemoveMember("group-1", "contact-1", "user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{
			{Key: "id", Value: "group-1"},
			{Key: "contactId", Value: "contact-1"},
		})

		gc.RemoveMember(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockGroupService(ctrl)
		gc := controller.NewGroupController(mockService)

		mockService.EXPECT().RemoveMember("group-1", "contact-1", "user-1").Return(domains.NewNotFoundError("member not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{
			{Key: "id", Value: "group-1"},
			{Key: "contactId", Value: "contact-1"},
		})

		gc.RemoveMember(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
