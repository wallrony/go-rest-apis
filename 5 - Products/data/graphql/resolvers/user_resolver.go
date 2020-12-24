package resolvers

import (
	"errors"
	"products/core/models"
	"products/core/utils"
	"products/data/database/controllers"

	"github.com/graphql-go/graphql"
)

// UserResolver type.
type UserResolver struct {
	DB *controllers.DBUserInstance
}

// Login function verifys if a user with the provided credentials exists.
func (resolver *UserResolver) Login(params graphql.ResolveParams) (interface{}, error) {
	credentials := models.AuthCredentials{
		Email:    params.Args["email"].(string),
		Password: params.Args["password"].(string),
	}

	if credentials.Email == "" || credentials.Password == "" {
		return nil, errors.New("missing-credentials")
	}

	user, err := resolver.DB.Login(credentials)

	if err != nil {
		return nil, err
	}

	token, err := utils.CreateToken(user.ID)

	if err != nil {
		return nil, errors.New("internal-error")
	}

	data := map[string]interface{}{
		"auth_token": token.AccessToken,
		"user":       user,
	}

	return data, nil
}

// ShowUser function returns a user that corresponds with the provided id
// in arguments.
func (resolver *UserResolver) ShowUser(params graphql.ResolveParams) (interface{}, error) {
	user, err := resolver.DB.ShowUser(params.Args["id"].(string))

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Register function add a new user in database with the provided data in
// arguments.
func (resolver *UserResolver) Register(params graphql.ResolveParams) (interface{}, error) {
	result, err := resolver.DB.AddUser(params.Args)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUser function updates user data in database that corresponds with
// the provided id in arguments.
func (resolver *UserResolver) UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	user, err := resolver.DB.UpdateUser(params.Args)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserPassword function updates user password in database that
// corresponds with the provided id in arguments.
func (resolver *UserResolver) UpdateUserPassword(params graphql.ResolveParams) (interface{}, error) {
	result, err := resolver.DB.UpdateUserPassword(params.Args)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// InactiveUser function updates an user is_active value to false that
// the id corresponds with the provided id in the arguments.
func (resolver *UserResolver) InactiveUser(params graphql.ResolveParams) (interface{}, error) {
	result, err := resolver.DB.InactiveUser(params.Args["id"].(string))

	if err != nil {
		return nil, err
	}

	return result, nil
}

// ActiveUser function updates an user is_active value to true that
// the id corresponds with the provided id in the arguments.
func (resolver *UserResolver) ActiveUser(params graphql.ResolveParams) (interface{}, error) {
	result, err := resolver.DB.ActiveUser(params.Args["id"].(string))

	if err != nil {
		return nil, err
	}

	return result, nil
}
