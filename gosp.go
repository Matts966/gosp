package main

import (
	"fmt"
	"io"
	"os"
	"text/scanner"
)

var scn scanner.Scanner

func init() {
	scn.Init(os.Stdin)
}

// Expr has data of expression.
type Expr string

func readExpr() (Expr, error) {
	var expr Expr
	c := scn.Next()
	for c != scanner.EOF {
		switch c {
		case ' ', '\n', '\r', '\t':
			return expr, nil
		default:
			strExpr := string(expr)
			strExpr += string(c)
			expr = Expr(strExpr)
		}
		c = scn.Next()
	}

	return expr, io.EOF
}

func eval(expr Expr) string {
	return string(expr)
}

func main() {
	for {
		expr, err := readExpr()
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(eval(expr))
	}
}
