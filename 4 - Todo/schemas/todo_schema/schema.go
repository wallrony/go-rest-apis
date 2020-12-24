package todoschema

import (
	"database/sql"
	queries "todo/data/graphql/queries/todo"
	server "todo/data/server"

	"github.com/graphql-go/graphql"
)

// GetSchema function returns todo GraphQL schema.
func GetSchema(instance *sql.DB) (server.Instance, error) {
	queries.InitializeRoot(instance)

	mRoot := queries.GetRoot()

	mSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    mRoot.Query,
			Mutation: mRoot.Mutation,
		},
	)

	if err != nil {
		return server.Instance{}, err
	}

	gqlServer := server.Instance{
		GqlSchema: mSchema,
	}

	return gqlServer, err
}
