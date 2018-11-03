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
		// w.Header().Set("")

		data := result.Data.(map[string]interface{})
		for _, v := range data {
			fmt.Fprint(w, v.(string))
		}
	})

	http.ListenAndServe(":8080", nil)
}
