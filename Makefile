
# ==============================================================================
# Docker
# TODO(Bagus): from env var to makefile var
# ADMIN_USER='admin' ADMIN_PASSWORD='admin' ADMIN_PASSWORD_HASH='$2a$14$1l.IozJx7xQRVmlkEQ32OeEEfP5mRxTpbDTCTcXRqn19gXD8YK1pO' docker-compose -f docker-compose.monitor.local.yml up -d
ADMIN_USER=admin
ADMIN_PASSWORD=admin
ADMIN_PASSWORD_HASH=$$2a$$14$$1l.IozJx7xQRVmlkEQ32OeEEfP5mRxTpbDTCTcXRqn19gXD8YK1pO

.PHONY: local-up
local-up:
	set -e
	@echo 'fetching 3893.4 MB of required docker images to run'
	docker-compose -f docker-compose.local.yml -f docker-compose.monitor.local.yml up -d
	@echo 'nevermind, already fetched'

.PHONY: local-down
local-down:
	docker-compose -f docker-compose.local.yml -f docker-compose.monitor.local.yml down

.PHONY: local-down-v
local-down-v:
	docker-compose -f docker-compose.local.yml -f docker-compose.monitor.local.yml down -v

.PHONY: local2-up
local2-up:
	@echo 'fetching 3893.4 MB of required docker images to run'
	docker-compose -f docker-compose.local.yml up -d
	@echo 'nevermind, already fetched'

.PHONY: local2-down
local2-down:
	docker-compose -f docker-compose.local.yml down

.PHONY: local2-down-v
local2-down-v:
	docker-compose -f docker-compose.local.yml down -v

.PHONY: docker-up
docker-up:
	@set -e
	@echo 'fetching 3893.4 MB of required docker images to run'
	docker-compose up -d
	@echo 'nevermind, already fetched'

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: docker-down-v
docker-down-v:
	docker-compose down -v

# docker run --rm caddy:2.6.2 caddy hash-password --plaintext 'admin'
.PHONY: caddy-hash-pw
caddy-hash-pw:
	docker run --rm caddy:2.6.2 caddy hash-password --plaintext $(ADMIN_PASSWORD)

.PHONY: mailhog-hash-pw
mailhog-hash-pw:
	mailhog bcrypt $(ADMIN_PASSWORD)

.PHONY: build-safety
build-safety:
	@set -e
	@echo 'build-safety docker images start'
	cd auth && pwd && make docker-build && echo 'done' && cd ..
	cd grpc-gateway && pwd && make docker-build && echo 'done' && cd ..
	cd mail && pwd && make docker-build && echo 'done' && cd ..
	cd safety && pwd && make docker-build && echo 'done' && cd ..
	cd user && pwd && make docker-build && echo 'done' && cd ..
	@echo 'build-safety docker images done'

.PHONY: push-safety
push-safety:
	@set -e
	@echo 'push-safety docker images start'
	docker push szczynk/grpc-safety_auth
	docker push szczynk/grpc-safety_gateway
	docker push szczynk/grpc-safety_mail
	docker push szczynk/grpc-safety_core
	docker push szczynk/grpc-safety_user
	@echo 'push-safety docker images done'

# ==============================================================================
# Fun

.PHONY: scc
scc:
	scc