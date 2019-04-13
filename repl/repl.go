package repl

import (
	"fmt"
	"io"
	"os"

	"github.com/Matts966/gosp/prims"
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/repl/reader"
	"github.com/Matts966/gosp/repl/scanner"
	"github.com/Matts966/gosp/types"
)

var (
	st          = "symbol_table"
	scn         scanner.Scanner
	env         types.Env = types.Env{}
	symbolTable *types.Cell
)

type repl struct {
	prompt      string
	scn         scanner.Scanner
	env         types.Env
	symbolTable *types.Cell
}

type Runnable interface {
	Run()
}

func New(r io.Reader, prompt string) Runnable {
	scn.Init(r)
	env.AddObj(st, nil)
	s, err := env.Find(st)
	if err != nil {
		panic(err)
	}
	ss := s.(*types.Cell)
	symbolTable = ss
	prims.AddPrims(&env)
	return &repl{
		prompt:      prompt,
		scn:         scn,
		env:         env,
		symbolTable: symbolTable,
	}
}

func (r *repl) Run() {
L:
	for {
		fmt.Print(r.prompt)
		obj, err := reader.ReadExpr(r.symbolTable, &r.scn)
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

		o, err := evaluator.Eval(&r.env, obj)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(o.String())
	}
}
