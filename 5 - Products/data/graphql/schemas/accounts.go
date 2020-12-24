package schemas

import (
	"database/sql"
	"products/core/models"
	"products/data/database/controllers"
	"products/data/graphql/resolvers"
	"products/data/graphql/types"

	"github.com/graphql-go/graphql"
)

var accountsRoot models.GQLSchemaRoot

// InitializeAccountsRoot initiates schema root, making a setup with all queries and mutations
// that will be used in accounts route.
func InitializeAccountsRoot(db *sql.DB) {
	resolver := resolvers.UserResolver{DB: &controllers.DBUserInstance{DB: db}}

	accountsRoot = models.GQLSchemaRoot{
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
					"users": &graphql.Field{
						Type: types.User,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.ShowUser,
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"register": &graphql.Field{
						Type: types.User,
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
					"update_user": &graphql.Field{
						Type: types.User,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"email": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.UpdateUser,
					},
					"update_user_password": &graphql.Field{
						Type: graphql.Boolean,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.UpdateUserPassword,
					},
					"inactive_user": &graphql.Field{
						Type: graphql.Boolean,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.InactiveUser,
					},
					"active_user": &graphql.Field{
						Type: graphql.Boolean,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.ActiveUser,
					},
				},
			},
		),
	}
}

// GetAccountsRoot function returns a graphql schema root fot accounts route.
func GetAccountsRoot() models.GQLSchemaRoot {
	return accountsRoot
}
