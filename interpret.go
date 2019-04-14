package gosp

import (
	"go/scanner"
	"strings"

	"github.com/Matts966/gosp/repl"
	"github.com/Matts966/gosp/types"
)

var (
	st          = "symbol_table"
	scn         scanner.Scanner
	env         types.Env = types.Env{}
	symbolTable *types.Cell
)

func Interpret(lispStr string) (types.Obj, error) {
	r := repl.New(strings.NewReader(lispStr), "")
	return r.Run()
}
