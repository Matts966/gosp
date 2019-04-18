package prims

import (
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimMinus is primitive function in form of (- ~).
var PrimMinus types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in minus caused error: %w", err)
	}
	argList := *args
	if nil == argList.Car {
		return types.Int{Value: 0}, nil
	}
	i, ok := argList.Car.(types.Int)
	if !ok {
		return nil, xerrors.New("not int values passed to function minus")
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
		return nil, xerrors.New("malformed minus")
	}
	for {
		if nil == argList.Car {
			break
		}
		i, ok := argList.Car.(types.Int)
		if !ok {
			return nil, xerrors.New("not int values passed to function minus")
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
			return nil, xerrors.New("malformed minus")
		}
	}

	return types.Int{Value: val}, nil
}
