package repository

import (
	"clickhouse_postgres/domain"
	"context"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type authorRepositoryImpl struct {
	db *pg.DB
}

func NewAuthorRepository(db *pg.DB) domain.AuthorRepository {
	return &authorRepositoryImpl{
		db: db,
	}
}

func (a *authorRepositoryImpl) InitializeAuthorTable(ctx context.Context) error {
	for _, model := range []interface{}{(*domain.Author)(nil)} {
		err := a.db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil

}

func (a *authorRepositoryImpl) GetTable() string {
	return "author"
}

func (a *authorRepositoryImpl) Save(ctx context.Context, author *domain.Author) error {
	err := a.db.Insert(author)

	return err
}

func (a *authorRepositoryImpl) GetByID(ctx context.Context, id int64) (*domain.Author, error) {
	author := &domain.Author{
		Id: id,
	}

	if err := a.db.Select(author); err != nil {
		return nil, err
	}

	return author, nil
}

func (a *authorRepositoryImpl) DropAuthorTable(ctx context.Context) error {
	err := a.db.DropTable(&domain.Author{}, &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})

	return err
}

func (a *authorRepositoryImpl) Close() {
	a.db.Close()
}
