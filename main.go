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

	fmt.Print("Press 'Enter' to: REQUEST DATA RATES OF SPECIFIC IPV4ADDRESSES")
	input.Scan()
	GetDataRates(rsClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO ALL LINKS")
	input.Scan()
	SubscribeToAllLinks(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO SPECIFIC LINK")
	input.Scan()
	SubscribeToSpecificLink(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO DATA RATES OF SPECIFIC IPV4ADDRESSES")
	input.Scan()
	SubscribeToDataRates(psClient)

	fmt.Print("Press 'Enter' to: SUBSCRIBE TO PACKETS SENT AND RECEIVED OF SPECIFIC IPV4ADDRESSES")
	input.Scan()
	SubscribeToPacketsSentAndReceived(psClient)
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
		log.Fatalf("Error when calling SubscribeToDataRates on PushService: %s", err)
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
	log.Print("--------------------")
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
		log.Fatalf("Error when calling SubscribeToPacketsSentAndReceived on PushService: %s", err)
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
	log.Print("--------------------")
}

func GetDataRates(client requestservice.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting Specific DataRates")

	ips := []string{
		// "10.18.8.53",
		// "10.18.8.54",
		// "invalid",
		"10.18.8.41",
	}

	propertyNames := []string{
		"DataRate",
	}

	message := &requestservice.TelemetryRequest{Ipv4Addresses: ips, PropertyNames: propertyNames}
	response, err := client.GetTelemetryData(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetDataRates on RequestService: %s", err)
	}
	// for _, telemetryData := range response.TelemetryData {
	// 	printResponse(telemetryData)
	// }
	// log.Print("--------------------")
	printResponse(response)
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
