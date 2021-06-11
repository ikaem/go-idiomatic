package main

import (
	"fmt"
	"strconv"
)

// func div(numerator int, denominator int) int {
// 	if denominator == 0 {
// 		return 0
// 	}
// 	return numerator / denominator
// }

// type MyFuncOpts struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func MyFunc(opts MyFuncOpts) {
// 	fmt.Println("this is print line", opts)
// }

// func addTo(base int, vals ...int) []int {
// 	out := make([]int, 0, len(vals)) // this creates a slice, initializes array of length 0, and sets capacity to 10)
// 	for _, v := range vals {
// 		out = append(out, base+v)
// 	}

// 	return out
// }

// func divAndremainder(numerator, denominator int) (int, int, error) {
// 	if denominator == 0 {
// 		return 0, 0, errors.New("Cannot divide by zero") // we had to import errors for this
// 	}
// 	return numerator / denominator, numerator % denominator, nil
// }

// func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) { // we have to name thse returns in the return types
// 	if denominator == 0 {
// 		err = errors.New("Cannot divide by zero")
// 		return result, remainder, err // i guess this will use zero values for these other variables, based on their types
// 	}
// 	result, remainder = numerator/denominator, numerator%denominator // i like this // just assign bunch of calculated values to bunch of calculated values

// 	// return result, remainder, err                  // same thing with error here
// 	return numerator / denominator, remainder, err // same thing with error here
// }

// func divAndremainder(numerator, denominator int) (result int, remainder int, err error) {
// 	if denominator == 0 {
// 		err = errors.New("Cannot divide by zero")
// 		return
// 	}

// 	result, remainder = numerator/denominator, numerator%denominator
// 	return
// }

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {

	// result := div(5, 2)
	// fmt.Println(result)

	// MyFunc(MyFuncOpts{LastName: "Patel"})
	// MyFunc(MyFuncOpts{FirstName: "No"})

	// fmt.Println(addTo(3))             // returns nothing, because we never loop because there is nothing to loop over
	// fmt.Println(addTo(3, 2))          // returnrs 5
	// fmt.Println(addTo(3, 2, 4, 6, 8)) // 5, 7, 9, 11
	// a := []int{4, 3}
	// fmt.Println(addTo(3, a...)) // this spreads a slice, and it makes it 3, 4, 3 as argumebts

	// result, remainder, err := divAndremainder(5, 2)
	// divAndremainder(5, 2)
	// _, _, err := divAndremainder(5, 2)

	// if err != nil { // check if there was error // cannot use if err
	// 	fmt.Println(err)
	// 	os.Exit(1) // not sure what this is
	// }
	// fmt.Println(result, remainder)

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"five", "/", "six"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}

}
