package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimNumeq is primitive function returning the equality of int.
var PrimNumeq types.Prim = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	argList := *args
	b := false
	val := 0
	for {
		if nil == argList.Car {
			break
		}
		i, ok := argList.Car.(types.Int)
		if !ok {
			return nil, fmt.Errorf("not int values passed to function plus")
		}
		if !b {
			val = i.Value
			b = true
		} else {
			if val != i.Value {
				return types.False{}, nil
			}
		}

		if argList.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)
		switch v := to.(type) {
		case types.Cell:
			argList = v
		default:
			return nil, fmt.Errorf("malformed plus")
		}
	}

	return types.True{}, nil
}
