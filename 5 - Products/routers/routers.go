package routers

import (
	"database/sql"
	"products/core/models"
	"products/data/graphql/handlers"
	"products/data/graphql/schemas"
	"products/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// SetAPIRoute function set api route and your sub routes, like accounts and
// core routes.
func SetAPIRoute(db *sql.DB, router *gin.Engine) {
	apiGroup := router.Group("/api", middlewares.AuthMiddleware)

	setAccountsRoute(db, apiGroup)
	setCoreRoute(db, apiGroup)
}

func setAccountsRoute(db *sql.DB, apiRouter *gin.RouterGroup) {
	accountsRouter := apiRouter.Group("/accounts")

	setAuthRoute(db, accountsRouter)

	schemas.InitializeAccountsRoot(db)

	rootSchema := schemas.GetAccountsRoot()

	accountsHandler := createHandlerInstance(rootSchema)
	accountsRouter.POST("/users", accountsHandler.GraphQL)
}

func setCoreRoute(db *sql.DB, apiRouter *gin.RouterGroup) {
	coreRouter := apiRouter.Group("/core")

	schemas.InitializeCoreRoot(db)

	rootSchema := schemas.GetCoreRoot()

	coreHandler := createHandlerInstance(rootSchema)
	coreRouter.POST("/products", coreHandler.GraphQL)
}

func setAuthRoute(db *sql.DB, accountsRouter *gin.RouterGroup) {
	schemas.InitializeAuthRoot(db)

	rootSchema := schemas.GetAuthRoot()

	authHandler := createHandlerInstance(rootSchema)

	accountsRouter.POST("/auth", authHandler.GraphQL)
}

func createHandlerInstance(rootSchema models.GQLSchemaRoot) handlers.Instance {
	gqlSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootSchema.Query,
			Mutation: rootSchema.Mutation,
		},
	)

	if err != nil {
		panic(err)
	}

	handler := handlers.Instance{
		GqlSchema: gqlSchema,
	}

	return handler
}
