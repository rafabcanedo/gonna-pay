package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/rest_errors"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			restErr := rest_errors.NewUnauthorizedRequestError("missing access token")
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		claims, err := ParseAccessToken(cookie)
		if err != nil {
			restErr := rest_errors.NewUnauthorizedRequestError("invalid or expired token")
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Set("userName", claims.Name)
		c.Next()
	}
}
