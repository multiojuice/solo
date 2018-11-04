package controller

import (
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
)

// HandleQuery: To handle every query
func HandleQuery(query string, schema graphql.Schema) []byte {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	rJson, _ := json.Marshal(result)
	return rJson
}
