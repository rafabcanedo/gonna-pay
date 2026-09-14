package domains_test

import (
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/stretchr/testify/assert"
)

func TestUser_EncryptPassword(t *testing.T) {
	t.Run("password is hashed after encryption", func(t *testing.T) {
		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")

		err := user.EncryptPassword()

		assert.NoError(t, err)
		assert.NotEqual(t, "senha123", user.Password)
	})

	t.Run("two calls produce different hashes", func(t *testing.T) {
		user1 := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		user2 := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")

		user1.EncryptPassword()
		user2.EncryptPassword()

		assert.NotEqual(t, user1.Password, user2.Password)
	})
}

func TestUser_ComparePassword(t *testing.T) {
	t.Run("correct password returns no error", func(t *testing.T) {
		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		user.EncryptPassword()

		err := user.ComparePassword("senha123")

		assert.NoError(t, err)
	})

	t.Run("wrong password returns error", func(t *testing.T) {
		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		user.EncryptPassword()

		err := user.ComparePassword("wrong-password")

		assert.Error(t, err)
	})
}
