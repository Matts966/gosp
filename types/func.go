package types

import "fmt"

type Func interface {
	Apply(*Env, *Cell) (Obj, error)
	Eq(Obj) bool
}

type UserFunc struct {
	params Obj
	body   Obj
	env    Obj
}

func isList(obj Obj) bool {
	_, ok := obj.(Cell)
	return obj == nil || ok
}

func (uf UserFunc) String() string {
	return "<function>"
}

func (uf UserFunc) Eq(o Obj) bool {
	switch u := o.(type) {
	case UserFunc:
		return uf.params == u.params && uf.body == u.body && uf.env == u.env
	case *UserFunc:
		return uf.params == u.params && uf.body == u.body && uf.env == u.env
	default:
		return false
	}
}

func (f UserFunc) Apply(env *Env, args *Cell) (Obj, error) {
	if !isList(args) {
		return nil, fmt.Errorf("argument must be a list")
	}
	return nil, fmt.Errorf("not implemented")
}

type PF func(env *Env, args *Cell) (Obj, error)

type PrimFuncs struct {
	F *PF
}

func (p PrimFuncs) String() string {
	return "<primitive>"
}

func (p PrimFuncs) Eq(o Obj) bool {
	if pr, ok := o.(PrimFuncs); ok {
		return p.F == pr.F
	}
	return false
}

func (p PrimFuncs) Apply(env *Env, args *Cell) (Obj, error) {
	return (*p.F)(env, args)
}
