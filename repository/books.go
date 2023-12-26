package repository

import (
	"context"
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/entity/book"
)

type booksImpl struct {
	db *entity.Client
}

func NewBooks(db *entity.Client) Books {
	return &booksImpl{
		db: db,
	}
}

func (b *booksImpl) ListBooks(ctx context.Context) ([]*entity.Book, error) {
	return b.db.Book.Query().
		WithAuthor().
		All(ctx)
}

func (b *booksImpl) GetBook(ctx context.Context, id int) (*entity.Book, error) {
	return b.db.Book.Query().
		Where(book.ID(id)).
		WithAuthor().
		Only(ctx)
}

func (b *booksImpl) AddBook(ctx context.Context, author *entity.User, title string, discription string) (*entity.Book, error) {
	return b.db.Book.Create().
		SetTitle(title).
		SetDiscription(discription).
		SetAuthor(author).
		Save(ctx)
}

func (b *booksImpl) DeleteBook(ctx context.Context, id int) error {
	return b.db.Book.DeleteOneID(id).Exec(ctx)
}
