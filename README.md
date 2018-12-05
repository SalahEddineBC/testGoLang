# Go GraphQL
this is a very Basic GraphQL server that returns always the same string while queried
# Installation
ensure that you Have GraphQL package for Golang  
``` go get 	"github.com/graphql-go/graphql" ```
# Running the server  
``` go run main.go ```
# Sending a query  
we use the endpoint `/graphQL`  
**using a POST request**  
		the query should be like this, since our server is very limited:  
  ```
  {
  	'query: {text}
  }    
  ```
notice that the header `application/json` should be present in the request  
**using a GET request**  
	the query should be like that:   
	`localhost:3000/graphql?query={text}`
