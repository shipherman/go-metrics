d := $(shell date +'%Y/%m/%d %H:%M:%S')
.PHONY: build
build:
	go build -ldflags "-X main.buildVersion=v1.0.1 -X 'main.buildDate=$d' -X main.buildCommit=`git rev-parse HEAD`" -o bin/server cmd/server/*.go
	go build -o bin/agent cmd/agent/*.go
	go build -o bin/checker checker/checker.go
	go build -o bin/certgen cmd/certgen/certgen.go

.PHONY: test
test:
	docker run -it -d -e POSTGRES_PASSWORD=pass -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 postgre
	go test ./...

.PHONY: checkserver
checkserver:
	bin/checker cmd/server/*.go

.PHONY: checkagent
checkagent:
	bin/checker cmd/agent/*.go
