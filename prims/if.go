package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimIf is in form of (if t texpr fexpr).
var PrimIf types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, fmt.Errorf("failed to get the length of args, err: %v", err)
	}
	if l < 2 {
		return nil, fmt.Errorf("malformed if")
	}
	// args.Car and args.Cdr is not nil because the length of args is 2
	cond, err := evaluator.Eval(env, reflect.Indirect(reflect.ValueOf(args.Car)).Interface().(types.Obj))
	if err != nil {
		return nil, fmt.Errorf("evaluating condition failed, cond: %+v", cond)
	}
	// If true, evaluate texpr
	if _, ok := cond.(types.False); !ok {
		return evaluator.Eval(env, reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Cell).Car)
	}

	els := reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Cell).Cdr
	if nil == els {
		return types.False{}, nil
	}
	if elsc, ok := reflect.Indirect(reflect.ValueOf(els)).Interface().(types.Cell); !ok {
		return nil, fmt.Errorf("dotted list was passed to primitive function if")
	} else {
		return evaluator.Progn(env, &elsc)
	}
}
