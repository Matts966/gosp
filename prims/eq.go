package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimEq is primitive function returning the equality in form of (eq 'a 'b 'c).
var PrimEq types.Prim = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	argList := *args
	for {
		if nil == argList.Car {
			break
		}

		if argList.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)

		switch v := to.(type) {
		case types.Cell:
			x := reflect.Indirect(reflect.ValueOf(argList.Car)).Interface()
			y := reflect.Indirect(reflect.ValueOf(v.Car)).Interface()
			if x != y {
				return types.False{}, nil
			}
			argList = v
		default:
			return nil, fmt.Errorf("malformed eq")
		}
	}

	return types.True{}, nil
}
