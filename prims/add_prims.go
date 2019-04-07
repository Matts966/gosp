package prims

import "github.com/Matts966/gosp/types"

func AddPrims(env *types.Env) {
	env.AddObj("quote", &PrimQuote)
	env.AddObj("+", &PrimPlus)
	env.AddObj("-", &PrimMinus)
}
