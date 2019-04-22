package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimMinus is primitive function in form of (- ~).
var PrimMinus types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in minus caused error: %w", err)
	}
	if nil == args.Car {
		return types.Int{Value: 0}, nil
	}
	i, ok := args.Car.(types.Int)
	if !ok {
		return nil, xerrors.New("not int values passed to function minus")
	}
	val := i.Value
	// unary minus returns negative value
	if args.Cdr == nil {
		return types.Int{Value: -val}, nil
	}
	switch v := args.Cdr.(type) {
	case *types.Cell:
		args = v
	default:
		return nil, xerrors.New("malformed minus")
	}
	for {
		if nil == args.Car {
			break
		}
		i, ok := args.Car.(types.Int)
		if !ok {
			return nil, xerrors.New("not int values passed to function minus")
		}
		val -= i.Value
		if args.Cdr == nil {
			break
		}
		switch v := args.Cdr.(type) {
		case *types.Cell:
			args = v
		default:
			return nil, xerrors.New("malformed minus")
		}
	}

	return types.Int{Value: val}, nil
}
