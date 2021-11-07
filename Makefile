include .env

DB_MIGRATION_DIR := database/migrations
DB_MYSQL_DSN := mysql://${DB_USER}:${DB_PASS}@tcp'('${DB_HOST}:${DB_PORT}')'/${DB_NAME}

# run executes the main.go file
run:
	go run ./cmd/api/main.go

# build compiles the project to /bin
build:
	go build -o ./cmd/api/ ./cmd/api/main.go

# dev runs a dev server
dev:
	CompileDaemon -exclude=*_test.go -build="make build" -command="./cmd/api/main"

# migrateup migrates up the mysql db
migrateup:
	migrate -path $(DB_MIGRATION_DIR) -database $(DB_MYSQL_DSN) -verbose up

# migratedown migrates down the mysql db
migratedown:
	migrate -path $(DB_MIGRATION_DIR) -database $(DB_MYSQL_DSN) -verbose down
