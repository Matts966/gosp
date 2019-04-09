package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimSetCar is in form of (setcar ~).
var PrimSetCar types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, fmt.Errorf("failed to get the length of args, err: %v", err)
	}
	if 2 != l {
		return nil, fmt.Errorf("malformed setcar")
	}
	args, err = evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	// args.Cdr is not nil because the length of args is 2
	acdr := reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Cell).Car

	if _, ok := args.Car.(*types.Cell); !ok {
		return nil, fmt.Errorf("not pointer value returned from env")
	}
	args.Car.(*types.Cell).Car = acdr
	return args.Car, nil
}
