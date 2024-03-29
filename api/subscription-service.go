package api

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func SubscribeToLinks(ctx context.Context, client jagw.SubscriptionServiceClient, subscription *jagw.TopologySubscription) jagw.SubscriptionService_SubscribeToLsLinksClient {
	stream, err := client.SubscribeToLsLinks(ctx, subscription)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToLsLinks on SubscriptionService: %s", err)
	}
	return stream
}

func SubscribeToLsNodeEdges(ctx context.Context, client jagw.SubscriptionServiceClient, subscription *jagw.TopologySubscription) jagw.SubscriptionService_SubscribeToLsNodeEdgesClient {
	stream, err := client.SubscribeToLsNodeEdges(ctx, subscription)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToLsNodeEdges on SubscriptionService: %s", err)
	}
	return stream
}

func SubscribeToTelemetryData(ctx context.Context, client jagw.SubscriptionServiceClient, subscription *jagw.TelemetrySubscription) jagw.SubscriptionService_SubscribeToTelemetryDataClient {
	stream, err := client.SubscribeToTelemetryData(ctx, subscription)
	if err != nil {
		log.Fatalf("Error when calling SubscribeToTelemetryData on SubscriptionService: %s", err)
	}
	return stream
}
