ENV = $1

.PHONY: install build

clean:
	rm -rf src/backend/assets/* && \
	touch src/backend/assets/.gitkeep
	rm -rf output/* && \
	mkdir output/assets && \
	touch output/assets/.gitkeep

front-install:
	cd src/front && \
	npm install

front-build:
	cd src/front && \
	cp -f config/env/${ENV}.json config/config.json && \
	npm run build && \
	rm -rf ../backend/assets/* && \
	rm -rf ../../../output/assetes/* && \
	touch ../backend/assets/.gitkeep && \
	touch ../../output/assets/.gitkeep && \
	cp -r dist/* ../backend/assets/ && \
	cp -r dist/* ../../output/assets/

backend-install:
	go mod download

backend-build:
	cd src/backend && \
	cp -f config/env/${ENV}.toml config/config.toml && \
	go build -o ../../output/log-uploader && \
	mkdir ../../output/config && \
	cp config/config.toml ../../output/config/config.toml

backend-run:
	cd src/backend && \
	cp -f config/env/${ENV}.toml config/config.toml && \
	go run main.go

backend-test:
	cd src/backend && \
	go test -v ./...

docker-build:
	make clean && \
	docker build -t log-uploader --build-arg ENV="${ENV}" .
