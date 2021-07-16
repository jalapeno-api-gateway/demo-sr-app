package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	rsproto "gitlab.ost.ch/ins/jalapeno-api/sr-app/proto"
)

func GetDataRates(client rsproto.ApiGatewayClient) {
	log.Printf("---------------------Requesting DataRates---------------------")
	ips := []string{"10.18.8.53", "10.18.8.54"}
	message := &rsproto.IPv4Addresses{Ipv4Address: ips}
	stream, err := client.GetDataRates(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetDataRates on RequestService: %s", err)
	}
	for {
		dataRate, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetDataRates(_) = _, %v", client, err)
		}
		log.Println(dataRate)
	}
	log.Printf("---------------------All DataRates received---------------------")
}

func GetNodes(client rsproto.ApiGatewayClient) {
	log.Printf("---------------------Requesting Nodes---------------------")
	ids := []string{"2_0_0_0000.0000.000a", "2_0_0_0000.0000.0001", "2_0_0_0000.0000.000c"}
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

func main() {
	log.Print("Starting SR-App ...")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("REQUEST_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := rsproto.NewApiGatewayClient(conn)

	GetNodes(client)
	GetDataRates(client)
}
