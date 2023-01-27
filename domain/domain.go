package domain

import "context"

type Author struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

type AuthorRepository interface {
	Save(ctx context.Context, author *Author) error
	GetByID(ctx context.Context, id int64) (*Author, error)

	InitializeAuthorTable(ctx context.Context) error
	DropAuthorTable(ctx context.Context) error
	Close()
}
