package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

// PrimSetq is in form of (setq symbol val).
var PrimSetq types.Prim = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, fmt.Errorf("failed to get the length of args, err: %v", err)
	}
	if 2 != l {
		return nil, fmt.Errorf("malformed setq")
	}
	// args.Car and args.Cdr is not nil because the length of args is 2
	sym, ok := reflect.Indirect(reflect.ValueOf(args.Car)).Interface().(types.Symbol)
	if !ok {
		return nil, fmt.Errorf("cannot setq value to other than symbol")
	}
	val := reflect.Indirect(reflect.ValueOf(args.Cdr)).Interface().(types.Cell).Car
	return env.Set(sym, val)
}
