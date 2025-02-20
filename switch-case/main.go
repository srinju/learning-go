package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("print ", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	whatamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("type is int")
		case bool:
			fmt.Println("case is bool")
		default:
			fmt.Println("dont know type  : ", t)
		}
	}

	whatamI(2)
	whatamI(false)
	whatamI("hey")
}
