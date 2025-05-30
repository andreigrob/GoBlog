package main

import (
	ft "fmt"
)

type Object struct {
}

func (*Object) Print() {
	ft.Println("Object")
}

type P interface {
	Print()
}

func Print[T P]() {
	var val T
	ft.Println(val)
	val.Print()
}

func main() {
	Print[*Object]()
}
