.PHONY: build
build:
	go build -o cmd/server/server cmd/server/*.go

.PHONY: test
test: build
	docker run -it -d -e POSTGRES_PASSWORD=pass -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 postgre
	go test ./...