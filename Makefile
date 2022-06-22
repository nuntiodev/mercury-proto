.PHONY: go-proto
go-proto:
	protoc -I. --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./mercury.proto
	
.PHONY: build-proto
build-proto:
	make go-proto