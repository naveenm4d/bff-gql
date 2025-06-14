BIN?=bff-gql
REGISTRY?=localhost

default: run
.PHONY : build run

build:
	GOOS=linux GOARCH=arm64 go build -o build/${BIN}

build-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional proto/blogs_svc.proto

clean:
	go clean
	rm -rf build
	
lint:
	golangci-lint run -c .golangci.yml

run: build
	./build/${BIN}

test:
	go test ./... -tags musl -coverprofile=coverage.txt -covermode count