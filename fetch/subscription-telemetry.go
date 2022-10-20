package fetch

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/demo-sr-app/api"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/protobuf/proto"
)

func SubscribeToTelemetryDataOfSpecificNode(client jagw.SubscriptionServiceClient) {
	log.Print("--------------------")
	log.Printf("Subscribing To Specific DataRates")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	makeCancellable(cancel)

	subscription := &jagw.TelemetrySubscription{
		SensorPath: proto.String("Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface"),
		Properties: []string{
			"data_rates/input_data_rate",
			"arp_information/arp_type_name",
			"arp_information/arp_timeout",
			"source",
		},
		Unflatten: proto.Bool(true),
		StringFilters: []*jagw.StringFilter{
			{
				Property: proto.String("source"),
				Value:    proto.String("XR-1"),
				Operator: jagw.StringOperator_EQUAL.Enum(),
			},
		},
	}

	// Subscribe via gRPC
	stream := api.SubscribeToTelemetryData(ctx, client, subscription)

	for {
		event, err := stream.Recv()
		if processEvent(ctx, event, err) {
			// As long as events can be processed successfully, print them
			log.Print("---------- EVENT ----------")
			prettyPrintTelemetryEvent(event)
		} else {
			// Some error occurred, abort subscription
			break

		}
	}
}
