package google

import (
	"context"

	webAuth "github.com/go-qbit/web/auth"

	"riceboards/db"
)

type Method struct {
	db   *db.Db
	auth *webAuth.Auth
}

func New(db *db.Db, auth *webAuth.Auth) *Method {
	return &Method{
		db:   db,
		auth: auth,
	}
}

func (m *Method) Caption(ctx context.Context) string {
	return `Google auth`
}

func (m *Method) Description(ctx context.Context) string {
	return `Sign in through Google`
}
