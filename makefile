#*============================================================================*#
#*=====*                          Generate GRPC                         *=====*#
#*============================================================================*#
.PHONY:protobuf-gen

protobuf-gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative internal/adapters/grpc/proto/v1/*.proto

protobuf-clean:
	@rm -rf internal/adapters/grpc/proto/v1/*.pb.go

#*============================================================================*#
#*=====*                          Run Server                            *=====*#
#*============================================================================*#
.PHONY:local-serve

local-serve:
	@go run . serve

#*============================================================================*#
#*=====*                           Dev Docker                           *=====*#
#*============================================================================*#
.PHONY: docker-side

docker-side:
	@docker compose --profile side up
