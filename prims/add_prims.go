package prims

import "github.com/Matts966/gosp/types"

var nameToPrims map[string]types.Obj = map[string]types.Obj{
	"t":      types.True{},
	"quote":  types.PrimFuncs{F: &PrimQuote},
	"+":      types.PrimFuncs{F: &PrimPlus},
	"-":      types.PrimFuncs{F: &PrimMinus},
	"<":      types.PrimFuncs{F: &PrimLessThan},
	"cons":   types.PrimFuncs{F: &PrimCons},
	"car":    types.PrimFuncs{F: &PrimCar},
	"cdr":    types.PrimFuncs{F: &PrimCdr},
	"define": types.PrimFuncs{F: &PrimDefine},
	"setcar": types.PrimFuncs{F: &PrimSetCar},
	"setq":   types.PrimFuncs{F: &PrimSetq},
	"if":     types.PrimFuncs{F: &PrimIf},
	"progn":  types.PrimFuncs{F: &PrimProgn},
	"=":      types.PrimFuncs{F: &PrimNumeq},
	"eq":     types.PrimFuncs{F: &PrimEq},
	"gensym": types.PrimFuncs{F: &PrimGensym},
	"intern": types.PrimFuncs{F: &PrimIntern},
	"lambda": types.PrimFuncs{F: &PrimLambda},
	"defun":  types.PrimFuncs{F: &PrimDefun},
	"print":  types.PrimFuncs{F: &PrimPrint},
	"server": types.PrimFuncs{F: &PrimServer},
}

func AddPrims(env *types.Env) {
	st, err := env.Find("symbol_table")
	if err != nil {
		panic(err)
	}
	for n, f := range nameToPrims {
		Intern(st.(*types.Cell), n)
		env.AddObj(n, f)
	}
}
