package auth_test

import (
	"os"
	"testing"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/auth"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("JWT_SECRET", "test-secret-key")
}

func TestGenerateAccessToken(t *testing.T) {
	t.Run("success - returns non-empty token", func(t *testing.T) {
		token, err := auth.GenerateAccessToken("user-1", "Rafael")

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("token contains correct claims", func(t *testing.T) {
		token, err := auth.GenerateAccessToken("user-1", "Rafael")
		assert.NoError(t, err)

		claims, err := auth.ParseAccessToken(token)

		assert.NoError(t, err)
		assert.Equal(t, "user-1", claims.Subject)
		assert.Equal(t, "Rafael", claims.Name)
	})
}

func TestParseAccessToken(t *testing.T) {
	t.Run("success - valid token parsed correctly", func(t *testing.T) {
		token, _ := auth.GenerateAccessToken("user-1", "Rafael")

		claims, err := auth.ParseAccessToken(token)

		assert.NoError(t, err)
		assert.Equal(t, "user-1", claims.Subject)
		assert.Equal(t, "Rafael", claims.Name)
	})

	t.Run("invalid token - returns error", func(t *testing.T) {
		_, err := auth.ParseAccessToken("token.invalido.aqui")

		assert.Error(t, err)
	})

	t.Run("token signed with different secret - returns error", func(t *testing.T) {
		os.Setenv("JWT_SECRET", "secret-diferente")
		token, _ := auth.GenerateAccessToken("user-1", "Rafael")

		os.Setenv("JWT_SECRET", "test-secret-key")
		_, err := auth.ParseAccessToken(token)

		assert.Error(t, err)
	})
}

func TestGenerateRefreshToken(t *testing.T) {
	t.Run("success - returns 64-char hex string", func(t *testing.T) {
		token, err := auth.GenerateRefreshToken()

		assert.NoError(t, err)
		assert.Len(t, token, 64)
	})

	t.Run("two calls return different tokens", func(t *testing.T) {
		token1, _ := auth.GenerateRefreshToken()
		token2, _ := auth.GenerateRefreshToken()

		assert.NotEqual(t, token1, token2)
	})
}

func TestHashToken(t *testing.T) {
	t.Run("same input always produces same hash", func(t *testing.T) {
		hash1 := auth.HashToken("meu-token")
		hash2 := auth.HashToken("meu-token")

		assert.Equal(t, hash1, hash2)
	})

	t.Run("different inputs produce different hashes", func(t *testing.T) {
		hash1 := auth.HashToken("token-a")
		hash2 := auth.HashToken("token-b")

		assert.NotEqual(t, hash1, hash2)
	})
}
