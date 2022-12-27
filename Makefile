PACKAGE=github.com/SamMHD/simple-broker

LDFLAG_SENDER = "-X 'main.TargetService=sender'"
LDFLAG_RECEIVER = "-X 'main.TargetService=receiver'"
LDFLAG_BROKER = "-X 'main.TargetService=broker'"
LDFLAG_DESTINATION = "-X 'main.TargetService=destination'"
LDFLAG_ALL = "-X 'main.TargetService=all'"

# clear-log is used to clear previous logs
clear-log:
	@echo "Clearing Previous Logs..."
	@rm *.log

# create-bin-folder is used to create bin folder (if not exists)
create-bin-folder:
	@mkdir -p bin

# proto is used to generate protobuf files and packages
proto:
	rm -f pb/*.go
	protoc --proto_path=./proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

#################### SENDER PACKAGE ####################
build-sender: create-bin-folder
	@echo "Building Sender Service..."
	go build -o bin/sender -ldflags=${LDFLAG_SENDER} main.go
	@echo "Build Done!"
	@echo

run-sender: build-sender 
	./bin/sender $(ARGS)
#################### SENDER PACKAGE ####################


#################### Receiver PACKAGE ####################
build-receiver: create-bin-folder
	@echo "Building Receiver Service..."
	go build -o bin/receiver -ldflags=${LDFLAG_RECEIVER} main.go
	@echo "Build Done!"
	@echo

run-receiver: build-Receiver
	./bin/receiver $(ARGS)
#################### Receiver PACKAGE ####################


#################### BROKER PACKAGE ####################
build-broker: create-bin-folder
	@echo "Building Broker Service..."
	go build -o bin/broker -ldflags=${LDFLAG_BROKER} main.go
	@echo "Build Done!"
	@echo

run-broker: build-broker
	./bin/broker $(ARGS)
#################### BROKER PACKAGE ####################


#################### DESTINATION PACKAGE ####################
build-destination: create-bin-folder
	@echo "Building Destination Service..."
	go build -o bin/destination -ldflags=${LDFLAG_DESTINATION} main.go
	@echo "Build Done!"
	@echo

run-destination: build-destination
	./bin/destination $(ARGS)
#################### DESTINATION PACKAGE ####################


#################### ALL-IN-ONE PACKAGE ####################
build: create-bin-folder
	@echo "Building All-in-One Service..."
	go build -o bin/all -ldflags=${LDFLAG_ALL} main.go
	@echo "Build Done!"
	@echo

run: build
	@echo "Running All-in-One Service..."
	./bin/all $(ARGS)
#################### ALL-IN-ONE PACKAGE ####################

# Build-All is used to build all services separately
build-all:
	@echo "Building All Services..."
	@make build-sender
	@make build-receiver
	@make build-broker
	@make build-destination

.PHONY: clear-log create-bin-folder proto build-sender run-sender build-Receiver run-Receiver build-broker run-broker build-destination run-destination build-all build run