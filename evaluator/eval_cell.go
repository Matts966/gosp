package evaluator

import (
	"reflect"

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
		if nil == cp.Cdr {
			return head, nil
		}
		//TODO(Matts966) ここでコピーが起きていてうまく更新できていないので治す。
		cdr, _ := reflect.Indirect(reflect.ValueOf(cp.Cdr)).Interface().(types.Obj)
		switch cd := cdr.(type) {
		case types.Cell:
			cp = &cd
		default:
			cc, err = Eval(env, cdr)
			cp.Cdr = cc

			return head, nil
		}
	}
}
