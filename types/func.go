package types

import "fmt"

type Func interface {
	Apply(*Env, *Obj)(*Obj, error)
	Print()
}

type UserFunc struct {
}

type Prim func(env *Env, args *Obj) (*Obj, error)

func isList(obj *Obj) bool {
	_, ok := (*obj).(Cell)
	return obj == nil || ok
}

func (uf UserFunc) Print() {
	fmt.Printf("Func: %+v", uf)
}

func (p Prim) Print() {
	fmt.Printf("Func: %+v", p)
}

func (f UserFunc) Apply(env *Env, args *Obj) (*Obj, error) {
	if !isList(args) {
		return nil, fmt.Errorf("argument must be a list")
	}
	return nil, fmt.Errorf("not implemented")
}

func (p Prim) Apply(env *Env, args *Obj) (*Obj, error) {
	return p(env, args)
}
