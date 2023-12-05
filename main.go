package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jalapeno-api-gateway/demo-sr-app/fetch"
	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func main() {
	log.Print("Starting SR-App.")
	rsPort, err := strconv.ParseInt(os.Args[2], 0, 16)
	if err != nil {
		log.Fatalf("Failed to parse port: %s", err)
	}
	ssPort, err := strconv.ParseInt(os.Args[3], 0, 16)
	if err != nil {
		log.Fatalf("Failed to parse port: %s", err)
	}

	// Get Endpoints
	rsEndpoint := jagw.JagwEndpoint{
		EndpointAddress: os.Args[1],
		EndpointPort:    uint16(rsPort),
	}
	ssEndpoint := jagw.JagwEndpoint{
		EndpointAddress: os.Args[1],
		EndpointPort:    uint16(ssPort),
	}

	// Setup Request Service Connection
	rsConnection, rsErr := jagw.NewJagwConnection(rsEndpoint)
	if rsErr != nil {
		log.Fatalf("Failed to setup request service connection: %s", rsErr)
	}
	defer rsConnection.Close()

	// Setup Subscription Service Connection
	ssConnection, ssErr := jagw.NewJagwConnection(ssEndpoint)
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

	fmt.Print("Is the collection LS_NODE_COORDINATES configured? (y/n): ")
	var answer string
	_, err := fmt.Scanln(&answer)
	if err == nil && answer == "Y" || err == nil && answer == "y" {
		fmt.Print("Press 'Enter' to: REQUEST LS_NODE_COORDINATES")
		input.Scan()
		fetch.GetCoordinates(rsClient)
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Skipping LS_NODE_COORDINATES request.")
	}
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
