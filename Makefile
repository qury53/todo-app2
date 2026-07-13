include .env
export
export PROJECT_ROOT=$(shell pwd)

run:
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/main.go
