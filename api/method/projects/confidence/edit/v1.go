package edit

import (
	"context"
	"errors"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"

	"riceboards/authctx"
	"riceboards/db"
)

type reqV1 struct {
	ProjectId string  `json:"project_id"`
	Id        uint32  `json:"id"`
	Caption   string  `json:"caption"`
	Weight    float64 `json:"weight"`
}

var errorsV1 struct {
	Unauthorized   rpc.ErrorFunc `desc:"Unauthorized"`
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
	InvalidLevel   rpc.ErrorFunc `desc:"Invalid level"`
	EmptyCaption   rpc.ErrorFunc `desc:"Caption must be filled"`
	AlreadyExists  rpc.ErrorFunc `desc:"Goal is already exists in the project"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*struct{}, error) {
	curUserId, err := authctx.GetCurUserId(ctx)
	if err != nil {
		if errors.Is(err, authctx.ErrUnauthorized) {
			return nil, errorsV1.Unauthorized("Unauthorized")
		}
		return nil, err
	}

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

	if db.FieldIsEmpty(r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}

	if err := m.db.ConfidentLevels.Edit(ctx, expr.Eq(m.db.ConfidentLevels.FieldExpr("id"), expr.Value(r.Id)), map[string]interface{}{
		"fk_owner_id": curUserId,
		"caption":     r.Caption,
		"weight":      r.Weight,
	}); err != nil {
		if db.IsDuplicateEntryErr(err) {
			return nil, errorsV1.AlreadyExists("The level is already added to the project")
		}
		return nil, err
	}

	return &struct{}{}, nil
}
