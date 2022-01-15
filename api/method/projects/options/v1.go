package options

import (
	"context"
	"sort"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"
)

type reqV1 struct {
	ProjectId string `json:"project_id"`
	WithUsers *bool  `json:"with_users"`
}

type optionsV1 struct {
	Caption           string        `json:"caption"`
	ReachFormat       uint8         `json:"reach_format"`
	ReachDescription  string        `json:"reach_description"`
	EffortDescription string        `json:"effort_description"`
	MoneySymbol       string        `json:"money_symbol"`
	ConfidentLevels   []confidentV1 `json:"confident_levels"`
	Goals             []goalV1      `json:"goals"`
	Teams             []teamV1      `json:"teams"`
	Users             []userV1      `json:"users,omitempty" field:"-"`
}

type confidentV1 struct {
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

type userV1 struct {
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

var errorsV1 struct {
	UnknownProject rpc.ErrorFunc `desc:"Unknown project"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*optionsV1, error) {
	var projects []optionsV1
	if err := m.db.Projects.GetAllToStruct(ctx, &projects, model.GetAllOptions{
		Filter: expr.Eq(m.db.Projects.FieldExpr("str_id"), expr.Value(r.ProjectId)),
		Limit:  1,
	}); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errorsV1.UnknownProject("Unknown project")
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

	if r.WithUsers != nil && *r.WithUsers {
		var projectsUsers []struct {
			User userV1
		}

		if err := m.db.ProjectsUsers.GetAllToStruct(ctx, &projectsUsers, model.GetAllOptions{
			Filter: expr.Eq(m.db.ProjectsUsers.FieldExpr("fk_project_id"), expr.Value(r.ProjectId)),
		}); err != nil {
			return nil, err
		}

		project.Users = make([]userV1, len(projectsUsers))
		for i, pu := range projectsUsers {
			project.Users[i] = pu.User
		}

		sort.Slice(project.Users, func(i, j int) bool {
			return project.Users[i].Fullname < project.Users[j].Fullname
		})
	}

	return &project, nil
}
