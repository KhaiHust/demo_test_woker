compose:
	docker-compose up -d

test:
	go test -coverprofile=coverage.out -coverpkg=./... ./...

coverage: test
	go tool cover -html=coverage.out -o coverage.html

.PHONY: compose test coverage