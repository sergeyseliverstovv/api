generate-api:
	protoc -I proto proto/api.proto\
	--go_out=.gen/go/api --go_opt=paths=source_relative\
	--go-grpc_out=.gen/go/api --go-grpc_opt=paths=source_relative\
	--grpc-gateway_out=.gen/go/api--grpc-gateway_opt=paths=source_relative 