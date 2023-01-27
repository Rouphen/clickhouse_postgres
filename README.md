# clickhouse_postgres
a simple demo for using click house and Postgres, but not unit test

## step 1
docker run -d --name clickhouse-server --ulimit nofile=262144:262144 -p 9000:9000 yandex/clickhouse-server:lastest

## step 2
docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=123456 -v postgres:/var/lib/postgresql/data postgres:14

## step 3
//repo, ctx := initializeClickHouseRepository()
repo, ctx := initializePostgresRepository()

## step 4
go run main.go
