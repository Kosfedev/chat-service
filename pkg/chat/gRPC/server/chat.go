package server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Kosfedev/chat-service/pkg/chat/gRPC"
)

type server struct {
	desc.UnimplementedChatV1Server
}

// NewServer is...
func NewServer() desc.ChatV1Server {
	return &server{}
}

// Create is...
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("chat name: %v\n", req.GetName())

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

// Get is...
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("user id: %d\n", req.GetId())

	return &desc.GetResponse{
		Id:         req.GetId(),
		Name:       gofakeit.Name(),
		UserIds:    []int64{1, 2, 3},
		MessageIds: []int64{3, 2, 1},
		CreatedAt:  timestamppb.New(gofakeit.Date()),
		UpdatedAt:  timestamppb.New(gofakeit.Date()),
	}, nil
}

// List is...
func (s *server) List(_ context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	log.Printf("user id: %v\n", req.GetUserId())
	log.Printf("user name: %v\n", req.GetName())

	return &desc.ListResponse{
		Chats: []*desc.Chat{{
			Id:         1,
			Name:       gofakeit.Name(),
			UserIds:    []int64{1, 2, 3},
			MessageIds: []int64{3, 2, 1},
			CreatedAt:  timestamppb.New(gofakeit.Date()),
			UpdatedAt:  timestamppb.New(gofakeit.Date()),
		}},
		Limit:  gofakeit.Int64(),
		Offset: gofakeit.Int64(),
	}, nil
}

// Update is...
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetId())
	log.Printf("user name: %v\n", req.GetName())

	return nil, nil
}

// Delete is...
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetId())

	return nil, nil
}

// AddUser is...
func (s *server) AddUser(_ context.Context, req *desc.AddUserRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetUserId())

	return nil, nil
}

// RemoveUser is...
func (s *server) RemoveUser(_ context.Context, req *desc.RemoveUserRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetUserId())

	return nil, nil
}
