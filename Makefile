GO := GO111MODULE=on go
DOCKER := DOCKER_DEFAULT_PLATFORM=linux/amd64

.PHONY: ci
ci:
	$(GO) mod tidy && \
	$(GO) mod download && \
	$(GO) mod verify && \
	$(GO) mod vendor && \
	$(GO) fmt ./... \

.PHONY: build
build:
	$(GO) build -mod=vendor -a -installsuffix cgo -tags musl -o main ./cmd/main.go

start:
	go run --tags dynamic $(shell pwd)/cmd/main.go

dev: 
	nodemon --exec go run --tags dynamic $(shell pwd)/cmd/main.go --signal SIGTERM

.PHONY: clean
clean:
	@rm -rf main ./vendor

postgres:
	export POSTGRESQL_URL='postgres://postgres:1688@localhost:5432/product_api?sslmode=disable'

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

proto:
	protoc pkg/**/pb/*.proto --go-grpc_out=.

buf:
	buf generate && go run ./scripts/swagtxt.go

swag:
	swag init -g cmd/main.go