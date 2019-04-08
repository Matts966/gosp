package prims

import (
	"fmt"

	"github.com/Matts966/gosp/types"
)

// PrimCons is primitive function in form of (cons a b).
var PrimCons types.Prim = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if l != 2 {
		return nil, fmt.Errorf("invalid number of args passed to PrimCons")
	}
	return args, nil
}
