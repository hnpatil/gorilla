package repository

import (
	"context"
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/entity/user"
)

type usersImpl struct {
	db *entity.Client
}

func NewUsers(db *entity.Client) Users {
	return &usersImpl{
		db: db,
	}
}

func (u *usersImpl) CreateUser(ctx context.Context, firstName string, lastName string, email string) (*entity.User, error) {
	return u.db.User.Create().
		SetFirstName(firstName).
		SetLastName(lastName).
		SetEmail(email).
		Save(ctx)
}

func (u *usersImpl) GetUser(ctx context.Context, email string) (*entity.User, error) {
	return u.db.User.Query().
		Where(user.Email(email)).
		Only(ctx)
}
