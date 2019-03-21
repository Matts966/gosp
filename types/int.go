package types

import "fmt"

// Int type.
type Int struct {
	Value int
}

func (i Int) Print() {
	fmt.Print(i.Value)
}
