package types

import (
	"strconv"
)

// Int type.
type Int struct {
	Value int
}

func (i Int) String() string {
	return strconv.Itoa(i.Value)
}
