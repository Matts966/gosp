package types

import "fmt"

// Env type
type Env struct {
	vars *Cell
	up   *Env
}

func (e Env) Print() {
	fmt.Print("Env: ")
	for {
		if nil == e.vars {
			return
		}
		(*(e.vars)).Print()
		if nil == e.up {
			return
		}
		e = *e.up
	}
}

// Find finds symbols from environment.
func (e *Env) Find(sym *Obj) (*Obj, error) {
	for p := e; p != nil; p = p.up {
		cell := p.vars
		for {
			if nil == cell {
				return nil, nil
			}
			if nil == cell.Car {
				return nil, nil
			}
			switch bind := (*(cell.Car)).(type) {
			case nil:
				break
			case Cell:
				s, ok := (*sym).(Symbol)
				if !ok {
					return nil, fmt.Errorf("passed sym is not Symbol")
				}
				bc, ok := (*(bind.Car)).(Symbol)
				if !ok {
					return nil, fmt.Errorf("symbol in env is not Symbol")
				}
				if s.Name == bc.Name {
					return cell.Car, nil
				}
			default:
				return nil, fmt.Errorf("unknown bind type, bind: %#v", bind)
			}

			switch next := (*(cell.Cdr)).(type) {
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

func (e *Env) AddObj(name string, obj Obj) {
	new := Cons(Cons(Symbol{
		Name: name,
	}, obj), e.vars)
	e.vars = &new
}
