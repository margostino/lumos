#!make
SHELL = /bin/sh
.DEFAULT: help

-include .env .env.local .env.*.local

# Defaults
DOCKER_COMPOSE = USERID=$(shell id -u):$(shell id -g) docker-compose ${compose-files}
ALL_ENVS := local ci
env ?= local

.PHONY: help
help:
	@echo "Lumos pipeline"
	@echo ""
	@echo "Usage:"
	@echo "  wiremock                       - Run Wiremock dependency for local testing"
	@echo "  vercel.dev                     - Replicate the Vercel deployment environment locally, allowing you to test your Serverless Functions, without requiring you to deploy each time a change is made"
	@echo "  run                            - Start main runner for local testing"
	@echo "  docker.wait                    - Waits until all docker containers have exited successfully and/or are healthy. Timeout: 180 seconds"
	@echo "  docker.logs                    - Generate one log file per each service running in docker-compose"
	@echo ""
	@echo "  ** The following tasks receive an env parameter to determine the environment they are being executed in. Default env=${env}, possible env values: ${ALL_ENVS}:"
	@echo "  docker.run.all                 - Run QuarkBox service and all it's dependencies with docker-compose (default env=${env})"
	@echo "  docker.run.dependencies        - Run only QuarkBox dependencies with docker-compose (default env=${env})". Note that `build` might need to be executed prior.
	@echo "  docker.stop                    - Stop and remove all running containers from this project using docker-compose down (default env=${env})"
	@echo ""
	@echo "Project-level environment variables are set in .env file:"
	@echo "  VERCEL=1"
	@echo "  VERCEL_ENV=development"
	@echo "  APP_DATA_BASE_URL=http://localhost:10000/lumos/data"
	@echo ""
	@echo "Note: Store protected environment variables in .env.local or .env.*.local"
	@echo ""

.PHONY: vercel.dev
vercel.dev:
	make v.stop
	vercel dev

.PHONY: run
run:
	go run runner/main.go

.PHONY: wiremock
wiremock: d.compose.down
	make d.compose.up
	make docker.wait
	docker-compose up -d
	docker-compose ps

.PHONY: docker.run.all
docker.run.all: d.compose.down
	make d.compose.up
	make docker.wait
	docker-compose ps

.PHONY: docker.run.dependencies
docker.run.dependencies: d.compose.down
	make d.compose.up
	make docker.wait
	docker-compose up -d
	docker-compose ps

.PHONY: docker.stop
docker.stop: d.compose.down

.PHONY: docker.logs
docker.logs:
	./bin/docker-logs

.PHONY: docker.wait
docker.wait:
	./bin/docker-wait

.PHONY: d.compose.up
d.compose.up:
	$(call DOCKER_COMPOSE) up -d --remove-orphans --build

.PHONY: d.compose.down
d.compose.down:
	$(call DOCKER_COMPOSE) down -v || true
	$(call DOCKER_COMPOSE) rm --force || true
	docker rm "$(docker ps -a -q)" -f || true

.PHONY: v.stop
v.stop:
	./bin/vercel-stop

