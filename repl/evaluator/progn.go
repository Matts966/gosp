package evaluator

import (
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

func getLastElement(c *types.Cell) (types.Obj, error) {
	for {
		switch cc := c.Cdr.(type) {
		case nil:
			return c.Car, nil
		case *types.Cell:
			c = cc
		default:
			return cc, nil
		}
	}
}

func Progn(env *types.Env, list *types.Cell) (types.Obj, error) {
	argList, err := EvalCell(env, *list)
	if err != nil {
		return nil, xerrors.Errorf("evaluating list in progn caused error: %w", err)
	}
	return getLastElement(argList)
}
