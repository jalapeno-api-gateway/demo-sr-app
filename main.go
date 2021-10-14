package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/jalapeno-api-gateway/demo-sr-app/proto/requestservice"
	"github.com/jalapeno-api-gateway/demo-sr-app/proto/subscriptionservice"
)

func main() {
	log.Print("Starting SR-App ...")

	server := os.Args[1]
	requestServicePort := os.Args[2]
	subscriptionServicePort := os.Args[3]
	requestService := fmt.Sprintf("%s:%s", server, requestServicePort)
	subscriptionService := fmt.Sprintf("%s:%s", server, subscriptionServicePort)

	//Connecting to Request Service
	var rsConn *grpc.ClientConn
	rsConn, rsErr := grpc.Dial(requestService, grpc.WithInsecure())
	if rsErr != nil {
		log.Fatalf("Could not connect to request service: %s", rsErr)
	}
	defer rsConn.Close()

	//Connecting to Subscription Service
	var psConn *grpc.ClientConn
	psConn, psErr := grpc.Dial(subscriptionService, grpc.WithInsecure())
	if psErr != nil {
		log.Fatalf("Could not connect to subscription service: %s", psErr)
	}
	defer psConn.Close()

	rsClient := requestservice.NewRequestServiceClient(rsConn)
	psClient := subscriptionservice.NewSubscriptionServiceClient(psConn)

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Press 'Enter' to: REQUEST ALL NODES")
	input.Scan()
	GetAllNodes(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST SPECIFIC NODES")
	input.Scan()
	GetSpecificNodes(rsClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO ALL LINKS")
	input.Scan()
	SubscribeToAllLinks(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO SPECIFIC LINK")
	input.Scan()
	SubscribeToSpecificLink(psClient)
}

func SubscribeToSpecificLink(client subscriptionservice.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific Link")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	keys := []string{
		"2_0_2_0_0000.0000.0001_2001:db8:12::1_0000.0000.0002_2001:db8:12::2",
	}
	message := &subscriptionservice.TopologySubscription{Keys: keys}
	stream, err := client.SubscribeToLSLinks(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToSpecificLinks on SubscriptionService: %s", err)
	}

	cancelled := make(chan bool, 1)
	go allowUserToCancel(cancelled)
	go func() {
		<-cancelled
		cancel()
	}()

	for {
		link, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			log.Fatalf("%v.SubscribeToSpecificLinks(_) = _, %v", client, err)
		}
		printResponse(link)
	}
}

func SubscribeToAllLinks(client subscriptionservice.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To All Links")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	message := &subscriptionservice.TopologySubscription{}
	stream, err := client.SubscribeToLSLinks(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToAllLinks on SubscriptionService: %s", err)
	}

	cancelled := make(chan bool, 1)
	go allowUserToCancel(cancelled)
	go func() {
		<-cancelled
		cancel()
	}()

	for {
		link, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			log.Fatalf("%v.SubscribeToAllLinks(_) = _, %v", client, err)
		}
		printResponse(link)
	}
}

func GetAllNodes(client requestservice.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting All Available Nodes")

	message := &requestservice.TopologyRequest{}
	response, err := client.GetLSNodes(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	}

	printResponse(response)
}

func GetSpecificNodes(client requestservice.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting Name and ASN of Two Specific Nodes ...")

	keys := []string{
		"2_0_0_0000.0000.0001",
		"2_0_0_0000.0000.0002",
	}
	propertyNames := []string{
		"Name",
		"ASN",
	}
	message := &requestservice.TopologyRequest{Keys: keys, PropertyNames: propertyNames}
	response, err := client.GetLSNodes(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	}

	printResponse(response)
}

func printResponse(response interface{}) {
	log.Print("---------- RESPONSE ----------")
	s, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("%s\n\n", string(s))
}

func allowUserToCancel(cancelled chan bool) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Press 'Enter' to cancel subscription\n")
	input.Scan()
	cancelled <- true
}
