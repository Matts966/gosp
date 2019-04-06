package scanner

import (
	"bufio"
	"io"
)

const (
	EOF = -(iota + 1)
)

type Scanner struct {
	reader *bufio.Reader
}

func (s *Scanner) Init(r io.Reader) {
	s.reader = bufio.NewReader(r)
}

func (s *Scanner) Peek() rune {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		panic(err)
	}
	s.reader.UnreadRune()
	return r
}

func (s *Scanner) Next() rune {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return EOF
		}
		panic(err)
	}
	return r
}
