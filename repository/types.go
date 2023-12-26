package repository

import (
	"context"
	"hanmantpatil/gorilla/entity"
	"time"
)

type Books interface {
	AddBook(ctx context.Context, author *entity.User, title string, discription string) (*entity.Book, error)
	ListBooks(ctx context.Context) ([]*entity.Book, error)
	GetBook(ctx context.Context, id int) (*entity.Book, error)
	DeleteBook(ctx context.Context, id int) error
}

type Users interface {
	CreateUser(ctx context.Context, firstName string, lastName string, email string) (*entity.User, error)
	GetUser(ctx context.Context, email string) (*entity.User, error)
}

type AuthCodes interface {
	CreateAuthCode(ctx context.Context, identifier string, code string, expiresAt *time.Time) (*entity.AuthCode, error)
	GetAuthCode(ctx context.Context, code string) (*entity.AuthCode, error)
}
