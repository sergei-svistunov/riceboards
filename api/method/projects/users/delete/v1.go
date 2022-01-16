package delete

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"
)

type reqV1 struct {
	ProjectId string `json:"project_id"`
	EMail     string `json:"e_mail"`
}

var errorsV1 struct {
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
	UnknownUser    rpc.ErrorFunc `desc:"Unknown user"`
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

	var users []struct {
		Id uint32
	}
	if err := m.db.Users.GetAllToStruct(ctx, &users, model.GetAllOptions{
		Filter: expr.Eq(m.db.Users.FieldExpr("email"), expr.Value(r.EMail)),
		Limit:  1,
	}); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errorsV1.UnknownUser("Unknown user")
	}

	if err := m.db.ProjectsUsers.Delete(ctx, expr.And(
		expr.Eq(m.db.ProjectsUsers.FieldExpr("fk_project_id"), expr.Value(projectId)),
		expr.Eq(m.db.ProjectsUsers.FieldExpr("fk_user_id"), expr.Value(users[0].Id)),
	)); err != nil {
		return nil, err
	}

	return &struct{}{}, nil
}
