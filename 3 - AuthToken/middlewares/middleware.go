package middleware

import (
	tokenutils "auth_token/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware Function is to authenticate user
// in any route that needs authorization.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenErr := tokenutils.TokenIsValid(c.Request)

		if tokenErr == nil {
			c.Next()

			return
		}

		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": tokenErr.Error(),
		})

		c.Abort()

		return
	}
}
