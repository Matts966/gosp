package scanner_test

import (
	"os"
	"strings"
	"testing"

	"github.com/Matts966/gosp/scanner"
)

var scn scanner.Scanner

const str string = "abcdefghijklmnopqrstuvwxyz-1234567890 \r\n\n"

func TestPeek(t *testing.T) {
	scn.Init(strings.NewReader(str))
	p := scn.Peek()
	for range str {
		p2 := scn.Peek()
		if p != p2 {
			t.Errorf("Peek failed, expected: %v, actual %v.", p, p2)
		}
	}
}

func TestNext(t *testing.T) {
	scn.Init(strings.NewReader(str))
	for _, s := range str {
		s2 := scn.Next()
		if s != s2 {
			t.Errorf("Peek failed, expected: %v, actual %v.", s, s2)
		}
	}
}

func TestEOF(t *testing.T) {
	scn.Init(strings.NewReader(str))
	for {
		if scanner.EOF == scn.Peek() {
			t.Logf("%v", scn.Next())
			break
		} else {
			t.Logf("%v", scn.Next())
		}
	}
}

func TestEOFUsingFile(t *testing.T) {
	f, err := os.Open(`./test_data.txt`)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	scn.Init(f)
	for {
		if scanner.EOF == scn.Peek() {
			t.Logf("%v", scn.Next())
			break
		} else {
			t.Logf("%v", scn.Next())
		}
	}
}

func TestBack(t *testing.T) {
	scn.Init(strings.NewReader(str))
	for {
		if scanner.EOF == scn.Peek() {
			scn.Next()
			scn.Back()
			if scanner.EOF != scn.Peek() {
				t.Fatal("Failed to back properly.")
			}
			break
		} else {
			n := scn.Next()
			scn.Back()
			if scn.Next() != n {
				t.Fatal("Failed to back properly.")
			}
		}
	}
}
