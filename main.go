package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	rsproto "gitlab.ost.ch/ins/jalapeno-api/sr-app/proto/request-service"
	psproto "gitlab.ost.ch/ins/jalapeno-api/sr-app/proto/push-service"
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

func SubscribeToDataRates(client psproto.PushServiceClient) {
	log.Printf("---------------------Subscribing to DataRates---------------------")
	ips := []string{"10.18.8.53", "10.18.8.54"}
	message := &psproto.IPv4Addresses{Ipv4Address: ips}
	stream, err := client.SubscribeToDataRates(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToDataRates on PushService: %s", err)
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
	log.Printf("--------------------- Subscription ended ---------------------")
}

func GetNodes(client rsproto.ApiGatewayClient) {
	log.Printf("---------------------Requesting Nodes---------------------")
	ids := []string{"2_0_0_0000.0000.000a", "2_0_0_0000.0000.0001", "2_0_0_0000.0000.0002", "2_0_0_0000.0000.0003", "2_0_0_0000.0000.0004", "2_0_0_0000.0000.0005", "2_0_0_0000.0000.0006", "2_0_0_0000.0000.0007", "2_0_0_0000.0000.000c"}
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

	//Connecting to Request Service
	var rsConn *grpc.ClientConn
	rsConn, rsErr := grpc.Dial(os.Getenv("REQUEST_SERVICE_ADDRESS"), grpc.WithInsecure())
	if rsErr != nil {
		log.Fatalf("Could not connect to request service: %s", rsErr)
	}
	defer rsConn.Close()

	//Connecting to Push Service
	var psConn *grpc.ClientConn
	psConn, psErr := grpc.Dial(os.Getenv("PUSH_SERVICE_ADDRESS"), grpc.WithInsecure())
	if psErr != nil {
		log.Fatalf("Could not connect to push service: %s", psErr)
	}
	defer psConn.Close()

	rsClient := rsproto.NewApiGatewayClient(rsConn)
	psClient := psproto.NewPushServiceClient(psConn)

	GetNodes(rsClient)
	GetDataRates(rsClient)
	SubscribeToDataRates(psClient)
}
