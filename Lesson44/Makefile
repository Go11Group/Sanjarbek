protoc:
	protoc --go_out=server --go_opt=paths=source_relative translate.proto
	protoc --go-grpc_out=server --go-grpc_opt=paths=source_relative translate.proto

run_server:
	go run server.go

run_client:
	go run client/client.go