package projects

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rbac"
	"github.com/go-qbit/storage-mysql"

	"riceboards/authctx"
)

var (
	PermGroup = rbac.NewPermissionsGroup("project", "Model project")

	PermView = PermGroup.NewPermission("view", "View projects")
)

type Table struct {
	*mysql.BaseModel
}

func New(db *mysql.MySQL) *Table {
	return &Table{
		BaseModel: mysql.NewBaseModel(
			db, "projects", []mysql.IMysqlFieldDefinition{
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
			},

			[]model.IFieldDefinition{},

			mysql.BaseModelOpts{
				BaseModelOpts: model.BaseModelOpts{
					PkFieldsNames: []string{"id"},
					DefaultFilter: func(ctx context.Context, m model.IModel) (model.IExpression, error) {
						curUid, err := authctx.GetCurUserId(ctx)
						if err != nil {
							return nil, err
						}

						if !rbac.HasPermission(ctx, PermView) {
							return expr.Value(false), nil
						}

						tProjectUsers := m.GetRelation("projects_users").ExtModel
						return expr.Or(
							expr.Eq(m.FieldExpr("fk_owner_id"), expr.Value(curUid)),
							expr.Any(m, tProjectUsers, expr.Eq(tProjectUsers.FieldExpr("fk_user_id"), expr.Value(curUid))),
						), nil
					},
				},
				Indexes: []mysql.Index{
					{[]string{"caption"}, true},
				},
			},
		),
	}
}
