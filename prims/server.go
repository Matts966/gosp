package prims

import (
	"net/http"
	"strconv"

	"github.com/Matts966/gosp/repl/evaluator"
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimServer is primitive function for setting up simple http file server.
var PrimServer types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	port := 80
	obj, err := evaluator.Eval(env, args.Car)
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in server caused error: %w", err)
	}
	if i, ok := obj.(types.Int); ok {
		port = i.Value
	}

	fileServer := http.StripPrefix("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":"+strconv.Itoa(port), fileServer)
	return types.False{}, nil
}
