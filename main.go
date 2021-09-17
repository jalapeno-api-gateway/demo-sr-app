package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	// "unsafe"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gitlab.ost.ch/ins/jalapeno-api/sr-app/proto/pushservice"
	"gitlab.ost.ch/ins/jalapeno-api/sr-app/proto/requestservice"
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

	//Connecting to Push Service
	var psConn *grpc.ClientConn
	psConn, psErr := grpc.Dial(os.Getenv("PUSH_SERVICE_ADDRESS"), grpc.WithInsecure())
	if psErr != nil {
		log.Fatalf("Could not connect to push service: %s", psErr)
	}
	defer psConn.Close()

	rsClient := requestservice.NewRequestServiceClient(rsConn)
	psClient := pushservice.NewPushServiceClient(psConn)

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
		"DataRate",
	}

	message := &requestservice.TelemetryRequest{Ipv4Addresses: ips, PropertyNames: propertyNames}
	stream, err := client.GetTelemetryData(context.Background(), message)
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
		printDataRate(dataRate)
	}
	log.Print("--------------------")
}

func SubscribeToDataRates(client pushservice.PushServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific DataRates")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ips := []string{"10.18.8.53", "10.18.8.54", "10.1.234.13"}
	propertyNames := []string{
		"DataRate",
	}

	message := &pushservice.TelemetrySubscription{Ipv4Addresses: ips, PropertyNames: propertyNames}
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
		printDataRateFromPushService(event.Ipv4Address, event.DataRate)
	}
	log.Print("--------------------")
}

func SubscribeToPacketsSentAndReceived(client pushservice.PushServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To PacketsSent")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ips := []string{"10.18.8.53", "10.18.8.54", "10.1.234.13"}
	propertyNames := []string{
		"PacketsSent",
		"PacketsReceived",
	}
	message := &pushservice.TelemetrySubscription{Ipv4Addresses: ips, PropertyNames: propertyNames}
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
		printPacketsSent(event.Ipv4Address, event.PacketsSent)
		printPacketsReceived(event.Ipv4Address, event.PacketsReceived)
	}
	log.Print("--------------------")
}

func SubscribeToEverything(client pushservice.PushServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Everything")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	message := &pushservice.TelemetrySubscription{}
	stream, err := client.SubscribeToTelemetryData(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToEverything on PushService: %s", err)
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
		printPacketsSent(event.Ipv4Address, event.PacketsSent)
		printPacketsReceived(event.Ipv4Address, event.PacketsReceived)
		printDataRateFromPushService(event.Ipv4Address, event.DataRate)
		printState(event.Ipv4Address, event.State)
		printLastStateTransitionTime(event.Ipv4Address, event.LastStateTransitionTime)
	}
	log.Print("--------------------")
}

//
//
// -----> SUBSCRIBING TO LINKS <-----
//
//

func SubscribeToSpecificLinks(client pushservice.PushServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific Links")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	keys := []string{
		"2_0_2_0_0000.0000.000b_2001:db8:117::11_0000.0000.0007_2001:db8:117::7",
	}
	message := &pushservice.TopologySubscription{Keys: keys}
	stream, err := client.SubscribeToLsLinks(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToSpecificLinks on PushService: %s", err)
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
		printLinkEvent(link)
	}
	log.Print("--------------------")
}

func SubscribeToAllLinks(client pushservice.PushServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To All Links")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	message := &pushservice.TopologySubscription{}
	stream, err := client.SubscribeToLsLinks(ctx, message)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToAllLinks on PushService: %s", err)
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
		printLinkEvent(link)
	}
	log.Print("--------------------")
}

//
//
// -----> REQUESTING NODES <-----
//
//

func GetAllNodes(client requestservice.RequestServiceClient) {
	log.Print("--------------------")
	log.Printf("Requesting All Available Nodes")
	propertyNames := []string{
		"Name",
		"Asn",
		"RouterIp",
	}
	message := &requestservice.TopologyRequest{PropertyNames: propertyNames}
	stream, err := client.GetLsNodes(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	}

	for {
		node, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetLsNodes(_) = _, %v", client, err)
		}
		printNode(node)
	}
	log.Print("--------------------")
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
		"Asn",
		"RouterIp",
	}
	message := &requestservice.TopologyRequest{Keys: keys, PropertyNames: propertyNames}
	stream, err := client.GetLsNodes(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	}

	for {
		node, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetLsNodes(_) = _, %v", client, err)
		}
		printNode(node)
	}
	log.Print("--------------------")
}

//
//
// -----> PRINTERS <-----
//
//

func printNode(node *requestservice.LsNode) {
	log.Printf(">>> Received Node \"%s\"\n", node.Name)
	log.Printf("  Key: %s", node.Key)
	log.Printf("  Name: %s", node.Name)
	log.Printf("  Asn: %d", node.Asn)
	log.Printf("  RouterIp: %s", node.RouterIp)
}

func printLink(link *requestservice.LsLink) {
	log.Printf(">>> Received Link\n")
	log.Printf("  Key: %s", link.Key)
	log.Printf("  RouterIp: %s", link.RouterIp)
	log.Printf("  PeerIp: %s", link.PeerIp)
	log.Printf("  LocalLinkIp: %s", link.LocalLinkIp)
	log.Printf("  RemoteLinkIp: %s", link.RemoteLinkIp)
	log.Printf("  IgpMetric: %d", link.IgpMetric)
}

func printLinkEvent(event *pushservice.LsLinkEvent) {
	log.Printf(">>> Received LinkEvent\n")
	log.Printf("  Action: %s", event.Action)
	log.Printf("  Key: %s", event.LsLink.Key)
	log.Printf("  RouterIp: %s", event.LsLink.RouterIp)
	log.Printf("  PeerIp: %s", event.LsLink.PeerIp)
	log.Printf("  LocalLinkIp: %s", event.LsLink.LocalLinkIp)
	log.Printf("  RemoteLinkIp: %s", event.LsLink.RemoteLinkIp)
	log.Printf("  IgpMetric: %d", event.LsLink.IgpMetric)
}

func printDataRate(response *requestservice.TelemetryResponse) {
	log.Printf(">>> Received DataRate\n")
	log.Printf("  Ipv4Address: %s", response.Ipv4Address)
	log.Printf("  DataRate: %d", response.DataRate)
}

func printDataRateFromPushService(ip string, dataRate int64) {
	log.Printf(">>> Received DataRate\n")
	log.Printf("  Ipv4Address: %s", ip)
	log.Printf("  DataRate: %d", dataRate)
}

func printPacketsSent(ip string, packetsSent int64) {
	log.Printf(">>> Received PacketsSent\n")
	log.Printf("  Ipv4Address: %s", ip)
	log.Printf("  PacketsSent: %d", packetsSent)
}

func printPacketsReceived(ip string, packetsReceived int64) {
	log.Printf(">>> Received PacketsReceived\n")
	log.Printf("  Ipv4Address: %s", ip)
	log.Printf("  PacketsReceived: %d", packetsReceived)
}

func printState(ip string, state string) {
	log.Printf(">>> Received State\n")
	log.Printf("  Ipv4Address: %s", ip)
	log.Printf("  State: %s", state)
}

func printLastStateTransitionTime(ip string, lastStateTransitionTime int64) {
	log.Printf(">>> Received LastStateTransitionTime\n")
	log.Printf("  Ipv4Address: %s", ip)
	log.Printf("  LastStateTransitionTime: %d", lastStateTransitionTime)
}

func allowUserToCancel(cancelled chan bool) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Press 'Enter' to cancel subscription\n")
	input.Scan()
	cancelled <- true
}
