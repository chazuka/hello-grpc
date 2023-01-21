BIN_DIR = bin
SERVER_DIR = server
CLIENT_DIR = client
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

.PHONY: greet calculator
project := greet calculator

all: $(project) ## Generate Pbs and build

greet: $@ ## Generate Pbs and build for greet
calculator: $@ ## Generate Pbs and build for calculator

$(project):
	protoc -I$@ --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}