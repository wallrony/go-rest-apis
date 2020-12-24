package middlewares

import (
	"net/http"
	"products/core/utils"

	"github.com/gin-gonic/gin"
)

var (
	authRoutes = []string{
		"/api/accounts/auth",
	}
)

// AuthMiddleware function verify, as a router
// middleware, if the user that's accessing
// the route is authorized.
func AuthMiddleware(context *gin.Context) {
	var isFreePath bool

	for _, path := range authRoutes {
		if path == context.Request.URL.Path {
			isFreePath = true

			break
		}
	}

	if isFreePath {
		context.Next()

		return
	}

	userID, err := utils.TokenIsValid(context.Request)

	if err == nil {
		context.Request.Header["user_id"] = []string{userID}
		context.Next()

		return
	}

	context.JSON(http.StatusUnauthorized, map[string]string{
		"message": err.Error(),
	})

	context.Abort()

	return
}
