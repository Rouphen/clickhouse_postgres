package main

import (
	"context"

	"fmt"
	"time"

	"clickhouse_postgres/domain"
	repoclickhouse "clickhouse_postgres/repository/clickhouse"
	repopostgres "clickhouse_postgres/repository/postgres"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/go-pg/pg"
)

func main() {
	//repo, ctx := initializeClickHouseRepository()
	repo, ctx := initializePostgresRepository()
	defer repo.Close()
	// create
	err := repo.InitializeAuthorTable(ctx)
	if err != nil {
		fmt.Println("failed to create table with : " + err.Error())
		return
	}

	// insert
	authorRob := &domain.Author{
		Id:   1,
		Name: "Rob",
		Age:  42,
	}
	repo.Save(ctx, authorRob)
	if err != nil {
		fmt.Println("failed to insert with : " + err.Error())
		return
	}

	// query
	author, errQyery := repo.GetByID(ctx, authorRob.Id)
	if errQyery != nil {
		fmt.Println("failed to query with : " + errQyery.Error())
		return
	}
	fmt.Println(author)

	//drop
	err = repo.DropAuthorTable(ctx)
	if err != nil {
		fmt.Println("failed to drop with : " + err.Error())
		return
	}
}

func initializeClickHouseRepository() (domain.AuthorRepository, context.Context) {
	//docker run -d --name clickhouse-server --ulimit nofile=262144:262144 -p 9000:9000 yandex/clickhouse-server:lastest
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		panic("failed to connect with : " + err.Error())
	}

	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}))

	repo := repoclickhouse.NewAuthorRepository(conn)

	return repo, ctx
}

func initializePostgresRepository() (domain.AuthorRepository, context.Context) {
	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "123456",
		Database: "postgres",
	})

	repo := repopostgres.NewAuthorRepository(db)

	return repo, context.Background()
}
