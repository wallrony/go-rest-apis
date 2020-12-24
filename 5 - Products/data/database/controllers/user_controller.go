package controllers

import (
	"database/sql"
	"errors"
	"products/core/models"
)

// DBUserInstance to userController.
type DBUserInstance struct {
	DB *sql.DB
}

// Login function search a user with the provided
// credentials and return a bool if the user exists or not.
func (instance *DBUserInstance) Login(credentials models.AuthCredentials) (*models.User, error) {
	stmt, err := instance.DB.Prepare("SELECT * FROM users WHERE email=$1 AND password=$2;")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(credentials.Email, credentials.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{}

	for rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.IsActive,
		)

		break
	}

	if user.ID == "" {
		return nil, errors.New("invalid-credentials")
	} else if !user.IsActive {
		return nil, errors.New("inactive-user")
	}

	return user, nil
}

// ShowUser function returns the data of a user with
// the correspondent id.
func (instance *DBUserInstance) ShowUser(userID string) (*models.User, error) {
	stmt, err := instance.DB.Prepare("SELECT * FROM users WHERE id=$1 AND is_active=1;")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)

	if err != nil {
		return nil, err
	}

	user := &models.User{}

	for rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.IsActive,
		)
	}

	if user.ID == "" {
		return nil, errors.New("not-found")
	}

	return user, nil
}

// AddUser function register a user in database with
// the arguments data.
func (instance *DBUserInstance) AddUser(data map[string]interface{}) (bool, error) {
	stmt, err := instance.DB.Prepare("INSERT INTO users(name, email, password, is_active) VALUES($1, $2, $3, $4) RETURNING id;")

	if err != nil {
		return false, err
	}

	rows, err := stmt.Query(
		data["name"],
		data["email"],
		data["password"],
		true,
	)

	if err != nil {
		return false, err
	}

	var id string

	for rows.Next() {
		rows.Scan(
			&id,
		)
	}

	if id == "" {
		return false, errors.New("internal-error")
	}

	return true, nil
}

// UpdateUser function updates user data in database that
// corresponds the passed id.
func (instance *DBUserInstance) UpdateUser(data map[string]interface{}) (*models.User, error) {
	stmt, err := instance.DB.Prepare("UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING *;")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(
		data["name"],
		data["email"],
		data["id"],
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{}

	for rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.IsActive,
		)
	}

	if user.ID == "" {
		return nil, errors.New("not-found")
	}

	return user, nil
}

// UpdateUserPassword updates the user password that
// corresponds with the passed id.
func (instance *DBUserInstance) UpdateUserPassword(data map[string]interface{}) (bool, error) {
	stmt, err := instance.DB.Prepare("UPDATE users SET password=$1 WHERE id=$2 RETURNING id;")

	if err != nil {
		return false, err
	}

	rows, err := stmt.Query(data["password"], data["id"])

	if err != nil {
		return false, err
	}

	var id string

	for rows.Next() {
		rows.Scan(&id)
	}

	if id == "" {
		return false, errors.New("not-found")
	}

	return true, err
}

// InactiveUser function inactive a user by setting is_active false in
// database.
func (instance *DBUserInstance) InactiveUser(id string) (bool, error) {
	stmt, err := instance.DB.Prepare("UPDATE users SET is_active=false WHERE id=$1;")

	if err != nil {
		return false, err
	}

	_, err = stmt.Query(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

// ActiveUser function active a user by setting is_active true in
// database.
func (instance *DBUserInstance) ActiveUser(id string) (bool, error) {
	stmt, err := instance.DB.Prepare("UPDATE users SET is_active=true WHERE id=$1;")

	if err != nil {
		return false, err
	}

	_, err = stmt.Query(id)

	if err != nil {
		return false, err
	}

	return true, nil
}
