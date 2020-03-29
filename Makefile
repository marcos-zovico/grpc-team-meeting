.PHONY: build


PROJECT_PATH=$(shell pwd)
PROTO_FILES=greet.proto
PROTO_PATH=greet/greetpb
PROTO_DEST=greet/greetpb
#https://hub.docker.com/r/znly/protoc/
PROTOC_DOCKER=znly/protoc


proto:
	@echo " ---   GENERATING PROTOBUF   --- "
	# @ mkdir -p $(PROTO_DEST)
	@ docker run --rm -v $(PROJECT_PATH):$(PROJECT_PATH) -w $(PROJECT_PATH) $(PROTOC_DOCKER)  --go_out=plugins=grpc:./$(PROTO_DEST) -I $(PROTO_PATH) $(PROTO_FILES)
	@echo " ---     FINISH PROTOBUF     --- "
