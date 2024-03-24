package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

import proto "chat/pkg/chat_v1"

const grpcPort = 50052

type Server struct {
	proto.UnimplementedChatServer
}

func (s *Server) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Printf("Create chat: %v", req.GetUsernames())
	log.Printf("Usernames count: %v", len(req.GetUsernames()))
	return &proto.CreateResponse{Id: 1}, nil
}

func (s *Server) Delete(ctx context.Context, req *proto.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat: %v", req.GetId())
	return &emptypb.Empty{}, nil
}

func (s *Server) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message from: %v", req.GetFrom())
	log.Printf("Message: %v", req.GetText())
	log.Printf("at: %v", req.GetTimestamp())
	return &emptypb.Empty{}, nil
}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	proto.RegisterChatServer(s, &Server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
