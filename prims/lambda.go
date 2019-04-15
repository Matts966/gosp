package prims

import (
	"github.com/Matts966/gosp/types"
)

// PrimLambda returns UserFuncs.
var PrimLambda types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	return makeFunc(env, args)
}
