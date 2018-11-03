package main

import (
	"encoding/json"
	"net/http"

	"github.com/multiojuice/solo/pkg/controller"
	"github.com/multiojuice/solo/pkg/schema"
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		schema := schema.GetSchema()
		result := controller.HandleQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
