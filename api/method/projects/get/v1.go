package get

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
)

type reqV1 struct {
	Id uint32 `json:"id"`
}

type projectV1 struct {
	Id      uint32 `json:"id"`
	Caption string `json:"caption"`
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*projectV1, error) {
	var projects []projectV1
	if err := m.db.Projects.GetAllToStruct(ctx, &projects, model.GetAllOptions{
		Filter: expr.Eq(m.db.Projects.FieldExpr("id"), expr.Value(r.Id)),
		Limit:  1,
	}); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, nil
	}

	return &projects[0], nil
}
