package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	core "todo/routers/core"
)

// SetRoute function set all routes that precedes with
// 'api' prefix.
func SetRoute(dbInstance *sql.DB, router *gin.Engine) {
	apiGroup := router.Group("/api")

	core.SetRoute(dbInstance, apiGroup)
}
