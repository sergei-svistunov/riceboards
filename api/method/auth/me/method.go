package me

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

func (m *Method) Caption(context.Context) string {
	return `Me`
}

func (m *Method) Description(context.Context) string {
	return `Returns current user`
}
