package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimCar is primitive function in form of (car '(a b c)).
var PrimCar types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	c, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, fmt.Errorf("nil was passed to function car")
	}
	cc := reflect.Indirect(reflect.ValueOf(c.Car)).Interface().(types.Obj)
	if cc, ok := cc.(types.Cell); ok {
		return cc.Car, nil
	}
	return nil, fmt.Errorf("not list value was passed to function car, value: %#v", args)
}
