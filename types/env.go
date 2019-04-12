package types

import (
	"fmt"
)

// Env type
type Env struct {
	vars *Cell
	up   *Env
}

func (e Env) String() string {
	s := "Env: "
	for {
		if nil == e.vars {
			return s
		}
		s += (*(e.vars)).String()
		if nil == e.up {
			return s
		}
		e = *e.up
	}
}

// Find finds symbols from environment and returns its pointer.
func (e *Env) Find(name string) (Obj, error) {
	for p := e; p != nil; p = p.up {
		cell := p.vars
		for {
			if nil == cell {
				return nil, nil
			}
			if nil == cell.Car {
				return nil, nil
			}
			switch bind := cell.Car.(type) {
			case nil:
				break
			case *Cell:
				bc, ok := bind.Car.(*Symbol)
				if !ok {
					return nil, fmt.Errorf("symbol in env is not pointer to Symbol")
				}
				if name == *bc.Name {
					return cell.Car, nil
				}
			default:
				return nil, fmt.Errorf("unknown bind type, bind: %#v", bind)
			}

			switch next := cell.Cdr.(type) {
			case nil:
				break
			case *Cell:
				cell = next
			default:
				return nil, fmt.Errorf("unknown next type, next: %#v", next)
			}
		}
	}
	return nil, nil
}

func (e *Env) Set(name string, obj Obj) (Obj, error) {
	for p := e; p != nil; p = p.up {
		cell := p.vars
		for {
			if nil == cell {
				return nil, fmt.Errorf("symbol %s not found", name)
			}
			if nil == cell.Car {
				return nil, fmt.Errorf("symbol %s not found", name)
			}
			switch bind := cell.Car.(type) {
			case nil:
				break
			case *Cell:
				bc, ok := bind.Car.(*Symbol)
				if !ok {
					return nil, fmt.Errorf("symbol in env should be Symbol")
				}
				if name == *bc.Name {
					bind.Cdr = obj
					return obj, nil
				}
			default:
				return nil, fmt.Errorf("unknown bind type, bind: %#v", bind)
			}

			switch next := cell.Cdr.(type) {
			case nil:
				break
			case *Cell:
				cell = next
			default:
				return nil, fmt.Errorf("unknown next type, next: %#v", next)
			}
		}
	}
	return nil, fmt.Errorf("symbol %s not found", name)
}

func (e *Env) AddObj(name string, obj Obj) {
	new := Cons(Cons(&Symbol{
		Name: &name,
	}, obj), e.vars)
	e.vars = new
}
