package confident_levels

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
			db, "confident_levels", []mysql.IMysqlFieldDefinition{
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

				&mysql.FloatField{
					Id:      "weight",
					NotNull: true,
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
