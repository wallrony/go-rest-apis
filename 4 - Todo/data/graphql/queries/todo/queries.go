package queries

import (
	"database/sql"

	"github.com/graphql-go/graphql"

	todocontroller "todo/data/db/todo_controller"
	todoresolver "todo/data/graphql/resolvers/todo_resolver"
	types "todo/data/graphql/types"
	"todo/models"
)

var root models.GQLRoot

// InitializeRoot function initialize graphql API root
// instance in root var.
func InitializeRoot(db *sql.DB) {
	todoResolver := todoresolver.ResolverTODO{Db: &todocontroller.DB{DB: db}}

	root = models.GQLRoot{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"todo_index": &graphql.Field{
						Type:    graphql.NewList(types.TODO),
						Resolve: todoResolver.Index,
					},
					"todo_show": &graphql.Field{
						Type: types.TODO,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: todoResolver.Show,
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"create_todo": &graphql.Field{
						Type: types.TODO,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"date": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: todoResolver.Add,
					},
					"update_todo": &graphql.Field{
						Type: types.TODO,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"date": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: todoResolver.Update,
					},
					"delete_todo": &graphql.Field{
						Type: types.TODO,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: todoResolver.Delete,
					},
				},
			},
		),
	}
}

// GetRoot function return graphql root instance.
func GetRoot() models.GQLRoot {
	return root
}
