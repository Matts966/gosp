package prims

import "github.com/Matts966/gosp/types"

var nameToFunc map[string]types.Prim = map[string]types.Prim{
	"quote":  types.Prim{F: &PrimQuote},
	"+":      types.Prim{F: &PrimPlus},
	"-":      types.Prim{F: &PrimMinus},
	"<":      types.Prim{F: &PrimLessThan},
	"cons":   types.Prim{F: &PrimCons},
	"car":    types.Prim{F: &PrimCar},
	"cdr":    types.Prim{F: &PrimCdr},
	"define": types.Prim{F: &PrimDefine},
	"setcar": types.Prim{F: &PrimSetCar},
	"setq":   types.Prim{F: &PrimSetq},
	"if":     types.Prim{F: &PrimIf},
	"progn":  types.Prim{F: &PrimProgn},
	"=":      types.Prim{F: &PrimNumeq},
	"eq":     types.Prim{F: &PrimEq},
	"gensym": types.Prim{F: &PrimGensym},
	"intern": types.Prim{F: &PrimIntern},
}

func AddPrims(env *types.Env) {
	st, err := env.Find("symbol_table")
	if err != nil {
		panic(err)
	}
	for n, f := range nameToFunc {
		Intern(st.(*types.Cell), n)
		env.AddObj(n, f)
	}
}
