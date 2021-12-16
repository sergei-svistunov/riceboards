package users

import (
	"github.com/go-qbit/model"
	"github.com/go-qbit/rbac"
	"github.com/go-qbit/storage-mysql"
)

var (
	PermGroupUser = rbac.NewPermissionsGroup("user", "Model user")
)

type Table struct {
	*mysql.BaseModel
}

func New(db *mysql.MySQL) *Table {
	return &Table{
		BaseModel: mysql.NewBaseModel(
			db, "users", []mysql.IMysqlFieldDefinition{
				&mysql.UintField{
					Id:            "id",
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.VarCharField{
					Id:      "fullname",
					Length:  255,
					NotNull: true,
				},

				&mysql.VarCharField{
					Id:      "email",
					Length:  255,
					NotNull: true,
				},

				&mysql.VarCharField{
					Id:      "avatar_url",
					Length:  255,
					NotNull: true,
				},

				&mysql.BooleanField{
					Id:      "is_admin",
					NotNull: true,
				},
			},

			[]model.IFieldDefinition{},

			mysql.BaseModelOpts{
				BaseModelOpts: model.BaseModelOpts{
					PkFieldsNames: []string{"id"},
				},
				Indexes: []mysql.Index{
					{[]string{"email"}, true},
				},
			},
		),
	}
}
