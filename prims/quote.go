package prims

import (
	"fmt"

	"github.com/Matts966/gosp/types"
)

// PrimQuote is in form of (quote ~).
var PrimQuote types.Prim = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if 1 != l {
		return nil, fmt.Errorf("malformed quote")
	}
	return argList.Car, nil
}
