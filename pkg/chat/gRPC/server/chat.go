package server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Kosfedev/chat-service/pkg/chat/gRPC"
)

type Server struct {
	desc.UnimplementedChatV1Server
}

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

func (s *Server) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
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

func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetId())
	log.Printf("user name: %v\n", req.GetName())

	return nil, nil
}

func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetId())

	return nil, nil
}

func (s *Server) AddUser(ctx context.Context, req *desc.AddUserRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetUserId())

	return nil, nil
}

func (s *Server) RemoveUser(ctx context.Context, req *desc.RemoveUserRequest) (*emptypb.Empty, error) {
	log.Printf("user id: %d\n", req.GetUserId())

	return nil, nil
}
