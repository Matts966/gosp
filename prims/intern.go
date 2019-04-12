package prims

import (
	"fmt"
	"reflect"

	"github.com/Matts966/gosp/types"
)

// Intern inputs symbol table and symbol name and returns pointer to symbol
func Intern(symbolTable *types.Cell, name string) (types.Obj, error) {
	s := types.Cons(&types.Symbol{Name: &name}, nil)
	if nil == symbolTable {
		symbolTable = s
		return s.Car, nil
	}
	for {
		if *symbolTable.Car.(*types.Symbol).Name == name {
			return symbolTable.Car, nil
		}
		if nil == symbolTable.Cdr {
			symbolTable.Cdr = s
			return s.Car, nil
		}
		symbolTable = symbolTable.Cdr.(*types.Cell)
	}
}

// PrimIntern finds symbol and returns it or creates new symbol if not found in form of (intern sym).
var PrimIntern types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	l, err := args.Length()
	if err != nil {
		return nil, fmt.Errorf("failed to get the length of args, err: %v", err)
	}
	if 1 != l {
		return nil, fmt.Errorf("malformed intern")
	}
	st := "symbol_table"
	symbolTable, err := env.Find(st)
	if err != nil {
		return nil, err
	}
	return Intern(symbolTable.(*types.Cell), *reflect.Indirect(reflect.ValueOf(args.Car)).Interface().(types.Symbol).Name)
}
