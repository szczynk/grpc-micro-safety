.PHONY: docker-build
docker-build:
	docker build --tag szczynk/postgres_cron:15.1 --progress=plain .

.PHONY: docker-run
docker-run:
	docker run --name postgres_cron -d szczynk/postgres_cron:15.1

.PHONY: docker-push
docker-push:
	docker push szczynk/postgres_cron:15.1