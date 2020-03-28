.PHONY: install build

backend-install:
	go mod download

backend-build:
	cd src/backend && \
	go build -o ../log-uploader

backend-run:
	cd src/backend && \
	go run main.go
