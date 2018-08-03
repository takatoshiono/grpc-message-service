package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/takatoshiono/grpc-message-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	key := flag.String("api-key", "", "API key")
	flag.Parse()

	conn, err := grpc.Dial(":50101", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	if *key != "" {
		log.Printf("Using API key: %s", *key)
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-api-key", *key))
	}

	client := pb.NewMessageServiceClient(conn)
	convReq := &pb.CreateConversationRequest{}
	conversation, err := client.CreateConversation(ctx, convReq)
	if err != nil {
		log.Fatalf("failed to CraeteConversation(): %v", err)
	}
	log.Printf("created %v", conversation)

	messReq := &pb.CreateMessageRequest{
		ConversationId: conversation.Id,
		Sender:         "bob",
		Body:           "This is bob. Hello alice.",
	}
	message, err := client.CreateMessage(ctx, messReq)
	if err != nil {
		log.Fatalf("failed to CreateMessage(): %v", err)
	}
	log.Printf("created %v", message)
}
