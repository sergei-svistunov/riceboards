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
	InvalidTean    rpc.ErrorFunc `desc:"Invalid team"`
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

	data, err := m.db.Teams.GetAll(ctx, []string{"id"}, model.GetAllOptions{
		Filter: expr.And(
			expr.Eq(m.db.Teams.FieldExpr("fk_project_id"), expr.Value(projectId)),
			expr.Eq(m.db.Teams.FieldExpr("id"), expr.Value(r.Id)),
		),
		Limit: 1,
	})
	if data.Len() == 0 {
		return nil, errorsV1.InvalidTean("Invalid team")
	}

	if err := m.db.Storage.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := m.db.IdeaTeams.Delete(ctx, expr.Eq(m.db.IdeaTeams.FieldExpr("fk_team_id"), expr.Value(r.Id))); err != nil {
			return err
		}

		if err := m.db.Teams.Delete(ctx, expr.Eq(m.db.Teams.FieldExpr("id"), expr.Value(r.Id))); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &struct{}{}, nil
}
