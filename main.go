package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jalapeno-api-gateway/demo-sr-app/fetch"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Print("Starting SR-App.")

	// Get Endpoints
	rsEndpoint := getRequestServiceEndpoint()
	ssEndpoint := getSubscriptionServiceEndpoint()

	// Setup Request Service Connection
	var rsConnection *grpc.ClientConn
	rsConnection, rsErr := grpc.Dial(rsEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if rsErr != nil {
		log.Fatalf("Failed to setup request service connection: %s", rsErr)
	}
	defer rsConnection.Close()

	// Setup Subscription Service Connection
	var ssConnection *grpc.ClientConn
	ssConnection, ssErr := grpc.Dial(ssEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if ssErr != nil {
		log.Fatalf("Failed to setup subscription service connection: %s", ssErr)
	}
	defer ssConnection.Close()

	// Create Clients
	rsClient := jagw.NewRequestServiceClient(rsConnection)
	ssClient := jagw.NewSubscriptionServiceClient(ssConnection)

	// Demo requests and subscriptions
	makeTopologyRequests(rsClient)
	makeTelemetryRequests(rsClient)
	makeTopologySubscriptions(ssClient)
	makeTelemetrySubscriptions(ssClient)
}

func getRequestServiceEndpoint() string {
	serverAddress := os.Args[1]
	requestServicePort := os.Args[2]
	return fmt.Sprintf("%s:%s", serverAddress, requestServicePort) // Returns Endpoint as <ip-address>:<port>
}

func getSubscriptionServiceEndpoint() string {
	serverAddress := os.Args[1]
	subscriptionServicePort := os.Args[3]
	return fmt.Sprintf("%s:%s", serverAddress, subscriptionServicePort) // Returns Endpoint as <ip-address>:<port>
}

func makeTopologyRequests(rsClient jagw.RequestServiceClient) {
	// Just for demo purposes:
	// Use as stdin scanner to wait for user input before continuing to next request
	input := bufio.NewScanner(os.Stdin)

	// Make requests
	fmt.Print("Press 'Enter' to: REQUEST ALL NODES")
	input.Scan()
	fetch.GetAllNodes(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST SPECIFIC NODES AND SPECIFIC PROPERTIES")
	input.Scan()
	fetch.GetSpecificNodes(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST ALL LINKS")
	input.Scan()
	fetch.GetAllLinks(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST SEPCIFC LINKS AND SPECIFIC PROPERTIES")
	input.Scan()
	fetch.GetSpecificLink(rsClient)
}

func makeTelemetryRequests(rsClient jagw.RequestServiceClient) {
	// Just for demo purposes:
	// Use as stdin scanner to wait for user input before continuing to next request
	input := bufio.NewScanner(os.Stdin)

	// Make requests
	fmt.Print("Press 'Enter' to: REQUEST DATA RATES OF PAST 60 SECONDS OF SPECIFIC NODE")
	input.Scan()
	fetch.GetDataRatesOfSpecificNode(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST LATEST MEASUREMENT WITHOUT UNFLATTEN")
	input.Scan()
	fetch.GetLatestMeasurement(rsClient, false)

	fmt.Print("Press 'Enter' to: REQUEST LATEST MEASUREMENT WITH UNFLATTEN")
	input.Scan()
	fetch.GetLatestMeasurement(rsClient, true)
}

func makeTopologySubscriptions(ssClient jagw.SubscriptionServiceClient) {
	// Just for demo purposes:
	// Use as stdin scanner to wait for user input before continuing to next request
	input := bufio.NewScanner(os.Stdin)

	// Make subscriptions
	fmt.Print("Press 'Enter' to: SUBSCRIBE TO ALL LINKS")
	input.Scan()
	fetch.SubscribeToAllLinks(ssClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO SPECIFIC LINK")
	input.Scan()
	fetch.SubscribeToSpecificLink(ssClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO LS NODE EDGE")
	input.Scan()
	fetch.SubscribeToLsNodeEdges(ssClient)
}

func makeTelemetrySubscriptions(ssClient jagw.SubscriptionServiceClient) {
	// Just for demo purposes:
	// Use as stdin scanner to wait for user input before continuing to next request
	input := bufio.NewScanner(os.Stdin)

	// Make subscriptions
	fmt.Print("Press 'Enter' to: SUBSCRIBE TO TELEMETRY DATA OF SPECIFIC NODE")
	input.Scan()
	fetch.SubscribeToTelemetryDataOfSpecificNode(ssClient)
}
