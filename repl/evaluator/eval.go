package evaluator

import (
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

func Eval(env *types.Env, obj types.Obj) (types.Obj, error) {
	switch o := obj.(type) {
	case types.Int:
		return obj, nil
	case *types.Symbol:
		bind, err := env.Find(*o.Name)
		if err != nil {
			return nil, xerrors.Errorf("finding object from env in Eval caused error: %w", err)
		}
		if nil == bind {
			return nil, xerrors.Errorf("undefined symbol: %s", *o.Name)
		}
		switch b := bind.(type) {
		case *types.Cell:
			return b.Cdr, nil
		default:
			return nil, xerrors.Errorf("unknown type symbol: %#v, b: %+v", o.Name, b)
		}
	case *types.Cell:
		// Function application
		fn, err := Eval(env, o.Car)
		if err != nil {
			return nil, xerrors.Errorf("falied to get functon: %+v, err: %+v", fn, err)
		}
		f, ok := fn.(types.Func)
		if !ok {
			return nil, xerrors.Errorf("the head of a list must be a function, head: %#v", fn)
		}
		if o.Cdr == nil {
			return apply(f, env, nil)
		}
		c, ok := o.Cdr.(*types.Cell)
		if !ok {
			return nil, xerrors.Errorf("args is not list, args: %#v", o.Cdr)
		}
		return apply(f, env, c)
	default:
		return o, nil
	}
	return nil, xerrors.Errorf("unknown type expression: %#v", obj)
}
