package schemas

import (
	"database/sql"
	"products/core/models"
	"products/data/database/controllers"
	"products/data/graphql/resolvers"
	"products/data/graphql/types"

	"github.com/graphql-go/graphql"
)

var coreRoot models.GQLSchemaRoot

// InitializeCoreRoot initiates schema root, making a setup with all queries and mutations
// that will be used in core route.
func InitializeCoreRoot(db *sql.DB) {
	resolver := resolvers.ProductResolver{DB: &controllers.DBProductInstance{DB: db}}

	coreRoot = models.GQLSchemaRoot{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"products": &graphql.Field{
						Type: graphql.NewList(types.Product),
						Args: graphql.FieldConfigArgument{
							"user_id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.IndexProducts,
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"products": &graphql.Field{
						Type: types.Product,
						Args: graphql.FieldConfigArgument{
							"user_id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"price": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Float),
							},
							"quantity": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.AddProduct,
					},
					"update_product": &graphql.Field{
						Type: types.Product,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"price": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Float),
							},
							"quantity": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.UpdateProduct,
					},
					"delete_product": &graphql.Field{
						Type: graphql.Boolean,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.DeleteProduct,
					},
				},
			},
		),
	}
}

// GetCoreRoot function returns a graphql schema root fot core route.
func GetCoreRoot() models.GQLSchemaRoot {
	return coreRoot
}
