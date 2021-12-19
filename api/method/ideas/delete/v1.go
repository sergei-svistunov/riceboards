package delete

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"
)

type reqV1 struct {
	Id uint32 `json:"id"`
}

var errorsV1 struct {
	IdeaNotFound rpc.ErrorFunc `desc:"Idea wasn't found'"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*struct{}, error) {
	var ideas []struct {
		Id uint32
	}
	if err := m.db.Storage.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := m.db.Ideas.GetAllToStruct(ctx, &ideas, model.GetAllOptions{
			Filter:    expr.Eq(m.db.Ideas.FieldExpr("id"), expr.Value(r.Id)),
			Limit:     1,
			ForUpdate: true,
		}); err != nil {
			return err
		}

		if len(ideas) == 0 {
			return errorsV1.IdeaNotFound("Invalid ID")
		}

		if err := m.db.IdeaGoals.Delete(ctx, expr.Eq(m.db.IdeaGoals.FieldExpr("fk_idea_id"), expr.Value(r.Id))); err != nil {
			return err
		}
		if err := m.db.IdeaTeams.Delete(ctx, expr.Eq(m.db.IdeaTeams.FieldExpr("fk_idea_id"), expr.Value(r.Id))); err != nil {
			return err
		}

		return m.db.Ideas.Delete(ctx, expr.Eq(m.db.Ideas.FieldExpr("id"), expr.Value(r.Id)))
	}); err != nil {
		return nil, err
	}

	return &struct{}{}, nil
}
