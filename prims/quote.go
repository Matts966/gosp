package prims

import (
	"fmt"

	"github.com/Matts966/gosp/types"
)

// PrimQuote is in form of (quote ~).
var PrimQuote types.Prim = func(env *types.Env, args types.Obj) (types.Obj, error) {
	argList, ok := args.(types.Cell)
	if !ok {
		return nil, fmt.Errorf("args is not list")
	}
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if 1 != l {
		return nil, fmt.Errorf("malformed quote")
	}
	return argList.Car, nil
}
