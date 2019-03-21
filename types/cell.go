package types

import "fmt"

// Cell type.
type Cell struct {
	Car Obj
	Cdr Obj
}

func (c Cell) Print() {
	fmt.Print("(")
L:
	for {
		if nil == c.Car {
			break L
		}
		c.Car.Print()
		switch c.Cdr.(type) {
		case nil:
			break L
		case *Cell:
			fmt.Print(" ")
			c = *c.Cdr.(*Cell)
		default:
			fmt.Print(" . ")
			c.Cdr.Print()
			break L
		}
	}
	fmt.Print(")")
}

// Cons gets Car, Cdr and return cell pointer.
func Cons(Car Obj, Cdr Obj) *Cell {
	return &Cell{
		Car: Car,
		Cdr: Cdr,
	}
}
