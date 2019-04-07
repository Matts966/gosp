package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

// PrimMinus is primitive function in form of (- ~).
var PrimMinus types.Prim = func(env *types.Env, args types.Obj) (types.Obj, error) {
	args, _ = reflect.Indirect(reflect.ValueOf(args)).Interface().(types.Obj)
	argList, ok := args.(types.Cell)
	if !ok {
		return nil, fmt.Errorf("args is not list")
	}

	if nil == argList.Car {
		return types.Int{Value: 0}, nil
	}
	i, ok := argList.Car.(types.Int)
	if !ok {
		return nil, fmt.Errorf("not int values passed to function minus")
	}
	val := i.Value
	// unary minus returns negative value
	if argList.Cdr == nil {
		return types.Int{Value: -val}, nil
	}
	to, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)
	switch v := to.(type) {
	case types.Cell:
		argList = v
	default:
		return nil, fmt.Errorf("malformed minus")
	}
	for {
		if nil == argList.Car {
			break
		}
		i, ok := argList.Car.(types.Int)
		if !ok {
			return nil, fmt.Errorf("not int values passed to function minus")
		}
		val -= i.Value
		if argList.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)
		switch v := to.(type) {
		case types.Cell:
			argList = v
		default:
			return nil, fmt.Errorf("malformed minus")
		}
	}

	return types.Int{Value: val}, nil
}
