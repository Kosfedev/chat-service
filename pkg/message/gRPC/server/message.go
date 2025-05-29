package server

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Kosfedev/chat-service/pkg/message/gRPC"
)

type server struct {
	desc.UnimplementedMessageV1Server
}

// NewServer is...
func NewServer() desc.MessageV1Server {
	return &server{}
}

// Create is...
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("chat id: %d\n", req.GetChatId())
	log.Printf("message text: %v\n", req.GetText())

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

// Get is...
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("message id: %d\n", req.GetId())

	return &desc.GetResponse{
		Id:        gofakeit.Int64(),
		ChatId:    gofakeit.Int64(),
		UserId:    gofakeit.Int64(),
		Text:      gofakeit.Quote(),
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: nil,
	}, nil
}

// List is...
func (s *server) List(_ context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	log.Printf("message ids: %d\n", req.GetIds())

	return &desc.ListResponse{
		Messages: []*desc.Message{{
			Id:        gofakeit.Int64(),
			ChatId:    gofakeit.Int64(),
			UserId:    gofakeit.Int64(),
			Text:      gofakeit.Quote(),
			CreatedAt: timestamppb.New(time.Now()),
			UpdatedAt: timestamppb.New(time.Now()),
		}, {
			Id:        gofakeit.Int64(),
			ChatId:    gofakeit.Int64(),
			UserId:    gofakeit.Int64(),
			Text:      gofakeit.Quote(),
			CreatedAt: timestamppb.New(time.Now()),
			UpdatedAt: timestamppb.New(time.Now()),
		}, {
			Id:        gofakeit.Int64(),
			ChatId:    gofakeit.Int64(),
			UserId:    gofakeit.Int64(),
			Text:      gofakeit.Quote(),
			CreatedAt: timestamppb.New(time.Now()),
			UpdatedAt: timestamppb.New(time.Now()),
		}},
		Limit:  0,
		Offset: 0,
	}, nil
}

// Update is...
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("message id: %d\n", req.GetId())
	log.Printf("message text: %v\n", req.GetText())

	return nil, nil
}

// Delete is...
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("message id: %d\n", req.GetId())

	return nil, nil
}
