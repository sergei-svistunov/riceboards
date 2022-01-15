package get

import (
	"context"
	"errors"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"

	"riceboards/authctx"
)

type reqV1 struct {
	Id string `json:"id"`
}

type projectV1 struct {
	Id      string `json:"id" field:"str_id"`
	Caption string `json:"caption"`
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*projectV1, error) {
	if _, err := authctx.GetCurUserId(ctx); err != nil {
		if errors.Is(err, authctx.ErrUnauthorized) {
			return nil, nil
		}
		return nil, err
	}

	var projects []projectV1
	if err := m.db.Projects.GetAllToStruct(ctx, &projects, model.GetAllOptions{
		Filter: expr.Eq(m.db.Projects.FieldExpr("str_id"), expr.Value(r.Id)),
		Limit:  1,
	}); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, nil
	}

	return &projects[0], nil
}
