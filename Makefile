compose:
	docker-compose up -d

test:
	go test -coverprofile=coverage.out -coverpkg=./... ./...

coverage:
	make test
	go tool cover -html=coverage.out -o coverage.html