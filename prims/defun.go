package prims

import (
	"reflect"

	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

func makeFunc(env *types.Env, list *types.Cell) (*types.UserFuncs, error) {
	aci := reflect.Indirect(reflect.ValueOf(list.Car)).Interface()
	_, ok := aci.(types.Cell)
	if !ok {
		if _, ok = aci.(types.False); !ok {
			return nil, xerrors.New("the args of function is not list")
		}
	}
	_, ok = reflect.Indirect(reflect.ValueOf(list.Cdr)).Interface().(types.Cell)
	if !ok {
		return nil, xerrors.New("the body of function is not list")
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
		return nil, xerrors.Errorf("getting the length of args in car caused error: %w", err)
	}
	if 3 != l {
		return nil, xerrors.New("malformed defun")
	}
	sym, ok := args.Car.(*types.Symbol)
	if !ok {
		return nil, xerrors.Errorf("the first argument was not symbol, but %#v", args.Car)
	}

	aci := reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface()
	acic, ok := aci.(types.Cell)
	if !ok {
		if _, ok = aci.(types.False); !ok {
			return nil, xerrors.New("the args and body of function is not list")
		}
	}
	df, err := makeFunc(env, &acic)
	if err != nil {
		return nil, xerrors.Errorf("making function in defun caused error: %w", err)
	}
	env.AddObj(*sym.Name, df)

	// Assign env again for implementing recurrsion.
	df.Env = *env
	return df, nil
}
