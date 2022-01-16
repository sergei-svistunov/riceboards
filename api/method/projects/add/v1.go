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

	var projectId uint32
	if err := m.db.Storage.DoInTransaction(ctx, func(ctx context.Context) error {
		pks, err := m.db.Projects.AddFromStructs(ctx, []struct {
			FkOwnerId uint32
			StrId     string
			Caption   string
		}{
			{curUserId, r.Id, r.Caption},
		}, model.AddOptions{})
		if err != nil {
			if db.IsDuplicateEntryErr(err) {
				return errorsV1.AlreadyExists("The project ID is already busy")
			}

			return err
		}
		projectId = pks.Data()[0][0].(uint32)

		if _, err := m.db.ConfidentLevels.AddFromStructs(ctx, []struct {
			FkProjectId uint32
			FkOwnerId   uint32
			Weight      float64
			Caption     string
		}{
			{projectId, curUserId, 1, "Self conviction"},
			{projectId, curUserId, 2, "Others opinions"},
			{projectId, curUserId, 3, "Estimates & plans"},
			{projectId, curUserId, 4, "Anecdotal evidence"},
			{projectId, curUserId, 5, "Market data"},
			{projectId, curUserId, 6, "User/customer evidence"},
			{projectId, curUserId, 8, "Test results"},
			{projectId, curUserId, 10, "Launch data"},
		}, model.AddOptions{}); err != nil {
			return err
		}

		if _, err := m.db.Goals.AddFromStructs(ctx, []struct {
			FkProjectId uint32
			FkOwnerId   uint32
			Caption     string
			Format      uint8
			Divider     float64
		}{
			{projectId, curUserId, "Money", 2, 1_000_000},
		}, model.AddOptions{}); err != nil {
			return err
		}

		if _, err := m.db.Teams.AddFromStructs(ctx, []struct {
			FkProjectId uint32
			FkOwnerId   uint32
			Caption     string
		}{
			{projectId, curUserId, "Development"},
		}, model.AddOptions{}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &projectV1{
		Id: projectId,
	}, nil
}
