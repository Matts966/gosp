package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimPlus is primitive function returning the sum of list.
var PrimPlus types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	argList := *args
	val := 0
	for {
		if nil == argList.Car {
			break
		}
		i, ok := argList.Car.(types.Int)
		if !ok {
			return nil, fmt.Errorf("not int values passed to function plus, value: %#v", argList.Car)
		}
		val += i.Value
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

	return types.Int{Value: val}, nil
}
