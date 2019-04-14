package evaluator

import (
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

func Progn(env *types.Env, list *types.Cell) (types.Obj, error) {
	argList, err := EvalCell(env, *list)
	if err != nil {
		return nil, err
	}
	return getLastElement(argList)
}
