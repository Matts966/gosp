package prims

import (
	"fmt"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimPrint is primitive function for printing object.
var PrimPrint types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	obj, err := evaluator.Eval(env, args.Car)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in print caused error: %w", err)
	}
	fmt.Println(obj.String())
	return types.False{}, nil
}
