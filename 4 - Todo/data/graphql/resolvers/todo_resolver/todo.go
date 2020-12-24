package todoresolver

import (
	"errors"
	todocontroller "todo/data/db/todo_controller"

	"github.com/graphql-go/graphql"
)

// ResolverTODO struct serves to reference to
// TodoDBAcoplate.
type ResolverTODO struct {
	Db *todocontroller.DB
}

// Index function returns all todo in database.
func (resolver *ResolverTODO) Index(params graphql.ResolveParams) (interface{}, error) {
	todos, err := resolver.Db.Index()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Show function returns a todo by passing id.
func (resolver *ResolverTODO) Show(params graphql.ResolveParams) (interface{}, error) {
	todo, err := resolver.Db.Show(params.Args["id"].(string))

	if err != nil {
		return nil, err
	} else if todo.ID == "" {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

// Add function send data to todo database controller
// to add a todo in database and return the register.
func (resolver *ResolverTODO) Add(params graphql.ResolveParams) (interface{}, error) {
	todo, err := resolver.Db.Add(params.Args)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Update function send data to todo database controller
// to updates a todo in database with the passed id and
// returns the todo updated.
func (resolver *ResolverTODO) Update(params graphql.ResolveParams) (interface{}, error) {
	todo, err := resolver.Db.Update(params.Args)

	if err != nil {
		return nil, err
	} else if todo.ID == "" {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

// Delete function sends data that need to be used in
// todo database controller to delete a todo.
func (resolver *ResolverTODO) Delete(params graphql.ResolveParams) (interface{}, error) {
	_, err := resolver.Db.Delete(params.Args)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
