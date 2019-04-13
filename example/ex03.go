package main

import (
	"fmt"

	"github.com/Matts966/gosp"
)

func main() {
	obj, err := gosp.Interpret(`
	(define sym (cons '111 '123)) 
	(setcar sym '789) 
	sym
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj.String())
}
