package types

import "fmt"

// Cell type.
type Cell struct {
	Car *Obj
	Cdr *Obj
}

func (c Cell) Print() {
	fmt.Print("(")
L:
	for {
		if nil == c.Car {
			break L
		}
		(*(c.Car)).Print()
		switch v := (*(c.Cdr)).(type) {
		case nil:
			break L
		case *Cell:
			fmt.Print(" ")
			c = *v
		default:
			fmt.Print(" . ")
			v.Print()
			break L
		}
	}
	fmt.Print(")")
}

func (c Cell) Length() (int, error) {
	len := 1
	for {
		if c.Cdr == nil {
			return len, nil
		}
		switch cdr := (*c.Cdr).(type) {
		case nil:
			return len, nil
		case Cell:
			c = cdr
			len++
		default:
			return 0, fmt.Errorf("cannot handle dotted list")
		}
	}
}

// Cons gets Car, Cdr and return cell pointer.
func Cons(Car Obj, Cdr Obj) Cell {
	return Cell{
		Car: &Car,
		Cdr: &Cdr,
	}
}
