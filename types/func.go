package types

import "fmt"

type Func interface {
	Apply(*Env, *Cell) (Obj, error)
	Eq(Obj) bool
}

type UserFuncs struct {
	Params Obj
	Body   Obj
	Env    Env
}

func (uf UserFuncs) String() string {
	return "<function>"
}

func (uf UserFuncs) Eq(o Obj) bool {
	switch u := o.(type) {
	case UserFuncs:
		return uf.Params == u.Params && uf.Body == u.Body && uf.Env == u.Env
	case *UserFuncs:
		return uf.Params == u.Params && uf.Body == u.Body && uf.Env == u.Env
	default:
		return false
	}
}

func (f UserFuncs) Apply(env *Env, args *Cell) (Obj, error) {
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
