# Store api

## Requirements

[go](https://golang.org/)

[CompileDaemon](https://github.com/githubnemo/CompileDaemon):
`go get github.com/githubnemo/CompileDaemon`

[Migrate](https://github.com/golang-migrate/migrate):
`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.0`

## Usage

### Environment variables
Create a `.env` file on the root project based on `.env.example`.

### Migrations

Create a new migration: `migrate -ext sql -dir ./migrations -seq <name>`.

Migrate up: `make migrateup`.

Migrate down: `make migratedown`.
