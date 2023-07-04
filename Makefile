APP_NAME=gai

default: clean build

build:
	@echo "Building... in to $(HOME)/bin/$(APP_NAME)"
	@go build -o $(HOME)/bin/$(APP_NAME) 

clean:
	@echo "Cleaning..."
	@rm -rf $(HOME)/bin/$(APP_NAME)

.PHONY: build clean