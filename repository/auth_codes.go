package repository

import (
	"context"
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/entity/authcode"
	"time"
)

type authCodesImpl struct {
	db *entity.Client
}

func NewAuthCodes(db *entity.Client) AuthCodes {
	return &authCodesImpl{
		db: db,
	}
}

func (a *authCodesImpl) CreateAuthCode(ctx context.Context, identifier string, code string, expiresAt *time.Time) (*entity.AuthCode, error) {
	return a.db.AuthCode.Create().
		SetCode(code).
		SetIdentifier(identifier).
		SetExpiresAt(*expiresAt).
		Save(ctx)
}

func (a *authCodesImpl) GetAuthCode(ctx context.Context, code string) (*entity.AuthCode, error) {
	return a.db.AuthCode.Query().Where(authcode.Code(code)).Only(ctx)
}
