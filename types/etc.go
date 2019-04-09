package types

// Dot type.
type Dot struct{}

func (d Dot) toString() string {
	return "."
}

// RParen type
type RParen struct{}

func (r RParen) toString() string {
	return ")"
}

// True type
type True struct{}

func (r True) toString() string {
	return "t"
}

// False type
type False struct{}

func (f False) toString() string {
	return "()"
}

// Comment type
type Comment struct{ String string }

func (c Comment) toString() string {
	return c.String
}
