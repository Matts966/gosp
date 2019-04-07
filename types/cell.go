package types

import (
	"fmt"
	"reflect"
)

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
		(c.Car).Print()
		if c.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(c.Cdr)).Interface().(Obj)
		switch v := to.(type) {
		case Cell:
			fmt.Print(" ")
			c = v
		default:
			fmt.Print(" . ")
			v.Print()
			break L
		}
	}
	fmt.Print(")")
}

func (c *Cell) Length() (int, error) {
	cp := c
	len := 1
	for {
		if nil == cp.Cdr {
			return len, nil
		}
		cdr, _ := reflect.Indirect(reflect.ValueOf(cp.Cdr)).Interface().(Obj)
		switch c := cdr.(type) {
		case Cell:
			cp = &c
			len++
		default:
			return 0, fmt.Errorf("cannot handle dotted list")
		}
	}
}

// Cons gets Car, Cdr and return cell pointer.
func Cons(Car Obj, Cdr Obj) Cell {
	return Cell{
		Car: Car,
		Cdr: Cdr,
	}
}
