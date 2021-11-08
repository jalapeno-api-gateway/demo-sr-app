package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"time"

	// "io"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	// "google.golang.org/protobuf/proto"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
)

func main() {
	log.Print("Starting SR-App ...")

	// For production
	server := os.Args[1]
	requestServicePort := os.Args[2]
	subscriptionServicePort := os.Args[3]
	requestService := fmt.Sprintf("%s:%s", server, requestServicePort)
	subscriptionService := fmt.Sprintf("%s:%s", server, subscriptionServicePort)

	// For dev env
	// requestService := os.Getenv("REQUEST_SERVICE_ADDRESS")
	// subscriptionService := os.Getenv("SUBSCRIPTION_SERVICE_ADDRESS")

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

	rsClient := jagw.NewRequestServiceClient(rsConn)
	// psClient := jagw.NewSubscriptionServiceClient(psConn)

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Press 'Enter' to: REQUEST ALL NODES")
	input.Scan()
	GetAllNodes(rsClient)

	// fmt.Print("Press 'Enter' to: REQUEST SPECIFIC NODES")
	// input.Scan()
	// GetSpecificNodes(rsClient)

	// fmt.Print("Press 'Enter' to: REQUEST DATA RATES OF SPECIFIC IPV4ADDRESSES")
	// input.Scan()
	// GetDataRates(rsClient)

	// fmt.Print("Press 'Enter' to: SUBSCRIBE TO ALL LINKS")
	// input.Scan()
	// SubscribeToAllLinks(psClient)

	// fmt.Print("Press 'Enter' to: SUBSCRIBE TO SPECIFIC LINK")
	// input.Scan()
	// SubscribeToSpecificLink(psClient)

	// fmt.Print("Press 'Enter' to: SUBSCRIBE TO DATA RATES OF SPECIFIC IPV4ADDRESSES")
	// input.Scan()
	// SubscribeToDataRates(psClient)

	// fmt.Print("Press 'Enter' to: SUBSCRIBE TO PACKETS SENT AND RECEIVED OF SPECIFIC IPV4ADDRESSES")
	// input.Scan()
	// SubscribeToPacketsSentAndReceived(psClient)
	
	// fmt.Print("Press 'Enter' to: REQUEST ALL NODE-EDGES")
	// input.Scan()
	// GetAllLsNodeEdges(rsClient)

	// fmt.Print("Press 'Enter' to: SUBSCRIBE TO LSNODE_EDGES")
	// input.Scan()
	// SubscribeToLsNodeEdges(psClient)
}

// func SubscribeToSpecificLink(client jagw.SubscriptionServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Subscribing To Specific Link")

// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	keys := []string{
// 		"2_0_2_0_0000.0000.0001_2001:db8:12::1_0000.0000.0002_2001:db8:12::2",
// 	}
// 	message := &jagw.TopologySubscription{Keys: keys}
// 	stream, err := client.SubscribeToLsLinks(ctx, message)
// 	if err != nil {
// 		log.Fatalf("Error when calling SubscribeToSpecificLinks on SubscriptionService: %s", err)
// 	}

// 	cancelled := make(chan bool, 1)
// 	go allowUserToCancel(cancelled)
// 	go func() {
// 		<-cancelled
// 		cancel()
// 	}()

// 	for {
// 		link, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			if ctx.Err() != nil {
// 				break
// 			}
// 			log.Fatalf("%v.SubscribeToSpecificLinks(_) = _, %v", client, err)
// 		}
// 		printResponse(link)
// 	}
// }

// func SubscribeToAllLinks(client jagw.SubscriptionServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Subscribing To All Links")

// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	message := &jagw.TopologySubscription{}
// 	stream, err := client.SubscribeToLsLinks(ctx, message)
// 	if err != nil {
// 		log.Fatalf("Error when calling SubscribeToAllLinks on SubscriptionService: %s", err)
// 	}

// 	cancelled := make(chan bool, 1)
// 	go allowUserToCancel(cancelled)
// 	go func() {
// 		<-cancelled
// 		cancel()
// 	}()

// 	for {
// 		link, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			if ctx.Err() != nil {
// 				break
// 			}
// 			log.Fatalf("%v.SubscribeToAllLinks(_) = _, %v", client, err)
// 		}
// 		printResponse(link)
// 	}
// }

// func SubscribeToDataRates(client jagw.SubscriptionServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Subscribing To Specific DataRates")

// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
	
// 	interfaceIds := []*jagw.InterfaceIdentifier{
// 		{Hostname: proto.String("XR-8"), LinkId: proto.Int32(10)},
// 	}

// 	propertyNames := []string{
// 		"DataRate",
// 	}

// 	message := &jagw.TelemetrySubscription{InterfaceIds: interfaceIds, PropertyNames: propertyNames}
// 	stream, err := client.SubscribeToTelemetryData(ctx, message)
// 	if err != nil {
// 		log.Fatalf("Error when calling SubscribeToDataRates on PushService: %s", err)
// 	}

// 	cancelled := make(chan bool, 1)
// 	go allowUserToCancel(cancelled)
// 	go func() {
// 		<-cancelled
// 		cancel()
// 	}()

