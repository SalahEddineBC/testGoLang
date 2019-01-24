package graphql

import (
	"github.com/SalahEddineBC/gRPCservice/client"
)

// Server contains
type Server struct {
	Client *client.Client
}

// NewGraphQLServer creates new GraphQL server
func NewGraphQLServer(serviceURL string) (*Server, error) {
	// Connect to account service
	C, err := client.NewClient(serviceURL)
	if err != nil {
		return nil, err
	}
	return &Server{
		C}, nil
}
