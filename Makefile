run:
	ENV=dev go run ./cmd/sports-line-processor/main.go

genproto:
	protoc --go_out=. --go-grpc_out=. ./api/proto/*.proto