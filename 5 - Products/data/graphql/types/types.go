package types

import "github.com/graphql-go/graphql"

// User graphql Type. Object to be identified in a
// GraphQL Query or Mutation.
var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Product graphql Type. Object to be identified in a
// GraphQL Query or Mutation.
var Product = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.String,
			},
			"quantity": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// AuthenticatedData graphql Type. Object to be identified in a
// GraphQL Query or Mutation.
var AuthenticatedData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthenticatedData",
		Fields: graphql.Fields{
			"auth_token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: User,
			},
		},
	},
)
