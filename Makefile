
PROJECT=sql2csv
BIN=$(CURDIR)/bin
EXEC=$(PROJECT)


all: build 

build:
	go build -o $(BIN)/

test:
	go test -v 

dep:
	go mod tidy
	
cc:
	GOOS=linux GOARCH=amd64 go build -o $(BIN)/$(EXEC) 
	GOOS=windows GOARCH=amd64 go build -o $(BIN)/$(EXEC).exe