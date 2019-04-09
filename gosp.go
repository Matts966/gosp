package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Matts966/gosp/evaluator"
	"github.com/Matts966/gosp/prims"
	"github.com/Matts966/gosp/reader"
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

func repl(prompt string) {
L:
	for {
		fmt.Print(prompt)
		obj, err := reader.ReadExpr(&scn)
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch obj.(type) {
		case nil:
			fmt.Println(fmt.Errorf("reading expression returns nil"))
			os.Exit(1)
		case types.Dot:
			fmt.Println(fmt.Errorf("stray dot"))
			os.Exit(1)
		case types.RParen:
			fmt.Println(fmt.Errorf("unmatched right parenthesis"))
			os.Exit(1)
		// Do not return any
		case types.Comment:
			continue L
		}

		o, err := evaluator.Eval(&env, obj)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(o.String())
	}
}

func main() {
	if len(os.Args) < 2 {
		repl("gosp~> ")
	} else if "test" == os.Args[1] {
		repl("")
	}
	for i, fp := range os.Args {
		if 0 == i {
			continue
		}
		if strings.HasSuffix(fp, ".gosp") {
			f, err := os.Open(fp)
			if err != nil {
				panic(err)
			}
			scn.Init(f)
			repl("")
		}
	}
}
