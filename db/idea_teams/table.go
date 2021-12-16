package idea_teams

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
			db, "idea_teams", []mysql.IMysqlFieldDefinition{
				&mysql.UintField{
					Id:            "id",
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.FloatField{
					Id:      "capacity",
					NotNull: true,
				},

				&mysql.TextField{
					Id: "comment",
				},
			},

			[]model.IFieldDefinition{},

			mysql.BaseModelOpts{
				BaseModelOpts: model.BaseModelOpts{
					PkFieldsNames: []string{"id"},
				},
				Indexes: []mysql.Index{
					{[]string{"fk_idea_id", "fk_team_id"}, true},
				},
			},
		),
	}
}
