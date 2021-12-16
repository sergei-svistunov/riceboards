package ideas

import (
	"github.com/go-qbit/model"
	"github.com/go-qbit/storage-mysql"
)

type Table struct {
	*mysql.BaseModel
}

func New(db *mysql.MySQL) *Table {
	return &Table{
		BaseModel: mysql.NewBaseModel(
			db, "ideas", []mysql.IMysqlFieldDefinition{
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

				&mysql.BooleanField{
					Id:      "ready_for_dev",
					NotNull: true,
				},

				&mysql.BigUintField{
					Id:     "reach",
					Length: 255,
				},

				&mysql.TextField{
					Id: "reach_comment",
				},

				&mysql.VarCharField{
					Id:     "link",
					Length: 255,
				},

				&mysql.TextField{
					Id: "confident_comment",
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
