package fetch

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/demo-sr-app/api"
	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func SubscribeToAllLinks(client jagw.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To All Links")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	makeCancellable(cancel)

	subscription := &jagw.TopologySubscription{}

	// Subscribe via gRPC
	stream := api.SubscribeToLinks(ctx, client, subscription)

	for {
		event, err := stream.Recv()
		if processEvent(ctx, event, err) {
			// As long as events can be processed successfully, print them
			log.Print("---------- EVENT ----------")
			prettyPrint(event)
		} else {
			// Some error occurred, abort subscription
			break

		}
	}
}

func SubscribeToSpecificLink(client jagw.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific Link")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	makeCancellable(cancel)

	subscription := &jagw.TopologySubscription{
		Keys: []string{
			"2_0_2_0_0000.0000.0007_2001:db8:37::7_0000.0000.0003_2001:db8:37::3",
		},
		Properties: []string{
			"Key",
			"LocalLinkIp",
			"RemoteLinkIp",
		},
	}

	// Subscribe via gRPC
	stream := api.SubscribeToLinks(ctx, client, subscription)

	for {
		event, err := stream.Recv()
		if processEvent(ctx, event, err) {
			// As long as events can be processed successfully, print them
			log.Print("---------- EVENT ----------")
			prettyPrint(event)
		} else {
			// Some error occurred, abort subscription
			break

		}
	}
}

func SubscribeToLsNodeEdges(client jagw.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To LSNodeEdges")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	makeCancellable(cancel)

	subscription := &jagw.TopologySubscription{
		Keys: []string{
			"2_0_2_0_0000.0000.0001_2001:db8:12::1_0000.0000.0002_2001:db8:12::2",
			"2_0_2_0_0000.0000.0001_2001:db8:12::2_0000.0000.0002_2001:db8:12::1",
		},
		Properties: []string{
			"Key",
			"LocalLinkIp",
			"RemoteLinkIp",
		},
	}

	// Subscribe via gRPC
	stream := api.SubscribeToLsNodeEdges(ctx, client, subscription)

	for {
		event, err := stream.Recv()
		if processEvent(ctx, event, err) {
			// As long as events can be processed successfully, print them
			log.Print("---------- EVENT ----------")
			prettyPrint(event)
		} else {
			// Some error occurred, abort subscription
			break

		}
	}
}
