package types

import "fmt"

type Func interface {
	Apply(*Env, *Cell) (Obj, error)
	toString() string
}

type UserFunc struct {
}

type Prim func(env *Env, args *Cell) (Obj, error)

func isList(obj Obj) bool {
	_, ok := obj.(Cell)
	return obj == nil || ok
}

func (uf UserFunc) toString() string {
	return "<function>"
}

func (p Prim) toString() string {
	return "<primitive>"
}

func (f UserFunc) Apply(env *Env, args *Cell) (Obj, error) {
	if !isList(args) {
		return nil, fmt.Errorf("argument must be a list")
	}
	return nil, fmt.Errorf("not implemented")
}

func (p Prim) Apply(env *Env, args *Cell) (Obj, error) {
	return p(env, args)
}
