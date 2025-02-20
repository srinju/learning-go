package main

import "fmt"

func main() {
	n := 4
	if n%2 == 0 {
		fmt.Println("n is even")
	} else {
		fmt.Println("n is odd!!")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even!!!")
	}

	if 8%2 == 0 || 4%2 == 0 {
		fmt.Println("both 8 and 4 are even")
	} else {
		fmt.Println("both are odd!")
	}
}
