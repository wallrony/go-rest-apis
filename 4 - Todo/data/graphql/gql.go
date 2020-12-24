package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// ExecuteQuery function returns the result and possible
// error in a query execution.
func ExecuteQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf(result.Errors[0].Error())
	}

	return result, nil
}
