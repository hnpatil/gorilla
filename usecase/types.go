package usecase

import (
	"context"
	"errors"
	"hanmantpatil/gorilla/entity"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	AuthCodeExpiredError int = 10001
	TokenInvalidError    int = 10002
	TokenExpiredError    int = 10003
)

type UsecaseError struct {
	ErrorCode int
	Message   string
}

func GetUsecaseError(err error) *UsecaseError {
	if err == nil {
		return nil
	}

	var e *UsecaseError
	if errors.As(err, &e) {
		return e
	}

	return nil
}

func (u *UsecaseError) Error() string {
	return u.Message
}

type Token struct {
	AuthToken    string    `json:"auth_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Books interface {
	AddBook(ctx Context, title string, description string) (*entity.Book, error)
	ListBooks(ctx Context) ([]*entity.Book, error)
	GetBook(ctx Context, id int) (*entity.Book, error)
	DeleteBook(ctx Context, id int) error
}

type Users interface {
	CreateUser(ctx Context, firstName string, lastName string) (*entity.User, error)
}

type Auth interface {
	GenerateAuthCode(ctx Context, identifier string) error
	GetAuthToken(ctx Context, authCode string) (*Token, error)
	RefereshToken(ctx Context, refreshToken string) (*Token, error)
	Authenticate(ctx Context, token string) (string, error)
}

type Context interface {
	GetLogger() logrus.FieldLogger
	GetUserID() string
	GetContext() context.Context
}
