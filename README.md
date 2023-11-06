### Requirements
* Go
* Docker 
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### Usage

Copy the `env.example` file to a `.env` file.
```bash
$ cp .env.example .env
```
Update the postgres variables declared in the new `.env` to match your preference.
There's a handy guide on the [Postgres' DockerHub](https://hub.docker.com/_/postgres).

Build and start the services with:
```bash
$ docker-compose up --build
```

### Migrations
Migrations are applied within the code after you run the application.

[//]: # (The database migration files are in `db/migrations` so feel free to simply source them )

[//]: # (directly at _localhost:5433_ with your credentials from _.env_. )

[//]: # (Alternatively, you can apply them using [golang-migrate/migrate]&#40;https://github.com/golang-migrate/migrate&#41; by running:)

[//]: # (```bash)

[//]: # ($ export POSTGRESQL_URL="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5433/$POSTGRES_DB?sslmode=disable")

[//]: # ($ migrate -database ${POSTGRESQL_URL} -path db/migrations up)

[//]: # (```)

[//]: # (_**NOTE:** Remember to replace the `$POSTGRES_*` variables with their actual values_)

### Development
After making your changes, you can rebuild the `server` service by running the commands below
```bash
$ docker-compose stop server
$ docker-compose build server
$ docker-compose up --no-start server
```


```bash
$ docker-compose start server
```