// 	for {
// 		event, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		// ctx.Err != nil if the context was canceled
// 		if ctx.Err() != nil {
// 			//client canceled so we exit the loop
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("%v.SubscribeToDataRates(_) = _, %v", client, err)
// 		}
// 		printResponse(event)
// 	}
// 	log.Print("--------------------")
// }

// func SubscribeToPacketsSentAndReceived(client jagw.SubscriptionServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Subscribing To PacketsSent And PacketsReceived")

// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	interfaceIds := []*jagw.InterfaceIdentifier{
// 		{Hostname: proto.String("XR-8"), LinkId: proto.Int32(10)},
// 	}

// 	propertyNames := []string{
// 		"PacketsSent",
// 		"PacketsReceived",
// 	}

// 	message := &jagw.TelemetrySubscription{InterfaceIds: interfaceIds, PropertyNames: propertyNames}
// 	stream, err := client.SubscribeToTelemetryData(ctx, message)
// 	if err != nil {
// 		log.Fatalf("Error when calling SubscribeToPacketsSentAndReceived on PushService: %s", err)
// 	}

// 	cancelled := make(chan bool, 1)
// 	go allowUserToCancel(cancelled)
// 	go func() {
// 		<-cancelled
// 		cancel()
// 	}()

// 	for {
// 		event, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		// ctx.Err != nil if the context was canceled
// 		if ctx.Err() != nil {
// 			//client canceled so we exit the loop
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("%v.SubscribeToPacketsSentAndReceived(_) = _, %v", client, err)
// 		}
// 		printResponse(event)
// 	}
// 	log.Print("--------------------")
// }

// func SubscribeToLsNodeEdges(client jagw.SubscriptionServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Subscribing To LsNodeEdges")

// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	message := &jagw.TopologySubscription{}
// 	stream, err := client.SubscribeToLsNodeEdges(ctx, message)
// 	if err != nil {
// 		log.Fatalf("Error when calling SubscribeToLsNodeEdges on PushService: %s", err)
// 	}

// 	cancelled := make(chan bool, 1)
// 	go allowUserToCancel(cancelled)
// 	go func() {
// 		<-cancelled
// 		cancel()
// 	}()

// 	for {
// 		event, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		// ctx.Err != nil if the context was canceled
// 		if ctx.Err() != nil {
// 			//client canceled so we exit the loop
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("%v.SubscribeToLsNodeEdges(_) = _, %v", client, err)
// 		}
// 		printResponse(event)
// 	}
// 	log.Print("--------------------")
// }

// func GetDataRates(client jagw.RequestServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Requesting Specific DataRates")

// 	interfaceIds := []*jagw.InterfaceIdentifier{
// 		{Hostname: proto.String("XR-8"), LinkId: proto.Int32(10)},
// 	}

// 	propertyNames := []string{
// 		"DataRate",
// 	}

// 	message := &jagw.TelemetryRequest{InterfaceIds: interfaceIds, PropertyNames: propertyNames}
// 	response, err := client.GetTelemetryData(context.Background(), message)
// 	if err != nil {
// 		log.Fatalf("Error when calling GetDataRates on RequestService: %s", err)
// 	}
// 	// for _, telemetryData := range response.TelemetryData {
// 	// 	printResponse(telemetryData)
// 	// }
// 	// log.Print("--------------------")
// 	printResponse(response)
// }

func GetAllNodes(client jagw.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting All Available Nodes")

	message := &jagw.TelemetryRequest{
		SensorPath: proto.String("Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface"),
		Properties: []string{
			"data_rates/output_data_rate",
		},
		StringFilters: []*jagw.StringFilter{
			{
				Property: proto.String("source"),
				Value: proto.String("XR-8"),
				Operator: jagw.StringOperator_EQUAL.Enum(),
			},
		},
		RangeFilter: &jagw.RangeFilter{
			EarliestTimestamp: proto.Int64(time.Now().Add(-1 * time.Minute).UnixNano()),
		},
	}

	response, err := client.GetTelemetryData(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	}

	printResponse(response)

	for _, t := range response.TelemetryData {
		var myTelemetryData MyTelemetryData
		json.Unmarshal([]byte(t), &myTelemetryData)
		printResponse(myTelemetryData)
	}

}

type MyTelemetryData struct {
	Time string
	// IpInformation_IpAddress string
	DataRates_OutputDataRate int
}

// func GetSpecificNodes(client jagw.RequestServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Requesting Name and ASN of Two Specific Nodes ...")

// 	keys := []string{
// 		"2_0_0_0000.0000.0001",
// 		"2_0_0_0000.0000.0002",
// 	}
// 	propertyNames := []string{
// 		"Name",
// 		"ASN",
// 	}
// 	message := &jagw.TopologyRequest{Keys: keys, PropertyNames: propertyNames}
// 	response, err := client.GetLsNodes(context.Background(), message)
// 	if err != nil {
// 		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
// 	}

// 	printResponse(response)
// }

// func GetAllLsNodeEdges(client jagw.RequestServiceClient) {
// 	log.Print("--------------------")
// 	log.Printf("Requesting All Available LsNodeEdges")

// 	message := &jagw.TopologyRequest{}
// 	response, err := client.GetLsNodeEdges(context.Background(), message)
// 	if err != nil {
// 		log.Fatalf("Error when calling GetAllLsNodeEdges on RequestService: %s", err)
// 	}

// 	printResponse(response)
// }

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
