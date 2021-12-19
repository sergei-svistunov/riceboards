package edit

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/rpc"

	"riceboards/db"
)

type reqV1 struct {
	Id                uint32 `json:"id"`
	Caption           string `json:"caption"`
	ReachFormat       uint8  `json:"reach_format"`
	ReachDescription  string `json:"reach_description"`
	EffortDescription string `json:"effort_description"`
	MoneySymbol       string `json:"money_symbol"`
}

var errorsV1 struct {
	UnknownProject   rpc.ErrorFunc `desc:"Unknown project"`
	EmptyCaption     rpc.ErrorFunc `desc:"Caption must be filled"`
	EmptyMoneySymbol rpc.ErrorFunc `desc:"Money symbol must be filled"`
	AlreadyExists    rpc.ErrorFunc `desc:"Caption is already busy"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*struct{}, error) {
	if db.FieldIsEmpty(r.Caption) {
		return nil, errorsV1.EmptyCaption("Caption must be filled")
	}
	if db.FieldIsEmpty(r.MoneySymbol) {
		return nil, errorsV1.EmptyMoneySymbol("Money symbol must be filled")
	}

	if err := m.db.Storage.DoInTransaction(ctx, func(ctx context.Context) error {
		data, err := m.db.Projects.GetAll(ctx, []string{"id"}, model.GetAllOptions{
			Filter:    expr.Eq(m.db.Projects.FieldExpr("id"), expr.Value(r.Id)),
			Limit:     1,
			ForUpdate: true,
		})
		if err != nil {
			return err
		}
		if data.Len() == 0 {
			return errorsV1.UnknownProject("Unknown project")
		}

		return m.db.Projects.Edit(ctx, expr.Eq(m.db.Projects.FieldExpr("id"), expr.Value(r.Id)), map[string]interface{}{
			"caption":            r.Caption,
			"reach_format":       r.ReachFormat,
			"reach_description":  r.ReachDescription,
			"effort_description": r.EffortDescription,
			"money_symbol":       r.MoneySymbol,
		})
	}); err != nil {
		return nil, err
	}

	return &struct{}{}, nil
}
