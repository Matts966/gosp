package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimSetq is in form of (setq symbol val).
var PrimSetq types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, xerrors.Errorf("failed to get the length of args, err: %v", err)
	}
	if 2 != l {
		return nil, xerrors.New("malformed setq")
	}
	// args.Car and args.Cdr is not nil because the length of args is 2
	sym, ok := args.Car.(*types.Symbol)
	if !ok {
		return nil, xerrors.New("cannot setq value to other than pointer to symbol")
	}
	cdr, ok := args.Cdr.(*types.Cell)
	if !ok {
		return nil, xerrors.New("list should be pointer to types.Cell")
	}
	val := cdr.Car
	val, err = evaluator.Eval(env, val)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in setq caused error: %w", err)
	}
	return env.Set(*sym.Name, val)
}
