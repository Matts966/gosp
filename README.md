# Gosp

This repository is for my learning of Golang by implementing lisp refering to [minilisp](https://github.com/rui314/minilisp).

Test code is inspired by the [minilisp](https://github.com/rui314/minilisp) repositoy.

You can write lisp in your go source, if you want

```Go

obj, err := gosp.Interpret(`
(define sym (cons '111 '123)) 
(setcar sym '789) 
sym
`)
if err != nil {
  panic(err)
}
fmt.Println("returned:", obj.String())

// (111 . 123)
// (789 . 123)
// (789 . 123)
// returned: (789 . 123)
```
