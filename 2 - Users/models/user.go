package models

import (
	"strconv"
)

// User type to use like a object model
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// UserConstruct Function is to create an User with map
func UserConstruct(data map[string]string) User {
	id, _ := strconv.ParseInt(data["id"], 10, 64)
	age, _ := strconv.ParseInt(data["age"], 10, 64)

	u := User{
		ID:   id,
		Age:  age,
		Name: data["name"],
	}

	return u
}
