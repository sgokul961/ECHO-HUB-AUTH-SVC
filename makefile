.PHONY: run

run:
	go run ./cmd


proto:


protoc --go_out=paths=source_relative:.        --go-grpc_out=paths=source_relative:.        pkg/pb/auth.proto
