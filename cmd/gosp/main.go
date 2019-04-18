package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Matts966/gosp/repl"
	"golang.org/x/xerrors"
)

const (
	prompt = "gosp~> "
	ext    = ".gosp"
)

var r repl.Runnable

func init() {
	if len(os.Args) < 2 {
		r = repl.New(os.Stdin, prompt)
	} else if "test" == os.Args[1] {
		r = repl.New(os.Stdin, "")
	}
	for i, fp := range os.Args {
		if 0 == i {
			continue
		}
		if strings.HasSuffix(fp, ext) {
			f, err := os.Open(fp)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
			r = repl.New(f, "")
		}
	}
	if nil == r {
		fmt.Printf("%+v\n", xerrors.New("file not found"))
	}
}

func main() {
	_, err := r.Run()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
