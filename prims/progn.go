package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
)

// PrimProgn is primitive function that evaluates all the item in cell and returns the last value.
var PrimProgn types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	return evaluator.Progn(env, args)
}
