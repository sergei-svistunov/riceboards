package list

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
)

type reqV1 struct {
	ProjectId uint32 `json:"project_id"`
}

type ideaV1 struct {
	Id               uint32       `json:"id"`
	Caption          string       `json:"caption"`
	Reach            *uint64      `json:"reach,omitempty"`
	ReachComment     *string      `json:"reach_comment,omitempty"`
	Link             *string      `json:"link,omitempty"`
	Confident        *uint32      `json:"confident,omitempty" field:"fk_confident_id"`
	ConfidentComment *string      `json:"confident_comment,omitempty"`
	Goals            []IdeaGoalV1 `json:"goals,omitempty" field:"idea_goals"`
	Teams            []IdeaTeamV1 `json:"teams,omitempty" field:"idea_teams"`
}

type IdeaGoalV1 struct {
	Id      uint32  `json:"id" field:"fk_goal_id"`
	Value   float64 `json:"value"`
	Comment *string `json:"comment"`
}

type IdeaTeamV1 struct {
	Id       uint32  `json:"id" field:"fk_team_id"`
	Capacity float64 `json:"capacity"`
	Comment  *string `json:"comment"`
}

func (m *Method) V1(ctx context.Context, r *reqV1) ([]ideaV1, error) {
	var ideas []ideaV1
	if err := m.db.Ideas.GetAllToStruct(ctx, &ideas, model.GetAllOptions{
		Filter: expr.Eq(m.db.Ideas.FieldExpr("fk_project_id"), expr.Value(r.ProjectId)),
	}); err != nil {
		return nil, err
	}

	return ideas, nil
}
