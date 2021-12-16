package options

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
)

type reqV1 struct {
	ProjectId uint32 `json:"project_id"`
}

type optionsV1 struct {
	ConfidentLevels []ConfidentV1 `json:"confident_levels"`
	Goals           []goalV1      `json:"goals"`
	Teams           []teamV1      `json:"teams"`
}

type ConfidentV1 struct {
	Id      uint32  `json:"id"`
	Caption string  `json:"caption"`
	Weight  float64 `json:"weight"`
}

type goalV1 struct {
	Id      uint32  `json:"id"`
	Caption string  `json:"caption"`
	Format  uint8   `json:"format"`
	Divider float64 `json:"divider"`
}

type teamV1 struct {
	Id      uint32 `json:"id"`
	Caption string `json:"caption"`
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*optionsV1, error) {
	var confidentLevels []ConfidentV1
	if err := m.db.ConfidentLevels.GetAllToStruct(ctx, &confidentLevels, model.GetAllOptions{
		Filter: expr.Eq(m.db.ConfidentLevels.FieldExpr("fk_project_id"), expr.Value(r.ProjectId)),
		OrderBy: []model.Order{
			{"weight", false},
		},
	}); err != nil {
		return nil, err
	}

	var goals []goalV1
	if err := m.db.Goals.GetAllToStruct(ctx, &goals, model.GetAllOptions{
		Filter: expr.Eq(m.db.Goals.FieldExpr("fk_project_id"), expr.Value(r.ProjectId)),
		OrderBy: []model.Order{
			{"caption", false},
		},
	}); err != nil {
		return nil, err
	}

	var teams []teamV1
	if err := m.db.Teams.GetAllToStruct(ctx, &teams, model.GetAllOptions{
		Filter: expr.Eq(m.db.Teams.FieldExpr("fk_project_id"), expr.Value(r.ProjectId)),
		OrderBy: []model.Order{
			{"caption", false},
		},
	}); err != nil {
		return nil, err
	}

	return &optionsV1{
		ConfidentLevels: confidentLevels,
		Goals:           goals,
		Teams:           teams,
	}, nil
}
