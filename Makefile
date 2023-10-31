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
	curl --location --request POST 'http://localhost:8000/login' --header 'Content-Type: application/json' -d '{"id": "string", "password": "Секретная строка"}'