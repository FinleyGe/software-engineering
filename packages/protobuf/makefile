all: api.proto
	protoc --go_out=../server/grpc/api --go_opt=paths=source_relative \
          --go-grpc_out=../server/grpc/api --go-grpc_opt=paths=source_relative \
      api.proto
