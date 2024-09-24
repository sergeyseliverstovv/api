.PHONY: mycommand

generate-api:
	protoc -I proto proto/api.proto \
	--go_out=gen/go --go_opt=paths=source_relative \
	--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=gen/go --grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true --openapiv1_out=.api.proto
	

generate:
	protoc --proto_path proto --proto_path vendor.protogen \
	--go_out=.gen/go --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=.gen/go --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=.gen/go --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	proto/api.proto

gener:
	protoc -I proto proto/api.proto --go_out=gen/go --go_opt=paths=source_relative --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen/go --grpc-gateway_opt=paths=source_relative

generate-api-v1:
	protoc -I proto proto/api.proto \
	--go_out=gen/go/v1 --go_opt=paths=source_relative \
	--go-grpc_out=gen/go/v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=gen/go/v1 --grpc-gateway_opt=paths=source_relative \

generate-swagger:
	   protoc --proto_path=proto \
          --openapi_out=swagger \
          proto/api.proto
   
