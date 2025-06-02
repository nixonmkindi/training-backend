# Migration
## Package
github.com/golang-migrate/migrate/v4

## Installation
Install the package using this command
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## Migration
Create migration sql using this command
migrate create -ext sql [name] 

Migrate the file by typing the following
export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/soso?sslmode=disable'
migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2

migrate -database ${POSTGRESQL_URL} -path migrations up

