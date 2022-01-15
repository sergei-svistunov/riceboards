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
	Id      string `json:"id"`
	Caption string `json:"caption"`
}

type projectV1 struct {
	Id uint32 `json:"id"`
}

var errorsV1 struct {
	Unauthorized  rpc.ErrorFunc `desc:"Unauthorized"`
	EmptyId       rpc.ErrorFunc `desc:"Id must be filled"`
	EmptyCaption  rpc.ErrorFunc `desc:"Caption must be filled"`
	AlreadyExists rpc.ErrorFunc `desc:"Caption is already busy"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*projectV1, error) {
	curUserId, err := authctx.GetCurUserId(ctx)
	if err != nil {
		if errors.Is(err, authctx.ErrUnauthorized) {
			return nil, errorsV1.Unauthorized("Unauthorized")
		}
		return nil, err
	}

	if db.FieldIsEmpty(r.Id) {
		return nil, errorsV1.EmptyId("Id must be filled")
	}

	if db.FieldIsEmpty(r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}

	pks, err := m.db.Projects.AddFromStructs(ctx, []struct {
		FkOwnerId uint32
		StrId     string
		Caption   string
	}{
		{curUserId, r.Id, r.Caption},
	}, model.AddOptions{})
	if err != nil {
		if db.IsDuplicateEntryErr(err) {
			return nil, errorsV1.AlreadyExists("The project ID is already busy")
		}

		return nil, err
	}

	return &projectV1{
		Id: pks.Data()[0][0].(uint32),
	}, nil
}
