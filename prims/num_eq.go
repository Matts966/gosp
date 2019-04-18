package prims

import (
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimNumeq is primitive function returning the equality of int in form of (eq 3 3 3).
var PrimNumeq types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in numeq(=) caused error: %w", err)
	}
	argList := *args
	b := false
	val := 0
	for {
		if nil == argList.Car {
			break
		}
		i, ok := argList.Car.(types.Int)
		if !ok {
			return nil, xerrors.New("not int values passed to function =")
		}
		if !b {
			val = i.Value
			b = true
		} else {
			if val != i.Value {
				return types.False{}, nil
			}
		}

		if argList.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)
		switch v := to.(type) {
		case types.Cell:
			argList = v
		default:
			return nil, xerrors.New("malformed =")
		}
	}

	return types.True{}, nil
}
