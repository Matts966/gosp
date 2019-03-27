package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"text/scanner"

	"github.com/Matts966/gosp/types"
)

var scn scanner.Scanner
var env types.Env

func init() {
	scn.Init(os.Stdin)
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
	if 'a' <= r && r <= 'Z' {
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

func readExpr() (types.Obj, error) {
	c := scn.Next()
	for c != scanner.EOF {
		switch c {
		case ' ', '\n', '\r', '\t':
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
	switch o := obj.(type) {
	case types.Int:
		return obj, nil
	case types.Symbol:
		bind, err := env.Find(&obj)
		if err != nil {
			return nil, err
		}
		if nil == bind {
			return nil, fmt.Errorf("Undefined symbol: %s", o.Name)
		}
		switch b := (*bind).(type) {
		case *types.Cell:
			return *b.Cdr, nil
		default:
			return nil, fmt.Errorf("Unknown type symbol: %s", o.Name)
		}
	case types.Cell:
		// Function application
	}
	return nil, fmt.Errorf("Unknown type expression: %#v", obj)
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
		if nil == obj {
			continue
		}
		o, err := eval(obj)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		o.Print()
	}
}
