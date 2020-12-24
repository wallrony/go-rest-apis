package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middlewares "auth_token/middlewares"
	models "auth_token/models"
	tokenutils "auth_token/utils/token"
)

var (
	router = gin.Default()
)

func main() {
	apiGroup := router.Group("/api")

	apiGroup.POST("/login", Login)
	apiGroup.POST("/test", middlewares.AuthMiddleware(), Teste)

	router.Use(gin.Recovery())

	router.Run(":8001")
}

var user = models.User{
	ID:       1,
	Username: "wallrony",
	Name:     "Wallisson Rony",
	Password: "123456",
}

// Teste function is to test AuthMiddleware.
func Teste(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "authenticated",
	})
}

// Login Function verify credentials of an user and return
// an authorization token
func Login(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Invalid body!",
		})

		return
	}

	// Compara o usuário
	if u.Username != user.Username || u.Password != user.Password {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid credentials!",
		})

		return
	}

	td, err := tokenutils.CreateToken(u.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Error creating an authorization code.",
		})

		return
	}

	tokens := map[string]string{
		"access_token": td.AccessToken,
	}

	c.JSON(http.StatusOK, tokens)
}
