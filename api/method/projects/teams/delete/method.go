package delete

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
	return `Project team deleting`
}

func (m *Method) Description(ctx context.Context) string {
	return `Delete the project team`
}
