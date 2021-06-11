package main

import "fmt"

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const untypedConstant = 10

	var a int = untypedConstant
	var b float64 = untypedConstant
	var c byte = untypedConstant

	const typeConstant int = 10

	var noWork float64 = typeConstant

	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
