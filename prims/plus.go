package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimPlus is primitive function returning the sum of list.
var PrimPlus types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating function in plus(+) caused error: %w", err)
	}
	val := 0
	for {
		if nil == args.Car {
			break
		}
		i, ok := args.Car.(types.Int)
		if !ok {
			return nil, xerrors.Errorf("not int values passed to function plus, value: %#v", args.Car)
		}
		val += i.Value
		if args.Cdr == nil {
			break
		}
		switch v := args.Cdr.(type) {
		case *types.Cell:
			args = v
		default:
			return nil, xerrors.New("malformed plus")
		}
	}

	return types.Int{Value: val}, nil
}
