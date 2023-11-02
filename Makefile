
#### RUN & MIGRATE POSTGRES
SETUP_DB: postgres create_db migrate_up

postgres:
	docker run --name postgres13 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5433:5432 -d --rm postgres:13-alpine

#connect:
#	docker exec -it postgres psql -h localhost -U root postgres13 -W

create_db:
	docker exec -it postgres13 createdb --username=root --owner=root social_net

drop_db:
	docker exec -it postgres13 dropdb social_net

migrate_up:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/social_net?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/social_net?sslmode=disable" -verbose down


############## SETUP SERVICE
SETUP_SERVICE: run

tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

codegen:
	oapi-codegen --config=./api/server.cfg.yaml ./api/openapi.json

env:
	cp .env.example .env

run:
	go run ./main.go

check:
	open http://localhost:8080/swagger/

check_curl:
	curl --location --request GET 'http://localhost:8080/_healthcheck/'



#######
docker_build:
	docker build -t social_net .

docker_remove_stopped_containers:
	docker compose rm -f

docker_compose:
	docker compose up --build -d