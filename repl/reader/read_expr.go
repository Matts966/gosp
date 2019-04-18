package reader

import (
	"io"
	"strconv"

	"github.com/Matts966/gosp/prims"
	"github.com/Matts966/gosp/repl/scanner"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
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

func readSymbol(symbolTable *types.Cell, scn *scanner.Scanner) (*types.Symbol, error) {

	n := scn.Peek()
	name := ""
	for isDigit(n) || isSymbolRune(n) || '-' == n {

		name += string(scn.Next())
		n = scn.Peek()
	}
	obj, err := prims.Intern(symbolTable, name)
	if err != nil {
		return nil, xerrors.Errorf("intern failed in readSymbol, symbolTable: %#v, name: %s", symbolTable, name)
	}

	return obj.(*types.Symbol), nil
}

func readQuote(symbolTable *types.Cell, scn *scanner.Scanner) (types.Obj, error) {
	cont, err := ReadExpr(symbolTable, scn)
	if err != nil {
		return nil, xerrors.Errorf("failed to read quote, err: %+v", err)
	}
	i, err := prims.Intern(symbolTable, "quote")
	if err != nil {
		return nil, xerrors.Errorf("interning \"quote\" failed in readQuote caused error: %w", err)
	}
	// is, err := prims.Intern(symbolTable, cont.String())
	return types.Cons(i, types.Cons(cont, nil)), nil
}

func readNumber(scn *scanner.Scanner, v int) int {
	for isDigit(scn.Peek()) {
		v = v*10 + int(scn.Next()-'0')
	}
	return v
}

func readList(symbolTable *types.Cell, scn *scanner.Scanner) (types.Obj, error) {
	obj, err := ReadExpr(symbolTable, scn)
	if err != nil {
		return obj, xerrors.Errorf("error occured while ReadExpr in readList symbolTable=%#v : %w", symbolTable, err)
	}
	switch obj.(type) {
	case nil:
		obj, err := ReadExpr(symbolTable, scn)
		if err != nil {
			return obj, xerrors.Errorf("error occured while 2nd ReadExpr in readList, symbolTable=%#v : %w", symbolTable, err)
		}
		if _, ok := obj.(types.RParen); ok {
			return types.Cell{}, nil
		}
		return nil, xerrors.New("unclosed parenthesis")
	case types.Dot:
		return nil, xerrors.New("stray dot")
	case types.RParen:
		return types.False{}, nil
	}
	head := types.Cons(obj, nil)

	var tail types.Obj
	tail = head
	for {
		obj, err := ReadExpr(symbolTable, scn)
		if err != nil {
			return nil, xerrors.Errorf("error occured while 3rd ReadExpr in readList symbolTable=%#v : %w", symbolTable, err)
		}
		switch obj.(type) {
		case nil:
			return nil, xerrors.New("unclosed parenthesis")
		case types.Dot:
			t, _ := tail.(*types.Cell)
			t.Cdr, err = ReadExpr(symbolTable, scn)
			if err != nil {
				return nil, xerrors.Errorf("error occured while ReadExpr for dotted list in readList symbolTable=%#v : %w", symbolTable, err)
			}
			nx, err := ReadExpr(symbolTable, scn)
			if err != nil {
				return nil, xerrors.Errorf("error occured while 2nd ReadExpr for dotted list in readList symbolTable=%#v : %w", symbolTable, err)
			}
			if _, ok := nx.(types.RParen); !ok {
				return nil, xerrors.New("closed parenthesis expected after dot")
			}
			return head, nil
		case types.RParen:
			return head, nil
		default:
			t, _ := tail.(*types.Cell)
			nc := types.Cons(obj, nil)
			t.Cdr = nc
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

func ReadExpr(symbolTable *types.Cell, scn *scanner.Scanner) (types.Obj, error) {
	p := scn.Peek()
	for p != scanner.EOF {
		switch p {
		case ' ', '\n', '\r', '\t':
			scn.Next()
		case '(':
			scn.Next()
			return readList(symbolTable, scn)
		case ')':
			scn.Next()
			return types.RParen{}, nil
		case '.':
			scn.Next()
			return types.Dot{}, nil
		case '\'':
			scn.Next()
			return readQuote(symbolTable, scn)
		case ';':
			scn.Next()
			return passLine(scn)
		case '-':

			scn.Next()

			if isDigit(scn.Peek()) {

				return types.Int{
					Value: -readNumber(scn, int(scn.Next()-'0')),
				}, nil
			}

			// Can not use scn.Back() for used scn.Peek() already.
			return prims.Intern(symbolTable, "-")

		default:
			if isDigit(p) {

				scn.Next()

				return types.Int{
					Value: readNumber(scn, int(p-'0')),
				}, nil
			}

			return readSymbol(symbolTable, scn)
		}
		p = scn.Peek()
	}
	return nil, io.EOF
}
