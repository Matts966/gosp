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
			switch bind := (*(cell.Car)).(type) {
			case nil:
				continue
			case *Cell:
				if sym == bind.Car {
					return cell.Car, nil
				}
			default:
				continue
			}

			switch next := (*(cell.Cdr)).(type) {
			case nil:
				break
			case *Cell:
				cell = next
			default:
				break
			}
		}
	}
	return nil, nil
}
