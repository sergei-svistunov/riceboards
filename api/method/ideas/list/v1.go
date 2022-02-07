package list

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"
)

type reqV1 struct {
	ProjectId string `json:"project_id"`
}

type ideaV1 struct {
	Id               uint32       `json:"id"`
	Caption          string       `json:"caption"`
	Comment          *string      `json:"comment,omitempty"`
	ReadyForDev      bool         `json:"ready_for_dev"`
	Reach            *uint64      `json:"reach,omitempty"`
	ReachComment     *string      `json:"reach_comment,omitempty"`
	Link             *string      `json:"link,omitempty"`
	Confident        *uint32      `json:"confident,omitempty" field:"fk_confident_id"`
	ConfidentComment *string      `json:"confident_comment,omitempty"`
	Goals            []IdeaGoalV1 `json:"goals,omitempty" field:"idea_goals"`
	Teams            []IdeaTeamV1 `json:"teams,omitempty" field:"idea_teams"`
	Done             bool         `json:"done"`
	Owner            UserV1       `json:"owner"`
}

type IdeaGoalV1 struct {
	Id      uint32  `json:"id" field:"fk_goal_id"`
	Value   float64 `json:"value"`
	Comment *string `json:"comment"`
	Owner   UserV1  `json:"owner"`
}

type IdeaTeamV1 struct {
	Id       uint32  `json:"id" field:"fk_team_id"`
	Capacity float64 `json:"capacity"`
	Comment  *string `json:"comment"`
	Owner    UserV1  `json:"owner"`
}

type UserV1 struct {
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

func (m *Method) V1(ctx context.Context, r *reqV1) ([]ideaV1, error) {
	projectId, err := m.db.Projects.GetIdByStrId(ctx, r.ProjectId)
	if err != nil {
		return nil, err
	}
	if projectId == 0 {
		return nil, errorsV1.UnknownProject("Unknown project")
	}

	var ideas []ideaV1
	if err := m.db.Ideas.GetAllToStruct(ctx, &ideas, model.GetAllOptions{
		Filter: expr.Eq(m.db.Ideas.FieldExpr("fk_project_id"), expr.Value(projectId)),
	}); err != nil {
		return nil, err
	}

	return ideas, nil
}
