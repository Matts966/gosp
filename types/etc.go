package types

import "fmt"

// Dot type.
type Dot struct {}

func (d Dot) Print() {
	fmt.Print(".")
}

// RParen type
type RParen struct {}

func (r RParen) Print() {
	fmt.Print(")")
}

// True type
type True struct {}

func (r True) Print() {
	fmt.Print("t")
}
