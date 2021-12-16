package edit

import (
	"context"
	"errors"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"

	"riceboards/authctx"
	"riceboards/db"
)

type reqV1 struct {
	Id               uint32              `json:"id"`
	Caption          *string             `json:"caption"`
	Link             *string             `json:"link"`
	Reach            *uint64             `json:"reach"`
	ReachComment     *string             `json:"reach_comment"`
	Confident        *uint32             `json:"confident"`
	ConfidentComment *string             `json:"confident_comment"`
	Goals            *map[uint32]float64 `json:"goals"`
	GoalsComments    *map[uint32]string  `json:"goals_comments"`
	Teams            *map[uint32]float64 `json:"teams"`
	TeamsComments    *map[uint32]string  `json:"teams_comments"`
}

var errorsV1 struct {
	Unauthorized  rpc.ErrorFunc `desc:"Unauthorized"`
	EmptyCaption  rpc.ErrorFunc `desc:"Caption must be filled"`
	AlreadyExists rpc.ErrorFunc `desc:"Caption is already busy"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*struct{}, error) {
	curUserId, err := authctx.GetCurUserId(ctx)
	if err != nil {
		if errors.Is(err, authctx.ErrUnauthorized) {
			return nil, errorsV1.Unauthorized("Unauthorized")
		}
		return nil, err
	}

	if r.Caption != nil && db.FieldIsEmpty(*r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}

	if err := m.db.Storage.DoInTransaction(ctx, func(ctx context.Context) error {
		ideaFields := map[string]interface{}{}

		if r.Caption != nil {
			ideaFields["caption"] = *r.Caption
		}
		if r.Link != nil {
			ideaFields["link"] = *r.Link
		}
		if r.Reach != nil {
			ideaFields["reach"] = *r.Reach
		}
		if r.ReachComment != nil {
			ideaFields["reach_comment"] = *r.ReachComment
		}
		if r.Confident != nil {
			ideaFields["fk_confident_id"] = *r.Confident
		}
		if r.ConfidentComment != nil {
			ideaFields["confident_comment"] = *r.ConfidentComment
		}

		if len(ideaFields) > 0 {
			if err := m.db.Ideas.Edit(ctx, expr.Eq(m.db.Ideas.FieldExpr("id"), expr.Value(r.Id)), ideaFields); err != nil {
				return err
			}
		}

		if r.Goals != nil && len(*r.Goals) > 0 {
			type goal struct {
				FkOwnerId uint32
				FkIdeaId  uint32
				FkGoalId  uint32
				Value     float64
				Comment   *string
			}
			var goals []goal
			for id, value := range *r.Goals {
				goals = append(goals, goal{
					FkOwnerId: curUserId,
					FkIdeaId:  r.Id,
					FkGoalId:  id,
					Value:     value,
				})
				if r.GoalsComments != nil {
					if c, exists := (*r.GoalsComments)[id]; exists {
						goals[len(goals)-1].Comment = &c
					}
				}
			}

			if _, err := m.db.IdeaGoals.AddFromStructs(ctx, goals, model.AddOptions{Replace: true}); err != nil {
				return err
			}
		}

		if r.Teams != nil && len(*r.Teams) > 0 {
			type team struct {
				FkOwnerId uint32
				FkIdeaId  uint32
				FkTeamId  uint32
				Capacity  float64
				Comment   *string
			}
			var teams []team
			for id, value := range *r.Teams {
				teams = append(teams, team{
					FkOwnerId: curUserId,
					FkIdeaId:  r.Id,
					FkTeamId:  id,
					Capacity:  value,
				})
				if r.TeamsComments != nil {
					if c, exists := (*r.TeamsComments)[id]; exists {
						teams[len(teams)-1].Comment = &c
					}
				}
			}

			if _, err := m.db.IdeaTeams.AddFromStructs(ctx, teams, model.AddOptions{Replace: true}); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &struct{}{}, nil
}
