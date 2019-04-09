package reader

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Matts966/gosp/scanner"
	"github.com/Matts966/gosp/types"
)

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

func readQuote(scn *scanner.Scanner) (types.Obj, error) {
	obj, err := ReadExpr(scn)
	if err != nil {
		return nil, fmt.Errorf("failed to read quote")
	}
	return types.Cons(types.Symbol{
		Name: "quote",
	}, types.Cons(obj, nil)), nil
}

func readNumber(scn *scanner.Scanner, v int) int {
	for isDigit(scn.Peek()) {
		v = v*10 + int(scn.Next()-'0')
	}
	return v
}

func readSymbol(scn *scanner.Scanner, c rune) *types.Symbol {
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

func readList(scn *scanner.Scanner) (types.Obj, error) {
	obj, err := ReadExpr(scn)
	if err != nil {
		return obj, err
	}
	switch obj.(type) {
	case nil:
		obj, err := ReadExpr(scn)
		if err != nil {
			return obj, err
		}
		if _, ok := obj.(types.RParen); ok {
			return types.Cell{}, nil
		}
		return nil, fmt.Errorf("unclosed parenthesis")
	case types.Dot:
		return nil, fmt.Errorf("stray dot")
	case types.RParen:
		return types.False{}, nil
	}
	head := types.Cons(obj, nil)

	var tail types.Obj
	tail = &head
	for {
		obj, err := ReadExpr(scn)
		if err != nil {
			return nil, err
		}
		switch obj.(type) {
		case nil:
			return nil, fmt.Errorf("unclosed parenthesis")
		case types.Dot:
			t, _ := tail.(*types.Cell)
			t.Cdr, err = ReadExpr(scn)
			if err != nil {
				return nil, err
			}
			nx, err := ReadExpr(scn)
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

func passLine(scn *scanner.Scanner) (types.Obj, error) {
	c := scn.Next()
	comment := string(c)
	for '\n' != c && scanner.EOF != c {
		// read Carriage Return
		if '\r' == c {
			if '\n' == scn.Peek() {
				scn.Next()
				break
			}
		}
		comment += string(c)
		c = scn.Next()
		continue
	}
	return types.Comment{comment}, nil
}

func ReadExpr(scn *scanner.Scanner) (types.Obj, error) {
	c := scn.Next()
	for c != scanner.EOF {
		switch c {
		case ' ', '\n', '\r', '\t':
		case '(':
			return readList(scn)
		case ')':
			return types.RParen{}, nil
		case '.':
			return types.Dot{}, nil
		case '\'':
			return readQuote(scn)
		case ';':
			return passLine(scn)
		case '-':
			if isDigit(scn.Peek()) {
				return types.Int{
					Value: -readNumber(scn, int(scn.Next()-'0')),
				}, nil
			}
			fallthrough
		default:
			if isDigit(c) {
				return types.Int{
					Value: readNumber(scn, int(c-'0')),
				}, nil
			}
			return readSymbol(scn, c), nil
		}
		c = scn.Next()
	}
	return nil, io.EOF
}
