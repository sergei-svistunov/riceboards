package edit

import (
	"context"

	"riceboards/db"
)

type Method struct {
	db *db.Db
}

func New(db *db.Db) *Method {
	return &Method{
		db: db,
	}
}

func (m *Method) Caption(ctx context.Context) string {
	return `Project goal editing`
}

func (m *Method) Description(ctx context.Context) string {
	return `Edit the project goal`
}
