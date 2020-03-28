.PHONY: install build

install:
	go mod download

build:
	cd src && \
	go build -o ../log-uploader

run:
	cd src && \
	go run main.go
