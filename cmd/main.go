package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	descChat "github.com/Kosfedev/chat-service/pkg/chat/gRPC"
	grpcChatServer "github.com/Kosfedev/chat-service/pkg/chat/gRPC/server"
	"github.com/Kosfedev/chat-service/pkg/message/http/handlers"
)

const (
	pathname     = "localhost"
	httpPort     = 8091
	grpcPort     = 8092
	chatsPostfix = "/chats"
	chatPostfix  = chatsPostfix + "/{id}"
)

func main() {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err = runHTTPServer()
	}()

	go func() {
		defer wg.Done()
		err = runGRPCServer()
	}()

	wg.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func runHTTPServer() error {
	r := chi.NewRouter()
	r.Post(chatsPostfix, handlers.CreateChatHandler)
	r.Delete(chatPostfix, handlers.DeleteChatHandler)
	r.Post(chatPostfix, handlers.CreateMessageHandler)

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", pathname, httpPort),
		Handler:      r,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	log.Printf("http server listening on: %d\n", httpPort)
	return server.ListenAndServe()
}

func runGRPCServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", pathname, grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	descChat.RegisterChatV1Server(s, &grpcChatServer.Server{})

	log.Printf("gRPC server listening on: %d\n", grpcPort)
	return s.Serve(lis)
}
