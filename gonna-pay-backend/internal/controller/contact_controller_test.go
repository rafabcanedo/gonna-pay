package controller_test

import (
	"encoding/json"
	"net/http"
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

func TestCreateContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		m := testutil.NewContactMock()
		mockService.EXPECT().Create(gomock.Any()).Return(m.Contact, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"phone":    m.Phone,
			"category": m.Category,
		})

		cc.CreateContact(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.Name, body.Name)
		assert.Equal(t, m.Category, body.Category)
	})

	t.Run("validation error - missing required fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
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
		testutil.SetAuthUser(ctx, testutil.UserID)
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

		m1 := testutil.NewContactMock()
		m2 := testutil.NewContactMock()
		m2.Contact.ID = "contact-2"
		mockService.EXPECT().FindAll(testutil.UserID).Return([]*domains.Contact{m1.Contact, m2.Contact}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
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

		m := testutil.NewContactMock()
		mockService.EXPECT().FindByID(m.ID, testutil.UserID).Return(m.Contact, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: m.ID}}, nil)

		cc.FindContactByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.ID, body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().FindByID(testutil.ContactID, testutil.UserID).Return(nil, domains.NewNotFoundError("contact not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}}, nil)

		cc.FindContactByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().FindByID(testutil.ContactID, testutil.UserID2).Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID2)
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}}, nil)

		cc.FindContactByID(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestUpdateContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		m := testutil.NewContactMock()
		m.Contact.Name = testutil.ContactUpdatedName
		mockService.EXPECT().Update(gomock.Any()).Return(m.Contact, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: m.ID}}, map[string]any{
			"name": testutil.ContactUpdatedName,
		})

		cc.UpdateContact(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.ContactResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, testutil.ContactUpdatedName, body.Name)
	})

	t.Run("validation error - invalid category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}}, map[string]any{
			"category": "Invalid",
		})

		cc.UpdateContact(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().Update(gomock.Any()).Return(nil, domains.NewNotFoundError("contact not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}}, map[string]any{
			"name": testutil.ContactUpdatedName,
		})

		cc.UpdateContact(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("forbidden", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().Update(gomock.Any()).Return(nil, domains.NewForbiddenError("access denied"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID2)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}}, map[string]any{
			"name": testutil.ContactUpdatedName,
		})

		cc.UpdateContact(ctx)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}

func TestDeleteContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().Delete(testutil.ContactID, testutil.UserID).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}})

		cc.DeleteContact(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockContactService(ctrl)
		cc := controller.NewContactController(mockService)

		mockService.EXPECT().Delete(testutil.ContactID, testutil.UserID).Return(domains.NewNotFoundError("contact not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.ContactID}})

		cc.DeleteContact(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
