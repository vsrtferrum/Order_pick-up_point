ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=vsrtf dbname=postgres  sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MOCKGEN_TAG=1.2.0
BINARY_NAME = gohw_4
MIGRATION_FOLDER = ./migrations
LOCAL_BIN:=$(CURDIR)/bin
PVZ_PROTO_PATH:="api/proto/pvz/v1"

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest


build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-limux ./cmd/main.go

run: build
	./bin/${BINARY_NAME}-linux

clean:
	go clean
	rm ./

build-compose:
	docker-compose build

up-all:
	docker-compose up -d zookeeper kafka1

down:
	docker-compose down

run-compose:
	go run ./cmd/route-kafka

.PHONY: migration-create
migration-create:
	goose -dir "$(`)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

# docker-compose aliases
.PHONY: compose-rs
compose-rs:
	make compose-rm
	make compose-up

.PHONY: compose-up
compose-up:
	docker-compose -p ws_5 up -d

.PHONY: compose-rm
compose-rm:
	docker-compose -p ws_5 rm -fvs

.PHONY: .generate-mockgen-deps
.generate-mockgen-deps:
ifeq ($(wildcard $(MOCKGEN_BIN)),)
	@GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@$(MOCKGEN_TAG)
endif

.PHONY: test
test:
	$(info running tests...)
	go test ./internal/tests/

.PHONY: .generate-mockgen
generate-mockgen:
	PATH="$(LOCAL_BIN):$$PATH" go generate -x -run=mockgen ./...

.PHONY: .generate-mock
generate-mockgen:
	find . -name '*_mock.go' -delete
	mockgen -source ./internal/handlers/input/inputstruct.go -destination=./internal/handlers/input/mocks/inputstruct_mock.go -package=mock_inputstruct
	mockgen -source ./internal/module/filters.go -destination=./internal/module/mocks/filters_mock.go -package=mock_filters
	mockgen -source ./internal/storage/storage.go -destination=./internal/storage/mocks/storage_mock.go -package=mock_storage
	mockgen -source ./internal/handlers/cli/cli.go -destination=./internal/handlers/cli/mocks/cli_mock.go -package=mock_cli
	mockgen -source ./internal/handlers/log/log.go -destination=./internal/handlers/log/mock/log_mock.go -package=mock_log
	mockgen -source ./internal/api/service.go -destination=./internal/api/mock/service_mock.go -package=mock_service



.vendor-proto: vendor-proto/google/protobuf vendor-proto/google/api vendor-proto/protoc-gen-openapiv2/options vendor-proto/validate

vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor.proto/grpc-ecosystem && \
 	cd vendor.proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor.proto/protoc-gen-openapiv2
	mv vendor.proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor.proto/protoc-gen-openapiv2
	rm -rf vendor.proto/grpc-ecosystem

vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor.proto/protobuf &&\
	cd vendor.proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor.proto/google
	mv vendor.proto/protobuf/src/google/protobuf vendor.proto/google
	rm -rf vendor.proto/protobuf

vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor.proto/googleapis && \
 	cd vendor.proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor.proto/google
	mv vendor.proto/googleapis/google/api vendor.proto/google
	rm -rf vendor.proto/googleapis

vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor.proto/tmp && \
		cd vendor.proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor.proto/validate
		mv vendor.proto/tmp/validate vendor.proto/
		rm -rf vendor.proto/tmp

.PHONY: generate
generate: .bin-deps .vendor-proto
	mkdir -p pkg/${PVZ_PROTO_PATH}
	protoc -I api/proto \
		-I vendor.proto \
		${PVZ_PROTO_PATH}/pvz.proto \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=./pkg/${PVZ_PROTO_PATH} --go_opt=paths=source_relative\
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=./pkg/${PVZ_PROTO_PATH} --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway --grpc-gateway_out ./pkg/${PVZ_PROTO_PATH} --grpc-gateway_opt  paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
		--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 --openapiv2_out=./pkg/${PVZ_PROTO_PATH} \
		--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate --validate_out="lang=go,paths=source_relative:pkg/api/proto/pvz/v1"

.PHONY: up-jaeger
up-jaeger:
	docker compose up -d jaeger
.PHONY: up-redis
up-redis:
	docker compose up -d redis
.PHONY: run-prometheus
run-prometheus:
	prometheus --config.file config/prometheus.yaml
