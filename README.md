# gRPC Demo in Golang

This is a simple demonstration of gRPC in Golang. The project includes a server and a client, with communication defined by a `.proto` file. The project uses `protoc` for protocol buffer compilation.


## Requirements

- [Golang](https://golang.org/doc/install) >= 1.16
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- gRPC and Protocol Buffers Go modules:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

## How To Run 
- change to server directory
 ```bash
  cd server
```
- run the server
  ```bash
  go run *.go
  ```

- change to client directory
 ```bash
  cd client
  ```
- run the server
  ```bash
  go run *.go
  ```


