package main

import "fmt"

func main() {

	var a = "intial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// if any value is not assigned it is 0
	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//or var foo int g = 2
	var foo int = 2
	g := 2
	fmt.Println(g)
	fmt.Println(foo)

}
