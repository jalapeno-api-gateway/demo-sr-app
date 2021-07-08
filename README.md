# sr-app
This application mocks a SR-App by making periodic requests to the Jalapeño API Gateway.

## Proto Buffers
When the file `request-service/request-service.proto` is updated, this command needs to be run to recompile the code:
```bash
$ /bin/protoc/bin/protoc --proto_path=./request-service --go_out=./request-service --go_opt=paths=source_relative --go-grpc_out=./request-service --go-grpc_opt=paths=source_relative ./request-service/request-service.proto
```