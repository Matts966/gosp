package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimCar is primitive function in form of (car '(a b c)).
var PrimCar types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	c, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in car caused error: %w", err)
	}
	if c == nil {
		return nil, xerrors.New("nil was passed to function car")
	}
	if cc, ok := c.Car.(*types.Cell); ok {
		return cc.Car, nil
	}
	return nil, xerrors.Errorf("not list value was passed to function car, value: %#v", args)
}
