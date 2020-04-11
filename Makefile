
PROJECT_PATH=$(shell pwd)
PROTO_FILES=greet.proto
PROTO_PATH=greet/greetpb
PROTO_DEST=greet/greetpb
#https://hub.docker.com/r/znly/protoc/
PROTOC_DOCKER=znly/protoc



# Generating stubs without docker:  protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
proto:
	@ echo " ---   GENERATING PROTOBUF   --- "
	@ docker run --rm -v $(PROJECT_PATH):$(PROJECT_PATH) -w $(PROJECT_PATH) $(PROTOC_DOCKER)  --go_out=plugins=grpc:./$(PROTO_DEST) -I $(PROTO_PATH) $(PROTO_FILES)
	@ echo " ---     FINISH PROTOBUF     --- "

run-server:
	@ go run greet/greet_server/server.go

run-client:
	@ go run greet/greet_client/client.go