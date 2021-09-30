package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	// "unsafe"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/jalapeno-api-gateway/demo-sr-app/proto/requestservice"
	"github.com/jalapeno-api-gateway/demo-sr-app/proto/subscriptionservice"
)

func main() {
	log.Print("Starting SR-App ...")
	//Connecting to Request Service
	var rsConn *grpc.ClientConn
	rsConn, rsErr := grpc.Dial(os.Getenv("REQUEST_SERVICE_ADDRESS"), grpc.WithInsecure())
	if rsErr != nil {
		log.Fatalf("Could not connect to request service: %s", rsErr)
	}
	defer rsConn.Close()

	//Connecting to Subscription Service
	var psConn *grpc.ClientConn
	psConn, psErr := grpc.Dial(os.Getenv("SUBSCRIPTION_SERVICE_ADDRESS"), grpc.WithInsecure())
	if psErr != nil {
		log.Fatalf("Could not connect to subscription service: %s", psErr)
	}
	defer psConn.Close()

	rsClient := requestservice.NewRequestServiceClient(rsConn)
	psClient := subscriptionservice.NewSubscriptionServiceClient(psConn)

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Press any key to: REQUEST ALL NODES")
	input.Scan()
	GetAllNodes(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST THREE SPECIFIC NODES")
	input.Scan()
	GetSpecificNodes(rsClient)

	fmt.Print("Press 'Enter' to: REQUEST DATA RATES OF SPECIFIC IPV4ADDRESSES")
	input.Scan()
	GetDataRates(rsClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO ALL LINKS")
	input.Scan()
	SubscribeToAllLinks(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO SPECIFIC LINKS")
	input.Scan()
	SubscribeToSpecificLinks(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO DATA RATES OF SPECIFIC IPV4ADDRESSES")
	input.Scan()
	SubscribeToDataRates(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO PACKETS SENT AND RECEIVED OF SPECIFIC IPV4ADDRESSES")
	input.Scan()
	SubscribeToPacketsSentAndReceived(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO EVERYTHING")
	input.Scan()
	SubscribeToEverything(psClient)
}

func GetDataRates(client requestservice.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting Specific DataRates")
	ips := []string{
		"10.18.8.53",
		"10.18.8.54",
		"invalid",
	}
	
	propertyNames := []string{
		"State",
	}

	message := &requestservice.TelemetryRequest{Ipv4Addresses: ips, PropertyNames: propertyNames}
	response, err := client.GetTelemetryData(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetDataRates on RequestService: %s", err)
	}
	printResponse(response)
}

func SubscribeToDataRates(client subscriptionservice.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific DataRates")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ips := []string{"10.18.8.53", "10.18.8.54", "10.1.234.13"}
	propertyNames := []string{
		"DataRate",
	}

	message := &subscriptionservice.TelemetrySubscription{Ipv4Addresses: ips, PropertyNames: propertyNames}
	stream, err := client.SubscribeToTelemetryData(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToDataRates on SubscriptionService: %s", err)
	}

	cancelled := make(chan bool, 1)
	go allowUserToCancel(cancelled)
	go func() {
		<-cancelled
		cancel()
	}()

	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		// ctx.Err != nil if the context was canceled
		if ctx.Err() != nil {
			//client canceled so we exit the loop
			break
		}
		if err != nil {
			log.Fatalf("%v.SubscribeToDataRates(_) = _, %v", client, err)
		}
		printResponse(event)
	}
}

func SubscribeToPacketsSentAndReceived(client subscriptionservice.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To PacketsSent")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ips := []string{"10.18.8.53", "10.18.8.54", "10.1.234.13"}
	propertyNames := []string{
		"PacketsSent",
		"PacketsReceived",
	}
	message := &subscriptionservice.TelemetrySubscription{Ipv4Addresses: ips, PropertyNames: propertyNames}
	stream, err := client.SubscribeToTelemetryData(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToPacketsSentAndReceived on SubscriptionService: %s", err)
	}

	cancelled := make(chan bool, 1)
	go allowUserToCancel(cancelled)
	go func() {
		<-cancelled
		cancel()
	}()

	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		// ctx.Err != nil if the context was canceled
		if ctx.Err() != nil {
			//client canceled so we exit the loop
			break
		}
		if err != nil {
			log.Fatalf("%v.SubscribeToPacketsSentAndReceived(_) = _, %v", client, err)
		}
		printResponse(event)
	}
}

func SubscribeToEverything(client subscriptionservice.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Everything")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	message := &subscriptionservice.TelemetrySubscription{}
	stream, err := client.SubscribeToTelemetryData(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToEverything on SubscriptionService: %s", err)
	}

	cancelled := make(chan bool, 1)
	go allowUserToCancel(cancelled)
	go func() {
		<-cancelled
		cancel()
	}()

	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		// ctx.Err != nil if the context was canceled
		if ctx.Err() != nil {
			//client canceled so we exit the loop
			break
		}
		if err != nil {
			log.Fatalf("%v.SubscribeToEverything(_) = _, %v", client, err)
		}
		printResponse(event)
	}
}

//
//
// -----> SUBSCRIBING TO LINKS <-----
//
//

func SubscribeToSpecificLinks(client subscriptionservice.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific Links")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	keys := []string{
		"2_0_2_0_0000.0000.000b_2001:db8:117::11_0000.0000.0007_2001:db8:117::7",
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

//
//
// -----> REQUESTING NODES <-----
//
//

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

//TODO: Error Handling
func GetSpecificNodes(client requestservice.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting Three Specific Nodes ...")
	keys := []string{
		"2_0_0_0000.0000.000a",
		"2_0_0_0000.0000.0001",
		"2_0_0_0000.0000.0002",
		"invalid",
	}
	propertyNames := []string{
		"RouterIP",
		"Name",
		"ASN",
		"Timestamp",
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
