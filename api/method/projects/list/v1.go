package list

import (
	"context"

	"github.com/go-qbit/model"
)

type reqV1 struct {
}

type projectV1 struct {
	Id      uint32 `json:"id"`
	Caption string `json:"caption"`
}

func (m *Method) V1(ctx context.Context, r *reqV1) ([]projectV1, error) {
	var projects []projectV1
	if err := m.db.Projects.GetAllToStruct(ctx, &projects, model.GetAllOptions{
		OrderBy: []model.Order{
			{"caption", false},
		},
	}); err != nil {
		return nil, err
	}

	return projects, nil
}
