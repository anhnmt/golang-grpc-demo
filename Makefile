

go.protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    hello/hello.proto

go.client:
	go run ./client/client.go

go.server:
	go run ./server/server.go