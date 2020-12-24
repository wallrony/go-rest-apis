package todocontroller

import (
	"database/sql"
	"fmt"
	"todo/models"
)

// DB type struct
type DB struct {
	*sql.DB
}

// Index function returns all todo that's associated
// with userID.
func (instance *DB) Index() ([]models.TODO, error) {
	stmt, err := instance.Prepare("SELECT * FROM todo;")

	if err != nil {
		err = fmt.Errorf("Error when preparing query to get all todos: %v", err.Error())

		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		err = fmt.Errorf("Error when execute query to get all todos: %v", err.Error())

		return nil, err
	}

	var todo models.TODO

	todoList := []models.TODO{}

	for rows.Next() {
		err = rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.Date,
		)

		if err != nil {
			err = fmt.Errorf("Error scanning data to convert to todo: %v", err.Error())

			return nil, err
		}

		todoList = append(todoList, todo)
	}

	return todoList, nil
}

// Add function add a todo register row in the
// database.
func (instance *DB) Add(data map[string]interface{}) (models.TODO, error) {
	stmt, err := instance.Prepare("INSERT INTO todo(name, date) VALUES($1, $2) RETURNING id;")

	if err != nil {
		err = fmt.Errorf("Error prepating to insert a todo: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	var rows *sql.Rows

	rows, err = stmt.Query(data["name"], data["date"])

	if err != nil {
		err = fmt.Errorf("Error executing insert todo query: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	var id string

	for rows.Next() {
		rows.Scan(&id)
	}

	todo, err := instance.Show(id)

	if err != nil {
		err = fmt.Errorf("Error executing select todo after insert: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	return todo, nil
}

// Show returns a todo with a specified id arg passed.
func (instance *DB) Show(id string) (models.TODO, error) {

	stmt, err := instance.Prepare("SELECT * FROM todo WHERE id=$1")

	if err != nil {
		err = fmt.Errorf("Error preparing to select a todo: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	rows, err := stmt.Query(id)

	if err != nil {
		err = fmt.Errorf("Error executing select todo query: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	var todo models.TODO

	for rows.Next() {
		rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.Date,
		)
	}

	return todo, nil
}

// Update function updates a todo in database with the id passed
// in data var.
func (instance *DB) Update(data map[string]interface{}) (models.TODO, error) {
	stmt, err := instance.Prepare("UPDATE todo SET name=$1, date=$2 WHERE id=$3;")

	if err != nil {
		err = fmt.Errorf("Error preparing update todo: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	_, err = stmt.Query(data["name"], data["date"], data["id"])

	if err != nil {
		err = fmt.Errorf("Error executing update todo query: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	todo, err := instance.Show(data["id"].(string))

	if err != nil {
		err = fmt.Errorf("Error executing select after update: %v", err.Error())

		return models.TODO{ID: "-1"}, err
	}

	return todo, nil
}

// Delete function deletes a todo in database by
// the passed id in data arguments.
func (instance *DB) Delete(data map[string]interface{}) (interface{}, error) {
	stmt, err := instance.Prepare("DELETE FROM todo WHERE id=$1")

	if err != nil {
		err = fmt.Errorf("Error preparing delete todo query: %v", err.Error())

		return nil, err
	}

	_, err = stmt.Query(data["id"])

	if err != nil {
		err = fmt.Errorf("Error executing delete todo query: %v", err.Error())

		return nil, err
	}

	return nil, nil
}
