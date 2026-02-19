test:
	cd backend && go test -v ./...

test-coverage:
	cd backend && go test -cover ./...

test-coverage-html:
	cd backend && go test -coverprofile=coverage.out ./...
	cd backend && go tool cover -html=coverage.out

test-clean:
	cd backend && go clean -testcache && go test -v ./...

build:
	docker-compose build

docker-up:
	docker-compose up -d