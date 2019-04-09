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

func (c Cell) toString() string {
	s := "("
L:
	for {
		if nil == c.Car {
			break L
		}
		s += c.Car.toString()
		if c.Cdr == nil {
			break
		}
		to, _ := reflect.Indirect(reflect.ValueOf(c.Cdr)).Interface().(Obj)
		switch v := to.(type) {
		case False:
			break L
		case Cell:
			s += " "
			c = v
		default:
			s += " . "
			s += v.toString()
			break L
		}
	}
	return s + ")"
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
