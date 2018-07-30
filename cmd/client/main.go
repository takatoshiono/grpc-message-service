package main

import (
	"context"
	"log"

	pb "github.com/takatoshiono/grpc-message-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50101", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)
	convReq := &pb.CreateConversationRequest{}
	conversation, err := client.CreateConversation(context.Background(), convReq)
	if err != nil {
		log.Fatalf("failed to CraeteConversation(): %v", err)
	}
	log.Printf("created %v", conversation)

	messReq := &pb.CreateMessageRequest{
		ConversationId: conversation.Id,
		Sender:         "bob",
		Body:           "This is bob. Hello alice.",
	}
	message, err := client.CreateMessage(context.Background(), messReq)
	if err != nil {
		log.Fatalf("failed to CreateMessage(): %v", err)
	}
	log.Printf("created %v", message)
}
