package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimCdr is primitive function in form of (cdr ~).
var PrimCdr types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	c, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in cdr caused error: %w", err)
	}
	if c == nil {
		return nil, xerrors.New("nil was passed to function cdr")
	}
	if cc, ok := c.Car.(*types.Cell); ok {
		return cc.Cdr, nil
	}
	return nil, xerrors.New("not list value was passed to function cdr")
}
