package main

func main() {

	// var x [3]int

	// var y = [3]int{1, 2, 3}

	// var z = [12]int{1, 5: 4, 6, 10: 100, 15}

	// var a = [...]int{1, 2, 3}

	// fmt.Println(y == a)

	// var b [2][3]int

	// fmt.Println(len(x))
	// var c = []int{10, 20, 30}
	// var e [][]int

	// var f []int
	// f = append(f, 10)
	// f = append(f, 12, 13, 114)
	// var d = []int{1, 5: 6, 7, 10, 11: 12}
	// f = append(f, d...)

	// fmt.Println(f)
	// fmt.Println(len(f))
	// fmt.Println(cap(f))

	// var x []int
	// fmt.Println(x, len(x), cap(x))

	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))

	// y := make([]int, 10, 20)

	// var data []int
	// var data2 = []int{}

	// defaultValues := []int{2, 4, 6, 8}

	/* 	x := []int{1, 2, 3, 4}
	   	y := x[:2]
	   	a := x[2:]
	   	b := x[1:3]
	   	c := x[:]
	*/

	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// z := x[1:]

	// x[1] = 20
	// y[0] = 10
	// z[1] = 30

	// a := x[:2:2]

	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// this is array
	// c := [4]int{5, 6, 7, 8}

	// this is now slice made after slicing hte above array
	// d := c[:2]

	// 	x := []int{1, 2, 3, 4}
	// 	y := make([]int, 4)

	// 	num := copy(y, x)
	// 	fmt.Println(y, num)
	// }

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// copy(y, x)

	// x := []int{1, 2, 3, 4}

	// y := make([]int, 2)
	// copy(y, x[2:])

	// x := []int{1, 2, 3, 4}
	// num := copy(x[:3], x[1:])
	// fmt.Println(x, num)

	// var s string = "hello there"

	// var b byte = s[6]
	// fmt.Println(b)

	// var s2 string = s[4:7]

	// fmt.Println(s2)

	// var a rune = 'x'
	// fmt.Println(a) // this printes 120 - utf code for x?

	// var s string = string(a)
	// fmt.Println(s) // prints x

	// var b byte = 'y'
	// fmt.Println(b) // prints 121

	// var s2 = string(b)
	// println(s2) // prints y

	/* 	var nilMap map[string]int

	   	println(len(nilMap)) // 0
	*/
	// totalWins := map[string]int{}

	// teams := map[string][]string{
	// 	"Orcas":   []string{"Fred", "Ralph", "Bijou"},
	// 	"Lions":   []string{"Karlo", "Ivan", "Martina"},
	// 	"Kittens": []string{"Jela", "Marko", "What"},
	// }

	// ages := make(map[int][]string, 10)

	// println(len(ages))

	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2

	// println(totalWins["Orcas"])
	// println(totalWins["Kittens"])
	// totalWins["Kittens"]++
	// println((totalWins["Kittens"]))
	// totalWins["Lions"] = 3
	// println(totalWins["Lions"])

	/* 	m := map[string]int{
	   		"hello": 5,
	   		"world": 0,
	   	}

	   	v, ok := m["hello"]
	   	println(v, ok) // 5, true

	   	v, ok = m["world"]
	   	println(v, ok) // 0 true

	   	v, ok = m["goodbye"]
	   	println(v, ok) // 0 false */

	// m := map[string]int{
	// 	"hello": 5,
	// 	"world": 10,
	// }

	// delete(m, "hello")

	// intSet := map[int]bool{}
	// vals := []int{5, 10, 2, 6, 2, 11}

	// for _, v := range vals {
	// 	intSet[v] = true
	// }

	// println("lenght of vals", len(vals), "length of the set", len(intSet))

	// println(intSet[5])
	// println(intSet[500])

	// if intSet[100] {
	// 	println("100 is in the set")
	// } else {
	// 	println("100 is not in the set")
	// }

	// type person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }

	// var fred person

	// bob := person{}

	// julia := person{
	// 	"Julia",
	// 	32,
	// 	"dog",
	// }

	// beth := person{
	// 	name: "Beth",
	// 	pet:  "dog",
	// }

	// beth.name = "not beth anymores"

	var car struct {
		make string
		age  int
	}

	car.make = "bmw"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
}
