package prims

import (
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimSetCar is in form of (setcar ~).
var PrimSetCar types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, xerrors.Errorf("failed to get the length of args, err: %v", err)
	}
	if 2 != l {
		return nil, xerrors.New("malformed setcar")
	}
	args, err = evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in setcar caused error: %w", err)
	}
	// args.Cdr is not nil because the length of args is 2
	acdr := reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Cell).Car

	if _, ok := args.Car.(*types.Cell); !ok {
		return nil, xerrors.New("not pointer value returned from env")
	}
	args.Car.(*types.Cell).Car = acdr
	return args.Car, nil
}
