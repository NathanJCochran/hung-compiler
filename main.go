package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

type I interface {
	Param(i interface {
		I
	})
	Return() (i interface {
		I
	})
}
