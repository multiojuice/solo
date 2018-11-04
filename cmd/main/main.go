package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/multiojuice/solo/pkg/controller"
	"github.com/multiojuice/solo/pkg/schema"
)

func main() {

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		mainSchema := schema.GetSchema()
		body, _ := ioutil.ReadAll(r.Body)
		result := controller.HandleQuery(string(body), mainSchema)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("")
		fmt.Fprint(w, string(result))
	})

	http.ListenAndServe(":8080", nil)
}
