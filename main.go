package main

func main() {}

type I interface {
	Param(interface {
		I
	})
	Return() interface {
		I
	}
}
