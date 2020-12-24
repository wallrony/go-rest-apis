package types

import "github.com/graphql-go/graphql"

// TODO Graphql type.
var TODO = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TODO",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
