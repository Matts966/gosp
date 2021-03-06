package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimCons is primitive function in form of (cons a b).
var PrimCons types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, xerrors.Errorf("getting length of args in cons caused error: %w", err)
	}
	if l != 2 {
		return nil, xerrors.New("invalid number of args passed to PrimCons")
	}

	args, err = evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in cons caused error: %w", err)
	}

	// Cdr is not nil because the length of list is 2.
	cdr, ok := args.Cdr.(*types.Cell)
	if !ok {
		return nil, xerrors.Errorf(
			"unknown value other than pointer to cell is passed to cons, value: %+v", cdr)
	}
	args.Cdr = cdr.Car

	return args, nil
}
