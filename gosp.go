package main

import (
	"fmt"
	"io"
	"os"
	"text/scanner"
)

// Expr has data of expression.
type Expr string

func readExpr() (Expr, error) {
	var expr Expr

	var scn scanner.Scanner
	scn.Init(os.Stdin)
	c := scn.Next()
	for c != scanner.EOF {
		switch c {
		case ' ', '\n', '\r', '\t':
		default:
			strExpr := string(expr)
			strExpr += string(c)
			expr = Expr(strExpr)
			return expr, nil
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
