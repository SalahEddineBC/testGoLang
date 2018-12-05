package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"text": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello from GraphQL", nil
				},
			},
		},
	},
)
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("those errors happend: %v", result.Errors)
	}
	return result
}

type queryStruct struct {
	Query string
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			decoder := json.NewDecoder(r.Body)
			var t queryStruct
			decoder.Decode(&t)
			result := executeQuery(t.Query, schema)
			json.NewEncoder(w).Encode(result)
		} else {
			result := executeQuery(r.URL.Query().Get("query"), schema)
			json.NewEncoder(w).Encode(result)
		}
	},
	)

	fmt.Println("Listening On Port 3000")
	http.ListenAndServe(":3000", nil)

}
