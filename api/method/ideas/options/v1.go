package options

import (
	"context"
	"sort"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
)

type reqV1 struct {
	ProjectId uint32 `json:"project_id"`
}

type optionsV1 struct {
	ReachFormat       uint8         `json:"reach_format"`
	ReachDescription  string        `json:"reach_description"`
	EffortDescription string        `json:"effort_description"`
	MoneySymbol       string        `json:"money_symbol"`
	ConfidentLevels   []ConfidentV1 `json:"confident_levels"`
	Goals             []goalV1      `json:"goals"`
	Teams             []teamV1      `json:"teams"`
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
	var projects []optionsV1
	if err := m.db.Projects.GetAllToStruct(ctx, &projects, model.GetAllOptions{
		Filter: expr.Eq(m.db.Projects.FieldExpr("id"), expr.Value(r.ProjectId)),
		Limit:  1,
	}); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, nil
	}

	project := projects[0]

	sort.Slice(project.ConfidentLevels, func(i, j int) bool {
		return project.ConfidentLevels[i].Weight < project.ConfidentLevels[j].Weight
	})

	sort.Slice(project.Goals, func(i, j int) bool {
		return project.Goals[i].Caption < project.Goals[j].Caption
	})

	sort.Slice(project.Teams, func(i, j int) bool {
		return project.Teams[i].Caption < project.Teams[j].Caption
	})

	return &project, nil
}
