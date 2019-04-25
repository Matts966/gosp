package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimNumeq is primitive function returning the equality of int in form of (eq 3 3 3).
var PrimNumeq types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in numeq(=) caused error: %w", err)
	}
	b := false
	val := 0
	for {
		if nil == args.Car {
			break
		}
		i, ok := args.Car.(types.Int)
		if !ok {
			return nil, xerrors.New("not int values passed to function =")
		}
		if !b {
			val = i.Value
			b = true
		} else {
			if val != i.Value {
				return types.False{}, nil
			}
		}

		if args.Cdr == nil {
			break
		}
		switch v := args.Cdr.(type) {
		case *types.Cell:
			args = v
		default:
			return nil, xerrors.New("malformed =")
		}
	}

	return types.True{}, nil
}
