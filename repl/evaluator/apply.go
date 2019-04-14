package evaluator

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

func apply(f types.Func, env *types.Env, args *types.Cell) (types.Obj, error) {
	f = reflect.Indirect(reflect.ValueOf(f)).Interface().(types.Func)
	if _, ok := f.(types.PrimFuncs); ok {
		return f.Apply(env, args)
	}

	uf, ok := f.(types.UserFuncs)
	if !ok {
		return nil, fmt.Errorf("not supported")
	}
	ne := uf.Env
	if nil == uf.Body {
		return Progn(&ne, nil)
	}
	b := reflect.Indirect(reflect.ValueOf(uf.Body)).Interface().(types.Cell)
	if nil == args {
		return Progn(&ne, &b)
	}
	eargs, err := EvalCell(env, *args)
	if err != nil {
		return nil, err
	}
	p, ok := uf.Params.(*types.Cell)
	if !ok {
		return nil, fmt.Errorf("params of user function was not pointer to cell")
	}
	for eargs != nil {
		found, err := ne.Find(p.Car.String())
		if err != nil || nil == found {
			ne.AddObj(p.Car.String(), eargs.Car)
		} else {
			ne.Set(p.Car.String(), eargs.Car)
		}

		ec, ok := eargs.Cdr.(*types.Cell)
		if !ok && eargs.Cdr != nil {
			return nil, fmt.Errorf("cell was not list while reading eargs: %#v", eargs)
		}
		eargs = ec

		p2, ok2 := p.Cdr.(*types.Cell)
		if !ok2 && p.Cdr != nil {
			return nil, fmt.Errorf("cell was not list while reading params, p: %#v", p)
		}
		p = p2

		if ok != ok2 {
			return nil, fmt.Errorf("number of argument does not match")
		}
	}
	return Progn(&ne, &b)
}
