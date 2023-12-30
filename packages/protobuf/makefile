all: api.proto
	make server
	make client

server: api.proto
	protoc --go_out=../server/grpc/api --go_opt=paths=source_relative \
					--go-grpc_out=../server/grpc/api --go-grpc_opt=paths=source_relative \
			api.proto

client: api.proto
	protoc --go_out=../hardware/grpc/api --go_opt=paths=source_relative \
					--go-grpc_out=../hardware/grpc/api --go-grpc_opt=paths=source_relative \
			api.proto
