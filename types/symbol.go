package types

import "fmt"

// Symbols runes
const Symbols string = "~!@#$%^&*-_=+:/?<>"

// Symbol type
type Symbol struct {
	Name string
}

func (s Symbol) Print() {
	fmt.Print(s.Name)
}
