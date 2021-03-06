package evaluator

import (
	"reflect"

	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

func apply(f types.Func, env *types.Env, args *types.Cell) (types.Obj, error) {
	f = reflect.Indirect(reflect.ValueOf(f)).Interface().(types.Func)
	if _, ok := f.(types.PrimFuncs); ok {
		return f.Apply(env, args)
	}

	uf, ok := f.(types.UserFuncs)
	if !ok {
		return nil, xerrors.New("not supported")
	}
	ne := uf.Env
	if nil == uf.Body {
		return Progn(&ne, nil)
	}
	b, ok := uf.Body.(*types.Cell)
	if !ok {
		return nil, xerrors.New("list should be pointer to types.Cell")
	}
	if nil == args {
		return Progn(&ne, b)
	}
	eargs, err := EvalCell(env, *args)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in apply caused error: %w", err)
	}
	p, ok := uf.Params.(*types.Cell)
	if !ok {
		return nil, xerrors.New("params of user function was not pointer to cell")
	}

	// Map for new env
	var m *types.Cell
	for eargs != nil {
		s, ok := p.Car.(*types.Symbol)
		if !ok {
			return nil, xerrors.New("not symbol parameter")
		}
		m = types.Cons(types.Cons(s, eargs.Car), m)

		ec, ok := eargs.Cdr.(*types.Cell)
		if !ok && eargs.Cdr != nil {
			return nil, xerrors.Errorf("cell was not list while reading eargs: %#v", eargs)
		}
		eargs = ec

		p2, ok2 := p.Cdr.(*types.Cell)
		if !ok2 && p.Cdr != nil {
			// Implement variadic function.
			s, ok := p.Cdr.(*types.Symbol)
			if !ok {
				return nil, xerrors.Errorf("cell was not list or last symbol while reading params, p: %#v", p)
			}
			m = types.Cons(types.Cons(s, eargs), m)
			break
		}
		p = p2

		if ok != ok2 {
			return nil, xerrors.New("number of argument does not match")
		}
	}
	ne.AddScope(m)

	return Progn(&ne, b)
}
