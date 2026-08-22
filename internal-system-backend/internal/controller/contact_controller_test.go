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

func TestCreateContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		contact := testutil.NewContactFixture()
		mockService.EXPECT().Create(gomock.Any()).Return(contact, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Ana",
			"email":    "ana@email.com",
			"phone":    "11999999999",
			"category": "Friend",
		})

		cc.CreateContact(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Ana", body.Name)
		assert.Equal(t, "Friend", body.Category)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"name": "Ana",
		})

		cc.CreateContact(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Ana",
			"email":    "ana@email.com",
			"phone":    "11999999999",
			"category": "Invalid",
		})

		cc.CreateContact(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestFindAllContacts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		c1 := testutil.NewContactFixture()
		c2 := testutil.NewContactFixture()
		c2.ID = "contact-2"
		mockService.EXPECT().FindAll("user-1").Return([]*domains.Contact{c1, c2}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, nil, nil)

		cc.FindAllContacts(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body []response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Len(t, body, 2)
	})
}

func TestFindContactByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		contact := testutil.NewContactFixture()
		mockService.EXPECT().FindByID("contact-1", "user-1").Return(contact, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "contact-1"}}, nil)

		cc.FindContactByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "contact-1", body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().FindByID("contact-1", "user-1").Return(nil, domains.NewNotFoundError("contact not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "contact-1"}}, nil)

		cc.FindContactByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().FindByID("contact-1", "user-2").Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-2")
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "contact-1"}}, nil)

		cc.FindContactByID(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestUpdateContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		updated := testutil.NewContactFixture()
		updated.Name = "Ana Silva"
		mockService.EXPECT().Update(gomock.Any()).Return(updated, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePut(ctx, gin.Params{{Key: "id", Value: "contact-1"}}, map[string]any{
			"name": "Ana Silva",
		})

		cc.UpdateContact(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Ana Silva", body.Name)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakePut(ctx, gin.Params{{Key: "id", Value: "contact-1"}}, map[string]any{
			"category": "Invalid",
		})

		cc.UpdateContact(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestDeleteContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().Delete("contact-1", "user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "contact-1"}})

		cc.DeleteContact(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().Delete("contact-1", "user-1").Return(domains.NewNotFoundError("contact not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "contact-1"}})

		cc.DeleteContact(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
