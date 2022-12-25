broker:
	@echo "Starting broker..."
	go run broker/main.go

destination:
	@echo "Starting destination..."
	go run destination/main.go

receiver:
	@echo "Starting receiver..."
	go run receiver/main.go

sender:
	@echo "Starting sender..."
	go run sender/main.go

all:
	@echo "Starting all..."
	@make broker
	@make destination
	@make receiver
	@make sender

proto:
	rm -f pb/*.go
	protoc --proto_path=./proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: broker destination receiver sender all proto