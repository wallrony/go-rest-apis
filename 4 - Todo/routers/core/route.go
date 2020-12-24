package core

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	todoschema "todo/schemas/todo_schema"
)

// SetRoute function set all routes that precedes with
// 'core' prefix.
func SetRoute(dbInstance *sql.DB, router *gin.RouterGroup) {
	todoSchema, err := todoschema.GetSchema(dbInstance)

	if err != nil {
		panic(err)
	}

	coreGroup := router.Group("/core")

	coreGroup.POST("/todo", todoSchema.GraphQL)
}
