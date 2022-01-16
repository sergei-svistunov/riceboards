package add

import (
	"context"
	"errors"

	"github.com/go-qbit/model"
	"github.com/go-qbit/rpc"

	"riceboards/authctx"
	"riceboards/db"
)

type reqV1 struct {
	ProjectId string `json:"project_id"`
	Caption   string `json:"caption"`
}

var errorsV1 struct {
	Unauthorized   rpc.ErrorFunc `desc:"Unauthorized"`
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
	EmptyCaption   rpc.ErrorFunc `desc:"Caption must be filled"`
	AlreadyExists  rpc.ErrorFunc `desc:"Team is already exists in the project"`
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

	if db.FieldIsEmpty(r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}

	_, err = m.db.Teams.AddFromStructs(ctx, []struct {
		FkProjectId uint32
		FkOwnerId   uint32
		Caption     string
	}{
		{projectId, curUserId, r.Caption},
	}, model.AddOptions{})
	if err != nil {
		if db.IsDuplicateEntryErr(err) {
			return nil, errorsV1.AlreadyExists("The team is already added to the project")
		}
		return nil, err
	}

	return &struct{}{}, nil
}
