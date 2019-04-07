package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/Matts966/gosp/prims"
	"github.com/Matts966/gosp/scanner"
	"github.com/Matts966/gosp/types"
)

var (
	scn scanner.Scanner
	env types.Env = types.Env{}
)

func init() {
	scn.Init(os.Stdin)
	prims.AddPrims(&env)
}

func readQuote() (types.Obj, error) {
	obj, err := readExpr()
	if err != nil {
		return nil, fmt.Errorf("failed to read quote")
	}
	return types.Cons(types.Symbol{
		Name: "quote",
	}, types.Cons(obj, nil)), nil
}

func isDigit(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err == nil {
		return true
	}
	return false
}

func isSymbolRune(r rune) bool {
	if 'A' <= r && r <= 'z' {
		return true
	}
	for _, s := range types.Symbols {
		if s == r {
			return true
		}
	}
	return false
}

func readNumber(v int) int {
	for isDigit(scn.Peek()) {
		v = v*10 + int(scn.Next()-'0')
	}
	return v
}

func readSymbol(c rune) *types.Symbol {
	n := scn.Peek()
	name := string(c)
	for isDigit(n) || isSymbolRune(n) {
		name += string(scn.Next())
		n = scn.Peek()
	}
	return &types.Symbol{
		Name: name,
	}
}

func readList() (types.Obj, error) {
	obj, err := readExpr()
	if err != nil {
		return obj, err
	}
	switch obj.(type) {
	case nil:
		return nil, fmt.Errorf("unclosed parenthesis")
	case types.Dot:
		return nil, fmt.Errorf("stray dot")
	case types.RParen:
		return nil, nil
	}
	head := types.Cons(obj, nil)
	var tail types.Obj
	tail = &head
	for {
		obj, err := readExpr()
		if err != nil {
			return nil, err
		}
		switch obj.(type) {
		case nil:
			return nil, fmt.Errorf("unclosed parenthesis")
		case types.Dot:
			t, _ := tail.(*types.Cell)
			t.Cdr, err = readExpr()
			if err != nil {
				return nil, err
			}
			nx, err := readExpr()
			if err != nil {
				return nil, err
			}
			if _, ok := nx.(types.RParen); !ok {
				return nil, fmt.Errorf("closed parenthesis expected after dot")
			}
			return head, nil
		case types.RParen:
			return head, nil
		default:
			t, _ := tail.(*types.Cell)
			nc := types.Cons(obj, nil)
			t.Cdr = &nc
			tail = t.Cdr
		}
	}
}

func readExpr() (types.Obj, error) {
	c := scn.Next()
	for c != scanner.EOF {
		switch c {
		case ' ', '\n', '\r', '\t':
		case '(':
			return readList()
		case ')':
			return types.RParen{}, nil
		case '.':
			return types.Dot{}, nil
		case '\'':
			return readQuote()
		case '-':
			if isDigit(scn.Peek()) {
				return types.Int{
					Value: -readNumber(int(scn.Next() - '0')),
				}, nil
			}
			fallthrough
		default:
			if isDigit(c) {
				return types.Int{
					Value: readNumber(int(c - '0')),
				}, nil
			}
			return readSymbol(c), nil
		}
		c = scn.Next()
	}
	return nil, io.EOF
}

func eval(obj types.Obj) (types.Obj, error) {
	obj, _ = reflect.Indirect(reflect.ValueOf(obj)).Interface().(types.Obj)
	switch o := obj.(type) {
	case types.Int:
		return obj, nil
	case types.Symbol:
		bind, err := env.Find(o)
		if err != nil {
			return nil, err
		}
		if nil == bind {
			return nil, fmt.Errorf("undefined symbol: %s", o.Name)
		}
		switch b := bind.(type) {
		case types.Cell:
			return b.Cdr, nil
		default:
			return nil, fmt.Errorf("unknown type symbol: %s, b: %+v", o.Name, b)
		}
	case types.Cell:
		// Function application
		fn, err := eval(o.Car)
		if err != nil {
			return nil, fmt.Errorf("falied to get functon: %+v, err: %+v", fn, err)
		}
		f, ok := fn.(types.Func)
		if !ok {
			return nil, fmt.Errorf("the head of a list must be a function, head: %#v", fn)
		}
		if o.Cdr == nil {
			return f.Apply(&env, nil)
		}
		ocd, _ := reflect.Indirect(reflect.ValueOf(o.Cdr)).Interface().(types.Obj)
		c, ok := ocd.(types.Cell)
		if !ok {
			return nil, fmt.Errorf("args is not list, args: %#v", o.Cdr)
		}
		return f.Apply(&env, &c)
	}
	return nil, fmt.Errorf("unknown type expression: %#v", obj)
}

func main() {
	for {
		obj, err := readExpr()
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch obj.(type) {
		case nil:
			continue
		case types.Dot:
			fmt.Println(fmt.Errorf("stray dot"))
			os.Exit(1)
		case types.RParen:
			fmt.Println(fmt.Errorf("unmatched right parenthesis"))
			os.Exit(1)
		}

		o, err := eval(obj)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		o.Print()
		fmt.Println()
	}
}
