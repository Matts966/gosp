package types

import "fmt"

// Symbol type
type Symbol struct {
	Name string
}

func (s Symbol) Print() {
	fmt.Print(s.Name)
}
