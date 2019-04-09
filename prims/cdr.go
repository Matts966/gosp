package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimCdr is primitive function in form of (cdr ~).
var PrimCdr types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	c, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, fmt.Errorf("nil was passed to function car")
	}
	cc := reflect.Indirect(reflect.ValueOf(c.Car)).Interface().(types.Obj)
	if cc, ok := cc.(types.Cell); ok {
		return cc.Cdr, nil
	}
	return nil, fmt.Errorf("not list value was passed to function car")
}
