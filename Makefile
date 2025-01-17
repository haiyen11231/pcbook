gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...