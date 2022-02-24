package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Get Endpoint as <ip-address>:<port> using the parameters passed by the user
	endpoint := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]) 

	// Setup Request Service Connection
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to setup request service connection: %s", err)
	}
	defer connection.Close()

	// Create Client
	client := jagw.NewRequestServiceClient(connection)

	// Request all Nodes
	request := &jagw.TopologyRequest{
		Keys: []string{},
		Properties: []string{},
	}

	response, err := client.GetLsNodes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsNodes on request service: %s", err)
	}

	prettyPrint(response)
}

func prettyPrint(any interface{}) {
	s, _ := json.MarshalIndent(any, "", "  ")
	fmt.Printf("%s\n\n", string(s))
}
