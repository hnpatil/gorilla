package usecase

import (
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/repository"
)

type booksImpl struct {
	books repository.Books
	users repository.Users
}

func NewBooks(books repository.Books, users repository.Users) Books {
	return &booksImpl{
		books: books,
		users: users,
	}
}

func (b *booksImpl) ListBooks(ctx Context) ([]*entity.Book, error) {
	return b.books.ListBooks(ctx.GetContext())
}

func (b *booksImpl) GetBook(ctx Context, id int) (*entity.Book, error) {
	return b.books.GetBook(ctx.GetContext(), id)
}

func (b *booksImpl) AddBook(ctx Context, title string, description string) (*entity.Book, error) {
	usr, err := b.users.GetUser(ctx.GetContext(), ctx.GetUserID())
	if err != nil {
		return nil, err
	}

	return b.books.AddBook(ctx.GetContext(), usr, title, description)
}

func (b *booksImpl) DeleteBook(ctx Context, id int) error {
	return b.books.DeleteBook(ctx.GetContext(), id)
}
