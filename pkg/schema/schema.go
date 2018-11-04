package schema

import (
	"log"
	"github.com/graphql-go/graphql"
	"github.com/multiojuice/solo/internal/vimeo"
)



func GetSchema() graphql.Schema {

	schemaConfig := graphql.SchemaConfig{Query: vimeo.GetVimeo()}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}
