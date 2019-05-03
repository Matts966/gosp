package gosp

import (
	"log"
)

// Acceptance tests for the gosp interpreter.
func ExampleSetcar() {
	_, err := Interpret(
		`
		(define sym (cons '111 '123)) 
		(setcar sym '789) 
		sym
		`,
	)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// Output:
	// (111 . 123)
	// (789 . 123)
	// (789 . 123)
}

func ExampleReccursion() {
	_, err := Interpret(
		`
		(defun fib (n) 
			(if (= n 1) 1
				(if (= n 2) 1
					(+ (fib (- n 1)) (fib (- n 2))))))
		(fib 10)
		`,
	)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// Output:
	// 55
}

func ExampleVariadicFunction() {
	_, err := Interpret(
		`
		(defun vari (x . y) 
			(cons x (car y)))
		(vari 1 '(1 1))
		(vari 1 '(2 1))
		(vari 1 '(3 2 1))
		`,
	)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// Output:
	// (1 . 1)
	// (1 . 2)
	// (1 . 3)
}
