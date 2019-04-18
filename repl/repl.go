package repl

import (
	"fmt"
	"io"

	"github.com/Matts966/gosp/prims"
	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/repl/reader"
	"github.com/Matts966/gosp/repl/scanner"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
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
	Run() (types.Obj, error)
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

func (r *repl) Run() (types.Obj, error) {
	var lastObj types.Obj
L:
	for {
		fmt.Print(r.prompt)
		obj, err := reader.ReadExpr(r.symbolTable, &r.scn)
		if err == io.EOF {
			return lastObj, nil
		}
		if err != nil {
			return nil, xerrors.Errorf("error occured while ReadExpr in repl.Run symbolTable=%#v : %w", r.symbolTable, err)
		}

		switch obj.(type) {
		case nil:
			return nil, xerrors.New("reading expression returns nil")
		case types.Dot:
			return nil, xerrors.New("stray dot")
		case types.RParen:
			return nil, xerrors.New("unmatched right parenthesis")
		// Do not return any
		case types.Comment:
			continue L
		}

		o, err := evaluator.Eval(&r.env, obj)
		if nil != err {
			return nil, xerrors.Errorf("error occured while evaluating obj in repl.Run obj=%#v : %w", obj, err)
		}
		fmt.Println(o.String())
		lastObj = o
	}
}
