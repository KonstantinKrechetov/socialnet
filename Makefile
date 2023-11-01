all: install generate run

install:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate:
	oapi-codegen --config=./api/server.cfg.yaml ./api/openapi.json

run:
	go run ./main.go

check:
	open http://localhost:8000/swagger/

check_curl:
	curl --location --request GET 'http://localhost:8000/_healthcheck/'





########################
docker_run_postgres:
	docker run --name postgres_cont -e POSTGRES_DB=postgres -e POSTGRES_PASSWORD=postgres -p 5433:5432 -d --rm postgres:13

docker_connect_to_container:
	docker exec -it postgres_cont psql -h localhost -U postgres postgres -W




##########
env:
	cp .env.example .env