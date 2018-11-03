package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var ACCESS_TOKEN string = "0c9de60ed26319d172042037ae22195e"
var URL string = "https://api.vimeo.com/"

func TestFunction() {
	// Schema
	fields := graphql.Fields{
		"video": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Printf("Hi %s", p.Args["title"].(string))
				url := URL + "videos?query=" + p.Args["title"].(string) + "&per_page=4"
				client := &http.Client{}
				req, _ := http.NewRequest("GET", url, nil)
				req.Header.Set("Authorization", "Bearer "+ACCESS_TOKEN)
				req.Header.Set("Accept", "application/json")
				res, err := client.Do(req)
				buf := new(bytes.Buffer)
				buf.ReadFrom(res.Body)
				newStr := buf.String()

				return newStr, err
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			video(title: "test")
		}`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	fmt.Printf("%s \n", r) // {“data”:{“hello”:”world”}}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
