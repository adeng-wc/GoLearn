package main

import "fmt"

func printArr(arr [5] int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArr2(arr *[5]int) {
	(*arr)[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}
	var grid [3][4]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	printArr(arr3)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArr2(&arr3)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

}
