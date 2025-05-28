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

func NewServer() desc.MessageV1Server {
	return &server{}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("chat id: %d\n", req.GetChatId())
	log.Printf("message text: %v\n", req.GetText())

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

func (s *server) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
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

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("message id: %d\n", req.GetId())
	log.Printf("message text: %v\n", req.GetText())

	return nil, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("message id: %d\n", req.GetId())

	return nil, nil
}
