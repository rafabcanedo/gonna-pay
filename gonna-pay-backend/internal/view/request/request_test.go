package request_test

import (
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/request"
	"github.com/stretchr/testify/assert"
)

func TestLoginRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    testutil.UserEmail,
			"password": testutil.UserPassword,
		})

		var req request.LoginRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("missing email", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"password": testutil.UserPassword,
		})

		var req request.LoginRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    "not-an-email",
			"password": testutil.UserPassword,
		})

		var req request.LoginRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing password", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email": testutil.UserEmail,
		})

		var req request.LoginRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestUpdateProfileRequest(t *testing.T) {
	t.Run("valid with all fields", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"name":  m.Name,
			"email": m.Email,
			"phone": m.Phone,
		})

		var req request.UpdateProfileRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("valid with empty body", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{})

		var req request.UpdateProfileRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"email": "not-an-email",
		})

		var req request.UpdateProfileRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestCreateUserRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"password": testutil.UserPassword,
			"phone":    m.Phone,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("missing name", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    m.Email,
			"password": testutil.UserPassword,
			"phone":    m.Phone,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing email", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"password": testutil.UserPassword,
			"phone":    m.Phone,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    "not-an-email",
			"password": testutil.UserPassword,
			"phone":    m.Phone,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing password", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":  m.Name,
			"email": m.Email,
			"phone": m.Phone,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("password too short", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"password": "123",
			"phone":    m.Phone,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing phone", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"password": testutil.UserPassword,
		})

		var req request.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestUpdateUserRequest(t *testing.T) {
	t.Run("valid with all fields", func(t *testing.T) {
		m := testutil.NewUserMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"password": testutil.UserPassword,
			"phone":    m.Phone,
		})

		var req request.UpdateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("valid with empty body", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{})

		var req request.UpdateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"email": "not-an-email",
		})

		var req request.UpdateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("password too short", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"password": "123",
		})

		var req request.UpdateUserRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestCreateContactRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"phone":    m.Phone,
			"category": m.Category,
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("missing name", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    m.Email,
			"phone":    m.Phone,
			"category": m.Category,
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing email", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"phone":    m.Phone,
			"category": m.Category,
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    "not-an-email",
			"phone":    m.Phone,
			"category": m.Category,
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing phone", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"category": m.Category,
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing category", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":  m.Name,
			"email": m.Email,
			"phone": m.Phone,
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"phone":    m.Phone,
			"category": "Invalid",
		})

		var req request.CreateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestUpdateContactRequest(t *testing.T) {
	t.Run("valid with all fields", func(t *testing.T) {
		m := testutil.NewContactMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"phone":    m.Phone,
			"category": m.Category,
		})

		var req request.UpdateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("valid with empty body", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{})

		var req request.UpdateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"email": "not-an-email",
		})

		var req request.UpdateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"category": "Invalid",
		})

		var req request.UpdateContactRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestCreateGroupRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := testutil.NewGroupMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"category": m.Category,
		})

		var req request.CreateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("valid with memberIds", func(t *testing.T) {
		m := testutil.NewGroupMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":      m.Name,
			"category":  m.Category,
			"memberIds": []string{testutil.ContactID},
		})

		var req request.CreateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("missing name", func(t *testing.T) {
		m := testutil.NewGroupMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"category": m.Category,
		})

		var req request.CreateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing category", func(t *testing.T) {
		m := testutil.NewGroupMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name": m.Name,
		})

		var req request.CreateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		m := testutil.NewGroupMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"category": "Invalid",
		})

		var req request.CreateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestUpdateGroupRequest(t *testing.T) {
	t.Run("valid with empty body", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{})

		var req request.UpdateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"category": "Invalid",
		})

		var req request.UpdateGroupRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestCreateCostRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   m.Name,
			"totalValue": m.TotalValue,
			"category":   m.Category,
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("missing costName", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"totalValue": m.TotalValue,
			"category":   m.Category,
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing totalValue", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"costName": m.Name,
			"category": m.Category,
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("totalValue zero", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   m.Name,
			"totalValue": 0,
			"category":   m.Category,
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("totalValue negative", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   m.Name,
			"totalValue": -1.0,
			"category":   m.Category,
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("missing category", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   m.Name,
			"totalValue": m.TotalValue,
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		m := testutil.NewCostMock()
		ctx, _ := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"costName":   m.Name,
			"totalValue": m.TotalValue,
			"category":   "Invalid",
		})

		var req request.CreateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}

func TestUpdateCostRequest(t *testing.T) {
	t.Run("valid with empty body", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{})

		var req request.UpdateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.NoError(t, err)
	})

	t.Run("totalValue negative", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"totalValue": -1.0,
		})

		var req request.UpdateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		ctx, _ := testutil.NewTestContext()
		testutil.MakePatch(ctx, nil, map[string]any{
			"category": "Invalid",
		})

		var req request.UpdateCostRequest
		err := ctx.ShouldBindJSON(&req)

		assert.Error(t, err)
	})
}
