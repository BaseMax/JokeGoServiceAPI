TEST_COMPOSER=docker-compose -f docker-compose-test.yml

all: build up down

up:
	docker-compose up

build: down
	docker-compose build

down:
	docker-compose down

test: test-build test-up test-down

test-up:
	${TEST_COMPOSER} up

test-build: test-down
	${TEST_COMPOSER} build

test-down:
	${TEST_COMPOSER} down

test-run-db:
	${TEST_COMPOSER} up db -d

run-db:
	docker-compose up db -d

run-local:
	export DB_HOSTNAME=localhost; go run .

test-local:
	export DB_HOSTNAME=localhost; go test -v ./... -cover -coverprofile=/tmp/cover.out
	go tool cover -func=/tmp/cover.out
