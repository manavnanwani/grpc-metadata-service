## Starting the Central Client

```
cd client
go run *.go
```

## Starting Multiple Servers

```
cd server
go run *.go
```

Each instance will register with the client and respond to metadata requests.


## Adding New Services

1. When adding new service, create proto buffers and generate the corresponding Go files using the following command:
```
protoc --go_out=. --go-grpc_out=. proto/service/service.proto
```

2. Register the service on both client and servers