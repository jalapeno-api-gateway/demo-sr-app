package fetch

import (
	"time"

	"github.com/jalapeno-api-gateway/demo-sr-app/api"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/protobuf/proto"
)

func GetDataRatesOfSpecificNode(rsClient jagw.RequestServiceClient) {
	request := &jagw.TelemetryRequest{
		SensorPath: proto.String("Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface"),
		Properties: []string{
			"source",
			"if_index",
			"data_rates/input_data_rate",
		},
		// StringFilters: []*jagw.StringFilter{
		// 	// Filter for measurements where source == "XR-3"
		// 	{
		// 		Property: proto.String("source"),
		// 		Value: proto.String("XR-3"),
		// 		Operator: jagw.StringOperator_EQUAL.Enum(),
		// 	},
		// },
		Unflatten: proto.Bool(true),
		RangeFilter: &jagw.RangeFilter{
			// Get all measurements from the last 60 seconds
			EarliestTimestamp: proto.Int64(time.Now().Add(-60 * time.Second).UnixNano()),
		},
	}
	response := api.RequestTelemetryData(rsClient, request)

	prettyPrintTelemetryData(response)
}

func GetLatestMeasurement(rsClient jagw.RequestServiceClient, unflatten bool) {
	request := &jagw.TelemetryRequest{
		SensorPath: proto.String("Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface"),
		Properties: []string{
			"source",
			"if_index",
			"data_rates/input_data_rate",
			"data_rates/output_data_rate",
			"arp_information/arp_timeout",
			"arp_information/arp_type_name",
		},
		StringFilters: []*jagw.StringFilter{
			// Filter for measurements where source == "XR-3"
			{
				Property: proto.String("source"),
				Value: proto.String("XR-3"),
				Operator: jagw.StringOperator_EQUAL.Enum(),
			},
		},
		Unflatten: proto.Bool(unflatten),
	}
	response := api.RequestTelemetryData(rsClient, request)

	prettyPrintTelemetryData(response)
}
