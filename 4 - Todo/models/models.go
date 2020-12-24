package models

import "github.com/graphql-go/graphql"

// GQLRoot Type Struct
type GQLRoot struct {
	Query    *graphql.Object
	Mutation *graphql.Object
}

// TODO type struct
type TODO struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Date string `json:"date,omitempty"`
}
