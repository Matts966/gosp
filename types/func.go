package types

import "fmt"

type Func interface {
	Apply(*Env, *Cell) (Obj, error)
	Eq(Obj) bool
}

type UserFunc struct {
	Params Obj
	Body   Obj
	Env    Env
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
		return uf.Params == u.Params && uf.Body == u.Body && uf.Env == u.Env
	case *UserFunc:
		return uf.Params == u.Params && uf.Body == u.Body && uf.Env == u.Env
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
