package main

import (
	"bufio"
	"fmt"
	"os"
)

// Expr has data of expression.
type Expr string

func readExpr() Expr {
	var expr Expr
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		for _, c := range line {
			fmt.Println(c)
			fmt.Println('\n')
			switch c {
			case '\n':
				continue
			default:
				strExpr := string(expr)
				strExpr += string(c)
				expr = Expr(strExpr)
			}
		}
		return Expr(line)
	}
}

func eval(expr Expr) string {
	return string(expr)
}

func main() {
	for {
		expr := readExpr()
		print(eval(expr))
	}
}
