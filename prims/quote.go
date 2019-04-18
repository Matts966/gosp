package prims

import (
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimQuote is in form of (quote ~).
var PrimQuote types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, xerrors.Errorf("evaluating args in quote caused error: %w", err)
	}
	if 1 != l {
		return nil, xerrors.New("malformed quote")
	}
	return argList.Car, nil
}
