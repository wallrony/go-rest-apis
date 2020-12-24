package schemas

import (
	"database/sql"
	"products/core/models"
	"products/data/database/controllers"
	"products/data/graphql/resolvers"
	"products/data/graphql/types"

	"github.com/graphql-go/graphql"
)

var authRoot models.GQLSchemaRoot

// InitializeAuthRoot initiates schema root, making a setup with all queries and mutations
// that will be used in accounts route.
func InitializeAuthRoot(db *sql.DB) {
	resolver := resolvers.UserResolver{DB: &controllers.DBUserInstance{DB: db}}

	authRoot = models.GQLSchemaRoot{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"login": &graphql.Field{
						Type: types.AuthenticatedData,
						Args: graphql.FieldConfigArgument{
							"email": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.Login,
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"register": &graphql.Field{
						Type: graphql.Boolean,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"email": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.Register,
					},
				},
			},
		),
	}
}

// GetAuthRoot function returns a graphql schema root fot accounts route.
func GetAuthRoot() models.GQLSchemaRoot {
	return authRoot
}
