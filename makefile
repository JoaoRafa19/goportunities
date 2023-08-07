.PHONY: default run build test docs clean

APP_NAME=gopornutities

default: run-with-docs


run-with-docs: 
	@swag init 
	@go run main.go
run: 
	@go run main.go 
build: 
	@go build -o $(APP_NAME) main.go 
build-win: 
	@go build -o $(APP_NAME).exe main.go
test:
	@go test ./ ... 
docs:
	@swag init 
clean:
	rm -f ./goportunities
	rm -f ./goportunities.exe
	rm -rf ./docs