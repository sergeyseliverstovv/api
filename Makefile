.PHONY: mycommand

generate-api:
	protoc -I proto proto/api.proto \
	--go_out=gen/go --go_opt=paths=source_relative \
	--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=gen/go --grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true --openapiv1_out=.api.proto
	

generate-api-v1:
	protoc -I proto proto/api.proto \
	--go_out=gen/go/v1 --go_opt=paths=source_relative \
	--go-grpc_out=gen/go/v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=gen/go/v1 --grpc-gateway_opt=paths=source_relative \


   
