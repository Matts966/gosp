package prims

import "github.com/Matts966/gosp/types"

func AddPrims(env *types.Env) {
	env.AddObj("quote", &PrimQuote)
	env.AddObj("+", &PrimPlus)
	env.AddObj("-", &PrimMinus)
	env.AddObj("<", &PrimLessThan)
	env.AddObj("cons", &PrimCons)
	env.AddObj("car", &PrimCar)
	env.AddObj("cdr", &PrimCdr)
	env.AddObj("define", &PrimDefine)
	env.AddObj("setcar", &PrimSetCar)
	env.AddObj("setq", &PrimSetq)
	env.AddObj("if", &PrimIf)
	env.AddObj("progn", &PrimProgn)
	env.AddObj("=", &PrimNumeq)
}
