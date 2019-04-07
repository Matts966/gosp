package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

// PrimPlus is primitive function returning the sum of list.
var PrimPlus types.Prim = func(env *types.Env, args types.Obj) (types.Obj, error) {
	args, _ = reflect.Indirect(reflect.ValueOf(args)).Interface().(types.Obj)
	argList, ok := args.(types.Cell)
	if !ok {
		return nil, fmt.Errorf("args is not list")
	}

	val := 0
	for {
		if nil == argList.Car {
			break
		}
		i, ok := argList.Car.(types.Int)
		if !ok {
			return nil, fmt.Errorf("not int values passed to function plus")
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
