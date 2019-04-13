package prims

import (
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
)

func getLastElement(c *types.Cell) (types.Obj, error) {
	for {
		switch cc := c.Cdr.(type) {
		case nil:
			return c.Car, nil
		case *types.Cell:
			c = cc
		case types.Cell:
			c = &cc
		default:
			return cc, nil
		}
	}
}

func Progn(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList, err := evaluator.EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	return getLastElement(argList)
}

// PrimProgn is primitive function that evaluates all the item in cell and returns the last value.
var PrimProgn types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	return Progn(env, args)
}
