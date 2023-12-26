package usecase

import (
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/repository"
)

type usersImpl struct {
	repo repository.Users
}

func NewUsers(repo repository.Users) Users {
	return &usersImpl{
		repo: repo,
	}
}

func (u *usersImpl) CreateUser(ctx Context, firstName string, lastName string) (*entity.User, error) {
	return u.repo.CreateUser(ctx.GetContext(), firstName, lastName, ctx.GetUserID())
}
