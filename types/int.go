package types

import (
	"strconv"
)

// Int type.
type Int struct {
	Value int
}

func (i Int) toString() string {
	return strconv.Itoa(i.Value)
}
