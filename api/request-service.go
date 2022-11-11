package api

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
)

func RequestCoordinates(client jagw.RequestServiceClient, request *jagw.LsNodeCoordinatesRequest) *jagw.LsNodeCoordinatesResponse {
	response, err := client.GetLsNodeCoordinates(context.TODO(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsNodesCoordinates on request service: %s", err)
	}
	return response
}

func RequestNodes(client jagw.RequestServiceClient, request *jagw.TopologyRequest) *jagw.LsNodeResponse {
	response, err := client.GetLsNodes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsNodes on request service: %s", err)
	}
	return response
}

func RequestLinks(client jagw.RequestServiceClient, request *jagw.TopologyRequest) *jagw.LsLinkResponse {
	response, err := client.GetLsLinks(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsLinks on request service: %s", err)
	}
	return response
}

func RequestTelemetryData(client jagw.RequestServiceClient, request *jagw.TelemetryRequest) *jagw.TelemetryResponse {
	response, err := client.GetTelemetryData(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetTelemetryData on request service: %s", err)
	}
	return response
}
