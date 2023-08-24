.PHONY: build
build:
	go build -o bin/server cmd/server/*.go
	go build -o bin/agent cmd/agent/*.go

.PHONY: test
test:
	docker run -it -d -e POSTGRES_PASSWORD=pass -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 postgre
	go test ./...