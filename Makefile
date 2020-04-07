.PHONY: install build

clean:
	rm -rf src/backend/assets/* && \
	touch src/backend/assets/.gitkeep

front-install:
	cd src/front && \
	npm install

front-build:
	cd src/front && \
	npm run build && \
	rm -rf ../backend/assets/* && \
	touch ../backend/assets/.gitkeep && \
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

docker-build:
	make clean && \
	docker build -t log-uploader .
