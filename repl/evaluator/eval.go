package evaluator

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

func Eval(env *types.Env, obj types.Obj) (types.Obj, error) {
	if obj == nil {
		return nil, fmt.Errorf("nil value is passed to function Eval")
	}
	obj, _ = reflect.Indirect(reflect.ValueOf(obj)).Interface().(types.Obj)
	switch o := obj.(type) {
	case types.Int:
		return obj, nil
	case types.Symbol:
		bind, err := env.Find(*o.Name)
		if err != nil {
			return nil, err
		}
		if nil == bind {
			return nil, fmt.Errorf("undefined symbol: %s", *o.Name)
		}
		switch b := bind.(type) {
		case *types.Cell:
			return b.Cdr, nil
		default:
			return nil, fmt.Errorf("unknown type symbol: %#v, b: %+v", o.Name, b)
		}
	case types.Cell:
		// Function application
		fn, err := Eval(env, o.Car)
		if err != nil {
			return nil, fmt.Errorf("falied to get functon: %+v, err: %+v", fn, err)
		}
		f, ok := fn.(types.Func)
		if !ok {
			return nil, fmt.Errorf("the head of a list must be a function, head: %#v", fn)
		}
		if o.Cdr == nil {
			return f.Apply(env, nil)
		}
		ocd, _ := reflect.Indirect(reflect.ValueOf(o.Cdr)).Interface().(types.Obj)
		c, ok := ocd.(types.Cell)
		if !ok {
			return nil, fmt.Errorf("args is not list, args: %#v", o.Cdr)
		}
		return f.Apply(env, &c)
	default:
		return o, nil
	}
	return nil, fmt.Errorf("unknown type expression: %#v", obj)
}
