package schema

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/graphql-go/graphql"
)

var ACCESS_TOKEN string = "0c9de60ed26319d172042037ae22195e"
var URL string = "https://api.vimeo.com/"

func GetSchema() graphql.Schema {
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
				bodyBytes, _ := ioutil.ReadAll(res.Body)
				return string(bodyBytes), err
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}