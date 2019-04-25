package prims

import (
	"github.com/Matts966/gosp/types"
	"golang.org/x/xerrors"
)

// PrimLessThan is primitive function in form of (> a b).
var PrimLessThan types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	argList := *args
	l, err := argList.Length()
	if err != nil {
		return nil, xerrors.Errorf("failed to get the length of args, err: %v", err)
	}
	if 2 != l {
		return nil, xerrors.New("malformed lt")
	}
	x, ok := argList.Car.(types.Int)
	if !ok {
		return nil, xerrors.Errorf("lt takes only int value, but the first argument was %#v",
			argList.Car)
	}
	c, ok := argList.Cdr.(*types.Cell)
	if !ok {
		return nil, xerrors.New("list should be pointer to types.Cell")
	}
	y, ok := c.Car.(types.Int)
	if !ok {
		return nil, xerrors.Errorf("lt takes only int value, but the second argument was %#v",
			argList.Car)
	}
	if x.Value < y.Value {
		return types.True{}, nil
	}

	return types.False{}, nil

}
