package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

// PrimLambda returns UserFuncss.
var PrimLambda types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if 2 != l {
		return nil, fmt.Errorf("malformed lambda")
	}
	aci := reflect.Indirect(reflect.ValueOf(argList.Car)).Interface()
	_, ok := aci.(types.Cell)
	if !ok {
		if _, ok = aci.(types.False); !ok {
			return nil, fmt.Errorf("the args of lambda is not list")
		}
	}
	_, ok = reflect.Indirect(reflect.ValueOf(argList.Cdr)).Interface().(types.Cell)
	if !ok {
		return nil, fmt.Errorf("the body of lambda is not list")
	}
	return &types.UserFuncs{
		Params: args.Car,
		Body:   args.Cdr,
		Env:    *env,
	}, nil
}
