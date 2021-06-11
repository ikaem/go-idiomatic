package main

import "fmt"

func main() {
	e := 16
	f, g := 11, 6
	h := 66

	const xe int = 64
	const (
		what = "id"
		no   = "huh?"
	)

	fmt.Println("hello world")

	var (
		a int
		b = 20
		c string
		d bool = false
	)

	var p, q int = 10, 20
	var r, s int
	var t, v = 10, "hello"

	var x = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)

	var isTrue bool = (x == 10)
}
