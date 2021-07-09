package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	rsproto "gitlab.ost.ch/ins/jalapeno-api/sr-app/proto"
)

func main() {
	log.Print("Starting SR-App ...")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("REQUEST_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := rsproto.NewApiGatewayClient(conn)
	ids := []int32{1, 2, 5, 10}
	message := &rsproto.NodeIds{Ids: ids}
	stream, err := client.GetNodes(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	}

	for {
		node, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetNodes(_) = _, %v", client, err)
		}
		log.Println(node)
	}
	log.Printf("---------------------All Nodes received---------------------")
}
