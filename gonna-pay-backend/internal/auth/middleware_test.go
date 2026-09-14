package auth_test

import (
	"net/http"
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/auth"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Run("missing cookie returns 401 and aborts", func(t *testing.T) {
		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)

		handler := auth.Middleware()
		handler(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.True(t, ctx.IsAborted())
	})

	t.Run("invalid token returns 401 and aborts", func(t *testing.T) {
		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)
		ctx.Request.Header.Set("Cookie", "access_token=invalid.token.here")

		handler := auth.Middleware()
		handler(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.True(t, ctx.IsAborted())
	})

	t.Run("valid token sets userID and userName in context", func(t *testing.T) {
		token, _ := auth.GenerateAccessToken("user-1", "Rafael")

		ctx, _ := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)
		ctx.Request.Header.Set("Cookie", "access_token="+token)

		handler := auth.Middleware()
		handler(ctx)

		assert.False(t, ctx.IsAborted())
		assert.Equal(t, "user-1", ctx.GetString("userID"))
		assert.Equal(t, "Rafael", ctx.GetString("userName"))
	})
}
