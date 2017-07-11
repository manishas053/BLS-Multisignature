package main

import "fmt"

func main() {
	
	var a [5]int
	fmt.Println("array1:", a)

	a[4] = 500
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	fmt.Println("length:",len(a))

	b := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("array2:",b)

	var twoD [2][3]int
	for i := 0; i<2; i++ {
		for j := 0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d:",twoD)
}
