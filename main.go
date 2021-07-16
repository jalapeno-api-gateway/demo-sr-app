package main

import (
	// "io"
	"log"
	"os"

	"golang.org/x/net/context"
	// "google.golang.org/grpc"

	// rsproto "gitlab.ost.ch/ins/jalapeno-api/sr-app/proto"
	"gitlab.ost.ch/ins/jalapeno-api/sr-app/kafka"
)

var lsNodeConsumer *kafka.Consumer

func main() {
	//Connect to Kafka
	ctx := context.Background()

	brokerAddress := os.Getenv("BROKER_ADDRESS")
	lsNodeTopic := "gobmp.parsed.ls_node"

	lsNodeConsumer = kafka.NewConsumer(ctx, brokerAddress, lsNodeTopic)

	consumeMessages(ctx, *lsNodeConsumer);

	// log.Print("Starting SR-App ...")
	// var conn *grpc.ClientConn
	// conn, err := grpc.Dial(os.Getenv("REQUEST_SERVICE_ADDRESS"), grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Could not connect: %s", err)
	// }
	// defer conn.Close()

	// client := rsproto.NewApiGatewayClient(conn)
	// ids := []string{"2_0_0_0000.0000.000a", "2_0_0_0000.0000.0001", "2_0_0_0000.0000.000c"}
	// message := &rsproto.NodeIds{Ids: ids}
	// stream, err := client.GetNodes(context.Background(), message)
	// if err != nil {
	// 	log.Fatalf("Error when calling GetNodes on RequestService: %s", err)
	// }

	// for {
	// 	node, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("%v.GetNodes(_) = _, %v", client, err)
	// 	}
	// 	log.Println(node)
	// }
	// log.Printf("---------------------All Nodes received---------------------")
}

func consumeMessages(ctx context.Context, consumer kafka.Consumer) {
	channel := make(chan []byte)
	go kafka.Consume(ctx, consumer, channel)

	for msg := range channel {
		log.Printf("%v\n", msg)
	}
}