package prims

import (
	"fmt"

	"github.com/Matts966/gosp/types"
)

// PrimLambda returns UserFuncs.
var PrimLambda types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, err
	}
	if 2 != l {
		return nil, fmt.Errorf("malformed lambda")
	}
	return makeFunc(env, args)
}
