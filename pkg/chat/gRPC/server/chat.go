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

func NewServer() desc.ChatV1Server {
	return &server{}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("chat name: %v\n", req.GetName())

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

func (s *server) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
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

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetId())
	log.Printf("user name: %v\n", req.GetName())

	return nil, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetId())

	return nil, nil
}

func (s *server) AddUser(ctx context.Context, req *desc.AddUserRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetUserId())

	return nil, nil
}

func (s *server) RemoveUser(ctx context.Context, req *desc.RemoveUserRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetUserId())

	return nil, nil
}
