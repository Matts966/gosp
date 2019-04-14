package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

func makeFunc(env *types.Env, list *types.Cell) (*types.UserFuncs, error) {
	aci := reflect.Indirect(reflect.ValueOf(list.Car)).Interface()
	_, ok := aci.(types.Cell)
	if !ok {
		if _, ok = aci.(types.False); !ok {
			return nil, fmt.Errorf("the args of function is not list")
		}
	}
	_, ok = reflect.Indirect(reflect.ValueOf(list.Cdr)).Interface().(types.Cell)
	if !ok {
		return nil, fmt.Errorf("the body of function is not list")
	}
	return &types.UserFuncs{
		Params: list.Car,
		Body:   list.Cdr,
		Env:    *env,
	}, nil
}

// PrimDefun returns UserFuncs.
var PrimDefun types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if 3 != l {
		return nil, fmt.Errorf("malformed defun")
	}
	sym, ok := args.Car.(*types.Symbol)
	if !ok {
		return nil, fmt.Errorf("the first argument was not symbol, but %#v", args.Car)
	}

	aci := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface()
	acic, ok := aci.(types.Cell)
	if !ok {
		if _, ok = aci.(types.False); !ok {
			return nil, fmt.Errorf("the args and body of function is not list")
		}
	}
	df, err := makeFunc(env, &acic)
	if err != nil {
		return nil, err
	}
	env.AddObj(*sym.Name, df)

	// Assign env again for implementing recurrsion.
	df.Env = *env
	return df, nil
}
