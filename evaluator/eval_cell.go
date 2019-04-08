package evaluator

import (
	"github.com/Matts966/gosp/types"
)

func EvalCell(env *types.Env, c types.Cell) (*types.Cell, error) {
	head := &c
	cp := &c
	for {
		cc, err := Eval(env, cp.Car)
		if err != nil {
			return nil, err
		}
		cp.Car = cc
		switch cd := cp.Cdr.(type) {
		case nil:
			return head, nil
		case *types.Cell:
			cp = cd
		case types.Cell:
			cp = &cd
		default:
			cc, err = Eval(env, cd)
			cp.Cdr = cc
			return head, nil
		}
	}
}
