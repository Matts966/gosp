package types

import (
	"reflect"

	"golang.org/x/xerrors"
)

// Cell type.
type Cell struct {
	Car Obj
	Cdr Obj
}

func (c Cell) String() string {
	s := "("
L:
	for {
		if nil == c.Car {
			break L
		}
		s += c.Car.String()
		cv := reflect.Indirect(reflect.ValueOf(c.Cdr))
		if !cv.IsValid() {
			break
		}
		to, _ := cv.Interface().(Obj)
		switch v := to.(type) {
		case False:
			break L
		case Cell:
			s += " "
			c = v
		default:
			s += " . "
			s += v.String()
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
			return 0, xerrors.New("cannot handle dotted list")
		}
	}
}

// Cons gets Car, Cdr and return cell pointer.
func Cons(Car Obj, Cdr Obj) *Cell {
	return &Cell{
		Car: Car,
		Cdr: Cdr,
	}
}
