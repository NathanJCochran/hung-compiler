package main

func main() {}

type I interface {
	Param(i interface {
		I
	})
	Return() (i interface {
		I
	})
}
