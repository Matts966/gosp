package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimDefine is primitive function in form of (define a 'a).
var PrimDefine types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, err
	}
	if l != 2 {
		return nil, fmt.Errorf("invalid number of args passed to PrimCons")
	}

	s, ok := reflect.Indirect(reflect.ValueOf(args.Car)).Interface().(types.Symbol)
	if !ok {
		return nil, fmt.Errorf("the car of the cell passed to define was not a symbol")
	}

	ac := reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Obj)
	// ac is cell because the lenght of args is 2
	acc, _ := ac.(types.Cell)
	val, err := evaluator.Eval(env, acc.Car)
	if err != nil {
		return nil, err
	}
	env.AddObj(*s.Name, val)
	return val, nil
}
