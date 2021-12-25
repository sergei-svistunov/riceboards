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
	ProjectId uint32  `json:"project_id"`
	Caption   string  `json:"caption"`
	Comment   *string `json:"comment"`
}

type ideaV1 struct {
	Id uint32 `json:"id"`
}

var errorsV1 struct {
	Unauthorized  rpc.ErrorFunc `desc:"Unauthorized"`
	EmptyCaption  rpc.ErrorFunc `desc:"Caption must be filled"`
	AlreadyExists rpc.ErrorFunc `desc:"Caption is already busy"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*ideaV1, error) {
	curUserId, err := authctx.GetCurUserId(ctx)
	if err != nil {
		if errors.Is(err, authctx.ErrUnauthorized) {
			return nil, errorsV1.Unauthorized("Unauthorized")
		}
		return nil, err
	}

	if db.FieldIsEmpty(r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}

	pks, err := m.db.Ideas.AddFromStructs(ctx, []struct {
		FkOwnerId   uint32
		FkProjectId uint32
		Caption     string
		Comment     *string
		ReadyForDev bool
	}{
		{curUserId, r.ProjectId, r.Caption, r.Comment, false},
	}, model.AddOptions{})
	if err != nil {
		if db.IsDuplicateEntryErr(err) {
			return nil, errorsV1.AlreadyExists("Caption is already busy")
		}

		return nil, err
	}

	return &ideaV1{
		Id: pks.Data()[0][0].(uint32),
	}, nil
}
