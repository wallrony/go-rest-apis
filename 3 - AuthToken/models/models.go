package models

// User Struct is a model to construct an User
type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omit"`
}

// TokenDetails Struct is a model to construct an object with tokens
type TokenDetails struct {
	AccessToken string
	AccessUUID  string
	AtExpires   int64
}
