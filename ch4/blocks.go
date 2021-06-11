package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// x := 10 // 10

	// if x > 5 {
	// 	println(x)

	// 	x := 5
	// 	println(x) // 5
	// }

	// println(x) // 10

	// x := 10
	// if x > 5 {
	// 	x, y := 5, 20
	// 	println(x, y)
	// }
	// println(x)

	// x := 10
	// fmt.Println(x)

	// fmt := "oops"
	// println(fmt)

	// println(true) // true
	// true := 10    // shadow true
	// println(true) // 10

	// n := rand.Intn(10) // i guess integer between 0 and 10

	// if n == 0 {
	// 	println("Too low")
	// } else if n > 5 {
	// 	println("This is too big")
	// } else {
	// 	println("A good number", n)
	// }

	// if n := rand.Intn(10); n == 0 { // here we basically say, if, when we create n, that n is 0
	// 	println("that is too low")
	// } else if n > 5 {
	// 	println("That is too big")
	// } else {
	// 	println("That is ok")
	// }

	// for i := 0; i < 10; i++ {
	// 	println(i)
	// }

	// i := 1
	// for i < 100 {
	// 	println(i)
	// 	i *= 2
	// }

	// for {
	// 	println("hello")
	// }

	// for {
	// 	if !CONDITON {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		println("FizBuzz")
	// 		continue
	// 	}

	// 	if i%3 == 0 {
	// 		println("Fizz")
	// 		continue
	// 	}

	// 	if i%5 == 0 {
	// 		println("Buzz")
	// 		continue
	// 	}

	// 	println(i)
	// }

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	println(v) // 0 2, 1 4, 2 6, 3 8, 4 10, 5 12
	// }

	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }

	// for i := 0; i < 3; i++ { // this just loops threee times, not sure why?
	// 	println("loop", i)
	// 	for k, v := range m {
	// 		println(k, v)
	// 	}
	// }

	// samples := []string{"hello", "apple_π!"}

	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		println(i, r, string(r))
	// 	}
	// 	println()
	// }

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }

	// fmt.Println(evenVals) // 2,4,6,8,10,12

	/* 	samples := []string{"hello", "apple"}

	outer:
		for _, sample := range samples {
			for i, r := range sample {
				fmt.Println(i, r, string(r))
				if string(r) == "l" {
					continue outer
				}
			}
		} */

	/* 	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	   	for _, word := range words { // so we have to loop over the thing to be able to get the value
	   		switch size := len(word); size { // we start switch, and we create a variable, then we say that we switch over that variable
	   		case 1, 2, 3, 4:
	   			fmt.Println(word, "is a short word")
	   		case 5:
	   			wordLen := len(word)
	   			fmt.Println(word, "is exactly the right length", wordLen)
	   		case 6, 7, 8, 9:
	   		default:
	   			fmt.Println(word, "is a long word")
	   		}
	   	} */

	// loop:
	// 	for i := 0; i < 10; i++ {
	// 		switch {
	// 		case i%2 == 0:
	// 			{

	// 				fmt.Println(i, "is event")
	// 			}
	// 		case i%7 == 0:
	// 			break loop
	// 		default:
	// 			fmt.Println("nothning here to see")
	// 		}
	// 	}

	// switch {
	// case a:
	// case b:
	// default:
	// }
	/*
		words := []string{"hi", "salutations", "hello"}
		for _, word := range words {
			switch wordLen := len(word); { // here we just declare a variable available during the switch
			case wordLen < 5:
				fmt.Println(word, "is a short word")
			case wordLen > 10:
				fmt.Println(word, "is a long word")
			default:
				fmt.Println(word, "is exactly the right lenght")
			}
		} */

	/* 	switch n := rand.Intn(10); {
	   	case n == 0:
	   		fmt.Println("something ")
	   	case n > 5:
	   		fmt.Println("something else ")
	   	default:
	   		fmt.Println("default")
	   	} */

	/* 	a := 10
		goto skip
		b := 20

	skip:
		c := 30
		fmt.Println(a, b, c)
		if c > a {
			goto inner
		}

		if a < b {
		inner:
			fmt.Println("a is less than b")
		} */

	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do somehting")
done:
	fmt.Println("do some complicated stuff no matter what ")
	fmt.Println(a)
}
