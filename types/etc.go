package types

// Dot type.
type Dot struct{}

func (d Dot) String() string {
	return "."
}

// RParen type
type RParen struct{}

func (r RParen) String() string {
	return ")"
}

// True type
type True struct{}

func (r True) String() string {
	return "t"
}

// False type
type False struct{}

func (f False) String() string {
	return "()"
}

// Comment type
type Comment struct{ Content string }

func (c Comment) String() string {
	return c.Content
}
