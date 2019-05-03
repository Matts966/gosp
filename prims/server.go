package prims

import (
	"net/http"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimServer is primitive function for setting up simple http file server.
var PrimServer types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	obj, err := evaluator.Eval(env, args.Car)
	address := "0.0.0.0:80"
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in server caused error: %w", err)
	}
	if s, ok := obj.(*types.Symbol); ok {
		if s != nil {
			address = *s.Name
		}
	}

	fileServer := http.StripPrefix("/", http.FileServer(http.Dir(".")))
	return types.False{}, xerrors.Errorf("setting up file server failed err: %+v", http.ListenAndServe(address, fileServer))
}
