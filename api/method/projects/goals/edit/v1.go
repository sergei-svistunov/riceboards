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
	Format    uint8   `json:"format"`
	Divider   float64 `json:"divider"`
}

var errorsV1 struct {
	Unauthorized   rpc.ErrorFunc `desc:"Unauthorized"`
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
	InvalidGoal    rpc.ErrorFunc `desc:"Invalid goal"`
	EmptyCaption   rpc.ErrorFunc `desc:"Caption must be filled"`
	AlreadyExists  rpc.ErrorFunc `desc:"Goal is already exists in the project"`
	InvalidFormat  rpc.ErrorFunc `desc:"Invalid format"`
	InvalidDivider rpc.ErrorFunc `desc:"Invalid divider"`
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

	data, err := m.db.Goals.GetAll(ctx, []string{"id"}, model.GetAllOptions{
		Filter: expr.And(
			expr.Eq(m.db.Goals.FieldExpr("fk_project_id"), expr.Value(projectId)),
			expr.Eq(m.db.Goals.FieldExpr("id"), expr.Value(r.Id)),
		),
		Limit: 1,
	})
	if data.Len() == 0 {
		return nil, errorsV1.InvalidGoal("Invalid goal")
	}

	if db.FieldIsEmpty(r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}

	if r.Format > 2 {
		return nil, errorsV1.InvalidFormat("Invalid format")
	}

	if r.Divider <= 0 {
		return nil, errorsV1.InvalidDivider("The divider must be great than 0")
	}

	if err := m.db.Goals.Edit(ctx, expr.Eq(m.db.Goals.FieldExpr("id"), expr.Value(r.Id)), map[string]interface{}{
		"fk_owner_id": curUserId,
		"caption":     r.Caption,
		"format":      r.Format,
		"divider":     r.Divider,
	}); err != nil {
		if db.IsDuplicateEntryErr(err) {
			return nil, errorsV1.AlreadyExists("The goal is already added to the project")
		}
		return nil, err
	}

	return &struct{}{}, nil
}
