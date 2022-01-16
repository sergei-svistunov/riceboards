package delete

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"
)

type reqV1 struct {
	ProjectId string `json:"project_id"`
	Id        uint32 `json:"id"`
}

var errorsV1 struct {
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
	InvalidLevel   rpc.ErrorFunc `desc:"Invalid level"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*struct{}, error) {
	projectId, err := m.db.Projects.GetIdByStrId(ctx, r.ProjectId)
	if err != nil {
		return nil, err
	}
	if projectId == 0 {
		return nil, errorsV1.UnknownProject("Unknown project")
	}

	data, err := m.db.ConfidentLevels.GetAll(ctx, []string{"id"}, model.GetAllOptions{
		Filter: expr.And(
			expr.Eq(m.db.ConfidentLevels.FieldExpr("fk_project_id"), expr.Value(projectId)),
			expr.Eq(m.db.ConfidentLevels.FieldExpr("id"), expr.Value(r.Id)),
		),
		Limit: 1,
	})
	if data.Len() == 0 {
		return nil, errorsV1.InvalidLevel("Invalid level")
	}

	if err := m.db.Storage.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := m.db.Ideas.Edit(ctx, expr.Eq(m.db.Ideas.FieldExpr("fk_confident_id"), expr.Value(r.Id)), map[string]interface{}{
			"fk_confident_id": nil,
		}); err != nil {
			return err
		}

		if err := m.db.ConfidentLevels.Delete(ctx, expr.Eq(m.db.ConfidentLevels.FieldExpr("id"), expr.Value(r.Id))); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &struct{}{}, nil
}
