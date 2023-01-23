# ==============================================================================
# Proto
# Generate protobuf in Go, gRPC for Go, gRPC Gateway for Go, and Swagger Documentation in JSON
.PHONY: proto
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
	--validate_out=pb --validate_opt=lang=go,paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-core_out=pb --grpc-core_opt=paths=source_relative,logtostderr=true \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=safety,logtostderr=true \
	proto/*.proto

# ==============================================================================
# Main
.PHONY: deps
deps:
	go mod tidy

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go build cmd/server/main.go

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: lint
lint:
	@echo  -e 'Starting golangci-lint'
	golangci-lint run ./...
	@echo  -e '\nStarting staticcheck'
	staticcheck -checks U1000 ./...

.PHONY: migrate
migrate:
	go run migrations/migration.go

# ==============================================================================
# Docker

.PHONY: docker-build
docker-build:
	docker build --tag szczynk/grpc-safety_core --progress=plain .

.PHONY: docker-run
docker-run:
	docker run --name grpc-safety_core -d -p 5000:5000 -p 7000:7000 szczynk/grpc-safety_core

# generate sql from GORM to golang-migrate
.PHONY: atlas
atlas:
	atlas migrate hash && \
	atlas migrate diff safety \
	--dir "file://migrations?format=golang-migrate" \
	--dev-url "postgres://postgres:postgres@:5432/safety?sslmode=disable" \
	--to "postgres://postgres:postgres@:5432/gorm?sslmode=disable"
