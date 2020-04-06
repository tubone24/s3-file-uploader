.PHONY: install build

front-install:
	cd src/front && \
	npm install

front-build:
	cd src/front && \
	npm run build && \
	rm -rf assets/* && \
	touch assets/.gitkeep && \
	cp -r dist/* ../backend/assets/

backend-install:
	go mod download

backend-build:
	cd src/backend && \
	go build -o ../../log-uploader

backend-run:
	cd src/backend && \
	go run main.go

backend-test:
	cd src/backend && \
	go test -v ./...
