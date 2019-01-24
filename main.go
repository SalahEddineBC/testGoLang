package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	grap "./graphql"
	api "github.com/SalahEddineBC/gRPCservice/api"
	"github.com/graphql-go/graphql"
)

// here we declare our query structure
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"text": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// this is the function excuted when we pass this query
					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					response, err := serv.Client.Client.GetReleaseNotes(ctx, &api.Empty{})
					if err != nil {
						return nil, err
					}
					return response, nil
				},
			},
		},
	},
)

//graphql schema that contains our type
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
var serv grap.Server

// function to handle graphql queries
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
	s, err := grap.NewGraphQLServer("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	serv = *s
	_ = serv
	//we create an endpoint here to listen at /graphql
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		//if we send a post request
		if r.Method == "POST" {
			//we extract the  struct that contains our query
			decoder := json.NewDecoder(r.Body)
			var t queryStruct
			decoder.Decode(&t)
			//we excute the query stored in the struct
			result := executeQuery(t.Query, schema)
			// we send back the result
			json.NewEncoder(w).Encode(result)
		} else {
			//if we send a GET request (or anyother type) we Handle the query from route parameters

			result := executeQuery(r.URL.Query().Get("query"), schema)
			// we send back the result
			json.NewEncoder(w).Encode(result)
		}
	},
	)
	// we start the server to listen on port 3000
	fmt.Println("Listening On Port 3000")
	http.ListenAndServe(":3000", nil)

}
