.PHONY: run build test clean

BINARY_NAME=inforce-task.exe
APP_PATH=./cmd/app/main.go

## run: run the application
run:
	@$(BINARY_NAME)

## build: build the application
build:
	@go build -o $(BINARY_NAME) $(APP_PATH)

## test: run tests
test:
	@go test ./internal/client/... ./internal/service/...

## clean: clean the project
clean:
	@go clean
	@rm -f $(BINARY_NAME)

swagger:
	@swag init -g $(APP_PATH)
