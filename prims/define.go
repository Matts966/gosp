package prims

import (
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimDefine is primitive function in form of (define a 'a).
var PrimDefine types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, xerrors.Errorf("getting length of args in define caused error: %w", err)
	}
	if l != 2 {
		return nil, xerrors.New("invalid number of args passed to PrimDefine")
	}

	s, ok := reflect.Indirect(reflect.ValueOf(args.Car)).Interface().(types.Symbol)
	if !ok {
		return nil, xerrors.New("the car of the cell passed to define was not a symbol")
	}

	// ac is cell because the lenght of args is 2.
	acc, _ := args.Cdr.(*types.Cell)
	val, err := evaluator.Eval(env, acc.Car)
	if err != nil {
		return nil, xerrors.Errorf("evaluating object in define caused error: %w", err)
	}
	env.AddObj(*s.Name, val)
	return val, nil
}
