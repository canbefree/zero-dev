protoc -I ./proto -I /usr/local/include proto/demo/*.proto --go_out=./proto/demo
protoc -I ./proto -I /usr/local/include proto/demo/*.proto  --go-grpc_out=./proto/demo
protoc -I ./proto -I /usr/local/include proto/demo/*.proto --grpc-gateway_out=./proto/demo