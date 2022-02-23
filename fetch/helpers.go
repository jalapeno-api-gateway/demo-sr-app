package fetch

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
)

// Making a subscription cancellable for demo purposes
func makeCancellable(cancel context.CancelFunc) {
	userInput := make(chan bool, 1)
	// Wait for user input in a separate go routine.
	go waitForUserInput(userInput)
	go func() {
		// Block here (in this separate go routine) until user has pressed a key, then cancel the subscription (through context)
		<-userInput
		cancel()
	}()
}

func waitForUserInput(userInput chan bool) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Press 'Enter' to cancel subscription\n")
	input.Scan()
	userInput <- true
}

func prettyPrint(any interface{}) {
	s, _ := json.MarshalIndent(any, "", "  ")
	fmt.Printf("%s\n\n", string(s))
}

func prettyPrintTelemetryData(response *jagw.TelemetryResponse) {
	telemetryDataSlice := make([]map[string]interface{}, len(response.TelemetryData))
	
	for i := 0; i < len(response.TelemetryData); i++ {
		telemetryDataSlice[i] = makeMapFromJsonString(response.TelemetryData[i])
	}
	
	responseAsMap := make(map[string]interface{})
	responseAsMap["TelemetryData"] = telemetryDataSlice

	prettyPrint(responseAsMap)
}

func prettyPrintTelemetryEvent(event *jagw.TelemetryEvent) {
	telemetryData := makeMapFromJsonString(*event.TelemetryData)
	
	responseAsMap := make(map[string]interface{})
	responseAsMap["TelemetryData"] = telemetryData

	prettyPrint(responseAsMap)
}

func makeMapFromJsonString(jsonString string) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		log.Panicf("%v\n", err)
	}
	return m
}

func processEvent(ctx context.Context, event interface{}, err error) bool {
	if err == io.EOF {
		// Server has reached EOF, subscription may end silently (you may handle it appropriately)
		return false
	}
	if err != nil {
		if ctx.Err() != nil {
			// User has cancelled the subscription manually by pressing a key on
			// the keyboard, subscription may end silently
			return false
		}
		// Some other error occured
		log.Fatalf("%v\n", err)
	}
	
	return true
}