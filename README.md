## hung-compiler

This program causes the go compiler to hang indefinitely:

```go
package main

func main() {}

type I interface {
	M1() interface {
		I
	}
	M2() interface {
		I
	}
}
```

If you remove either of the two interface method definitions (M1 or M2), it
compiles. Furthermore, it seems to exhibit the same behavior whether the
anonymous interfaces are for the method return types, parameter types, or a
mixture of the two.

Tested with go1.10.2 (5/5/2018)

### Try it yourself

Download the code:

`go get -d github.com/nathanjcochran/hung-compiler`

Try to compile it:

`go install github.com/nathanjcochran/hung-compiler`
