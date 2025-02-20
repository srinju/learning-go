package main

import "fmt"

func main() {

	var a [5]int

	fmt.Println("emp arr by default : ", a)

	a[4] = 100
	fmt.Println("set : ", a)
	fmt.Println("get : ", a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//with ... we can make the compiler count the number of elements in the array
	//b = [...]int{1,2,3,4,5}
	fmt.Println("array b is : ", b)
	fmt.Println("length of array b : ", len(b))

	b = [...]int{100, 3: 400, 500} //3rd index e 400 then 500 orom bhabe
	fmt.Println(b)

	b = [...]int{0, 1, 4: 4} //4th index e 4 orm bhabe
	fmt.Println(b)

	//two d arrays >>

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD arr : ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
