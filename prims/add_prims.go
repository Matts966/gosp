package prims

import "github.com/Matts966/gosp/types"

func AddPrims(env *types.Env) {
	env.AddObj("quote", types.Prim{F: &PrimQuote})
	env.AddObj("+", types.Prim{F: &PrimPlus})
	env.AddObj("-", types.Prim{F: &PrimMinus})
	env.AddObj("<", types.Prim{F: &PrimLessThan})
	env.AddObj("cons", types.Prim{F: &PrimCons})
	env.AddObj("car", types.Prim{F: &PrimCar})
	env.AddObj("cdr", types.Prim{F: &PrimCdr})
	env.AddObj("define", types.Prim{F: &PrimDefine})
	env.AddObj("setcar", types.Prim{F: &PrimSetCar})
	env.AddObj("setq", types.Prim{F: &PrimSetq})
	env.AddObj("if", types.Prim{F: &PrimIf})
	env.AddObj("progn", types.Prim{F: &PrimProgn})
	env.AddObj("=", types.Prim{F: &PrimNumeq})
	env.AddObj("eq", types.Prim{F: &PrimEq})
}
