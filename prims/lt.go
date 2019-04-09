package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

// PrimLessThan is primitive function in form of (> a b).
var PrimLessThan types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, fmt.Errorf("failed to get the length of args, err: %v", err)
	}
	if 2 != l {
		return nil, fmt.Errorf("malformed lt")
	}
	x, ok := argList.Car.(types.Int)
	if !ok {
		return nil, fmt.Errorf("lt takes only int value, but the first argument was %#v",
			argList.Car)
	}
	cdr, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)
	c, _ := cdr.(types.Cell)
	y, ok := c.Car.(types.Int)
	if !ok {
		return nil, fmt.Errorf("lt takes only int value, but the second argument was %#v",
			argList.Car)
	}
	if x.Value < y.Value {
		return types.True{}, nil
	} else {
		return types.Cell{}, nil
	}
}
