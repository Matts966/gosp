package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimCons is primitive function in form of (cons a b).
var PrimCons types.Prim = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if l != 2 {
		return nil, fmt.Errorf("invalid number of args passed to PrimCons")
	}

	// Cdr is not nil because the length of list is 2.
	cdr := reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Cell)
	args.Cdr = cdr.Car

	args, err = evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}

	return args, nil
}
