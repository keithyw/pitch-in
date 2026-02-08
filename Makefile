test:
	go test -v ./backend/...

build:
	docker-compose build

docker-up:
	docker-compose up -d