LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

generate:
	make generate-chat-api generate-message-api

generate-chat-api:
	mkdir -p pkg/chat/gRPC
	protoc 	--proto_path api/chat_v1 \
            --go_out=pkg/chat/gRPC --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/chat/gRPC --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            api/chat_v1/chat.proto

generate-message-api:
	mkdir -p pkg/message/gRPC
	protoc 	--proto_path api/message_v1 \
            --go_out=pkg/message/gRPC --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/message/gRPC --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            api/message_v1/message.proto