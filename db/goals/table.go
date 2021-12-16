package goals

import (
	"github.com/go-qbit/model"
	"github.com/go-qbit/storage-mysql"
)

type Table struct {
	*mysql.BaseModel
}

var oneFloat = 1.0

func New(db *mysql.MySQL) *Table {
	return &Table{
		BaseModel: mysql.NewBaseModel(
			db, "goals", []mysql.IMysqlFieldDefinition{
				&mysql.UintField{
					Id:            "id",
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.VarCharField{
					Id:      "caption",
					Length:  255,
					NotNull: true,
				},

				&mysql.TinyUintField{
					Id:      "format",
					NotNull: true,
				},

				&mysql.FloatField{
					Id:      "divider",
					NotNull: true,
					Default: &oneFloat,
				},
			},

			[]model.IFieldDefinition{},

			mysql.BaseModelOpts{
				BaseModelOpts: model.BaseModelOpts{
					PkFieldsNames: []string{"id"},
				},
				Indexes: []mysql.Index{
					{[]string{"fk_project_id", "caption"}, true},
				},
			},
		),
	}
}
