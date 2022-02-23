package fetch

import (
	"github.com/jalapeno-api-gateway/demo-sr-app/api"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
)

func GetAllNodes(rsClient jagw.RequestServiceClient) {
	request := &jagw.TopologyRequest{}
	response := api.RequestNodes(rsClient, request)

	prettyPrint(response)
}

func GetSpecificNodes(rsClient jagw.RequestServiceClient) {
	request := &jagw.TopologyRequest{
		Keys: []string{
			"2_0_0_0000.0000.0005",
			"2_0_0_0000.0000.000d",
		},
		PropertyNames: []string{
			"Key",
			"Name",
			"Asn",
			"RouterIp",
		},
	}
	response := api.RequestNodes(rsClient, request)

	prettyPrint(response)
}