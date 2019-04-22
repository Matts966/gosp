package prims

import (
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimIf is in form of (if t texpr fexpr).
var PrimIf types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, xerrors.Errorf("failed to get the length of args, err: %v", err)
	}
	if l < 2 {
		return nil, xerrors.New("malformed if")
	}
	// args.Car and args.Cdr is not nil because the length of args is 2
	cond, err := evaluator.Eval(env, reflect.Indirect(reflect.ValueOf(args.Car)).Interface().(types.Obj))
	if err != nil {
		return nil, xerrors.Errorf("evaluating condition failed, cond: %+v", cond)
	}
	// If true, evaluate texpr
	if _, ok := cond.(types.False); !ok {
		cd, ok := args.Cdr.(*types.Cell)
		if !ok {
			return nil, xerrors.Errorf("evaluating expr failed in function if, expr: %+v", args.Cdr)
		}
		return evaluator.Eval(env, cd.Car)
	}

	cd, ok := args.Cdr.(*types.Cell)
	if !ok {
		return nil, xerrors.New("dotted list was passed to primitive function if")
	}
	elsc, ok := cd.Cdr.(*types.Cell)
	if nil == cd.Cdr {
		return types.False{}, nil
	}
	if !ok {
		return nil, xerrors.New("dotted list was passed to primitive function if")
	}
	return evaluator.Progn(env, elsc)
}
