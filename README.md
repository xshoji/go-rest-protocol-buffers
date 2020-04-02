## About

Implementation of rest api by go based on [Protocol Buffers](https://developers.google.com/protocol-buffers) (proto3).

This api is implemented using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

## Requirements

```
brew install protobuf
go get -u google.golang.org/grpc
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Build and run

```
// generage code
./generate_model.sh -m usersapi

// run grpc and gateway server
go run cmd/main/main.go
```

## How to implement

### Define proto file

Create and Define user.proto in `./proto/user.proto`.

> Language Guide (proto3)  ｜  Protocol Buffers  ｜  Google Developers  
> https://developers.google.com/protocol-buffers/docs/proto3

### Generate golang codes

```
./generate_model.sh -m usersapi
```

### Implement service

Implement user.pb.go interfaces in `./usersapi/impl/user.go`.

```
package proto_impl
...
type SampleService struct{}

func (*SampleService) Create(context context.Context, user *proto.User) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) Read(ctx context.Context, in *proto.SampleServiceSelector) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) ReadAll(ctx context.Context, in *empty.Empty) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) Update(context context.Context, user *proto.User) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) Delete(context context.Context, in *proto.SampleServiceSelector) (*proto.SampleServiceResponse, error) {...}

```

### Register handler to gateway.go

```
// newGateway()
...
	err = proto.RegisterSampleServiceHandler(ctx, mux, conn)
...
```

### Register service to main.go

```
// main()
...
		proto.RegisterSampleServiceServer(s, new(proto_impl.UserService))
...
```

## Call api (grpc-gateway)

#### Create

```
curl \
  -vvv \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '
{
  "name": "john-'$(openssl rand -hex 15 | fold -w 15 |head -n 1)'",
  "nickname": "nick-'$(openssl rand -hex 15 | fold -w 15 |head -n 1)'",
  "age": '$((RANDOM%100+1))',
  "birth": {
    "datetime": "2020-03-31T17:09:09.916Z",
    "weight": '$((RANDOM%100+1))',
    "hospital": "string"
  },
  "addresss": {
    "country": "US",
    "zipCode": '$((RANDOM%100+1))',
    "state": "string",
    "city": "string"
  }
}
' \
'http://localhost:8080/v1/users'
```

#### Get

```
// Get by id
curl -X GET http://localhost:8080/v1/users/1

// Get All
curl -X GET http://localhost:8080/v1/users
```

####  Update

```
curl \
  -vvv \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '
{
  "name": "john-updated"
}
' \
'http://localhost:8080/v1/users/1'
```

#### Delete

```
curl -X DELETE http://localhost:8080/v1/users/1
```

## References

> grpc-gatewayでgRPCサーバをRESTで叩けるようにする - Carpe Diem  
> https://christina04.hatenablog.com/entry/2017/11/15/034455

> gRPCを使ってみた - チャーリー！のテクメモ  
> http://charleysdiary.hatenablog.com/entry/2016/09/08/163909

> Protocol Buffers(proto3)でoptionalをどう扱うか - Qiita  
> https://qiita.com/disc99/items/a8ac2a264f322bc6d6e5

> Usage ｜ grpc-gateway  
> https://grpc-ecosystem.github.io/grpc-gateway/docs/usage.html

**Validation**

> envoyproxy/protoc-gen-validate： protoc plugin to generate polyglot message validators  
> https://github.com/envoyproxy/protoc-gen-validate
