export INCLUDE_DIR=/tmp/protobuf/include:.


build-proto:
	@rm -rf ./proto/*.go
	@cd proto && protoc -I=${INCLUDE_DIR} --go_out=. ./*.proto
	@cd proto && mv grpc-starter/proto/*.go . && rm -rf grpc-starter