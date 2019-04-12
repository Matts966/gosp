package types

// Symbols runes
const Symbols string = "~!@#$%^&*-_=+:/?<>"

// Symbol type
type Symbol struct {
	Name *string
}

func (s Symbol) String() string {
	return *s.Name
}
