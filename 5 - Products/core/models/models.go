package models

import "github.com/graphql-go/graphql"

// TokenDetails type.
type TokenDetails struct {
	AccessToken string
}

// AuthCredentials type.
type AuthCredentials struct {
	Email    string
	Password string
}

// User type.
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}

// Product type.
type Product struct {
	ID          string  `json:"id,omitempty"`
	UserID      string  `json:"user_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
}

// RequestErrorStats type.
type RequestErrorStats struct {
	Code    int
	Message string
}

// GQLSchemaRoot type.
type GQLSchemaRoot struct {
	Query    *graphql.Object
	Mutation *graphql.Object
}
