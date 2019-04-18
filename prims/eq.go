package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimEq is primitive function returning the equality in form of (eq 'a 'b 'c).
var PrimEq types.PF = func(env *types.Env, cell *types.Cell) (types.Obj, error) {
	args, err := evaluator.EvalCell(env, *cell)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in eq caused error: %w", err)
	}
	argList := *args
	for {
		if nil == argList.Car {
			break
		}
		if argList.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Obj)
		switch v := to.(type) {
		case types.Cell:
			if argList.Car.String() != v.Car.String() {
				return types.False{}, nil
			}
			switch ac := argList.Car.(type) {
			case types.Func:
				if !ac.Eq(v.Car) {
					return types.False{}, nil
				}
			case types.Symbol, *types.Symbol:
				fmt.Printf("%#v, %#v\n", ac, v.Car)
				if ac != v.Car {
					return types.False{}, nil
				}
			}
			argList = v
		default:
			return nil, xerrors.New("malformed eq")
		}
	}

	return types.True{}, nil
}
