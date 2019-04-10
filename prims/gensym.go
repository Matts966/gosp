package prims

import (
	"strconv"

	"github.com/Matts966/gosp/types"
)

var count = -1

// PrimGensym generates new symbol in form of (gensym).
var PrimGensym types.PF = func(env *types.Env, args *types.Cell) (types.Obj, error) {
	count++
	return types.Symbol{Name: "G__" + strconv.Itoa(count)}, nil
}
