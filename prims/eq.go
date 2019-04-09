package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimEq is primitive function returning the equality in form of (eq 'a 'b 'c).
var PrimEq types.PF = func(env *types.Env, cell *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *cell)
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
			if argList.Car.String() != v.Car.String() {
				return types.False{}, nil
			}
			if f, ok := argList.Car.(types.Func); ok {
				if !f.Eq(v.Car) {
					return types.False{}, nil
				}
			}
			argList = v
		default:
			return nil, fmt.Errorf("malformed eq")
		}
	}

	return types.True{}, nil
}
