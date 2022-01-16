package add

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"

	"riceboards/db"
)

type reqV1 struct {
	ProjectId string `json:"project_id"`
	EMail     string `json:"e_mail"`
}

var errorsV1 struct {
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
	UnknownUser    rpc.ErrorFunc `desc:"Unknown user"`
	AlreadyExists  rpc.ErrorFunc `desc:"User is already exists in the project"`
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

	_, err = m.db.ProjectsUsers.AddFromStructs(ctx, []struct {
		FkProjectId uint32
		FkUserId    uint32
	}{
		{projectId, users[0].Id},
	}, model.AddOptions{})
	if err != nil {
		if db.IsDuplicateEntryErr(err) {
			return nil, errorsV1.AlreadyExists("The user is already added to the project")
		}
		return nil, err
	}

	return &struct{}{}, nil
}
