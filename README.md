# sr-app
This application mocks a SR-App by making periodic requests to the Jalapeño API Gateway.

## gRPC
When the file `proto/request-service/requestservice.proto` is updated, this command needs to be run to recompile the code:
```bash
$ protoc --proto_path=./proto/requestservice --go_out=./proto/requestservice --go_opt=paths=source_relative --go-grpc_out=./proto/requestservice --go-grpc_opt=paths=source_relative ./proto/requestservice/requestservice.proto
```
When the file `proto/push-service/pushservice.proto` is updated, this command needs to be run to recompile the code:
```bash
$ protoc --proto_path=./proto/pushservice --go_out=./proto/pushservice --go_opt=paths=source_relative --go-grpc_out=./proto/pushservice --go-grpc_opt=paths=source_relative ./proto/pushservice/pushservice.proto
```

## Setting Up Development Environment
Make sure you have setup the [global development environment](https://gitlab.ost.ch/ins/jalapeno-api/request-service/-/wikis/Development-Environment) first.

### Step 1: Initialize Okteto
- Clone the repository:
```bash
$ git clone ssh://git@gitlab.ost.ch:45022/ins/jalapeno-api/sr-app.git
```
- Initialize okteto:
```bash
$ okteto init
```
- Replace content of okteto.yml with the following:
```yml
name: sr-app
autocreate: true
image: okteto/golang:1
command: bash
namespace: jagw-dev-<namespace-name>
securityContext:
  capabilities:
    add:
      - SYS_PTRACE
volumes:
  - /go/pkg/
  - /root/.cache/go-build/
  - /root/.vscode-server
  - /go/bin/
  - /bin/protoc/
sync:
  - .:/usr/src/app
forward:
  - 2346:2345
  - 8081:8080
environment:
  - REQUEST_SERVICE_ADDRESS=rs:9000
  - PUSH_SERVICE_ADDRESS=ps:9000
  - BROKER_ADDRESS=10.20.1.24:31133
```
