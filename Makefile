.PHONY: build
build:
	go run cmd/main.go

.DEFAULT_GOAL := build