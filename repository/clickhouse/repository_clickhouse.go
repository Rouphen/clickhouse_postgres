package repository

import (
	"context"
	"errors"
	"fmt"

	"clickhouse_postgres/domain"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type authorRepositoryImpl struct {
	conn driver.Conn
}

func NewAuthorRepository(conn driver.Conn) domain.AuthorRepository {
	return &authorRepositoryImpl{
		conn: conn,
	}
}

func (a *authorRepositoryImpl) InitializeAuthorTable(ctx context.Context) error {
	err := a.conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS author (
			id   Int64,
			name String,
			age  Int64
		) engine=Memory
	`)

	return err
}

func (a *authorRepositoryImpl) GetTable() string {
	return "author"
}

func (a *authorRepositoryImpl) Save(ctx context.Context, author *domain.Author) error {
	insertSQL := fmt.Sprintf("INSERT INTO author VALUES (%d, '%s', %d)", author.Id, author.Name, author.Age)

	return a.conn.Exec(ctx, insertSQL)
}

func (a *authorRepositoryImpl) GetByID(ctx context.Context, id int64) (*domain.Author, error) {
	querySQL := fmt.Sprintf("SELECT * FROM author WHERE id=%d", id)
	row := a.conn.QueryRow(ctx, querySQL)
	if row == nil {
		return nil, errors.New("can not found")
	}

	author := &domain.Author{}
	var authorId, authorAge int64
	var authorName string
	if err := row.Scan(&authorId, &authorName, &authorAge); err != nil {
		return nil, err
	}

	author.Id = authorId
	author.Name = authorName
	author.Age = authorAge

	return author, nil
}

func (a *authorRepositoryImpl) DropAuthorTable(ctx context.Context) error {
	return a.conn.Exec(ctx, "DROP TABLE IF EXISTS author")
}

func (a *authorRepositoryImpl) Close() {
	a.conn.Close()
}
