package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf(" s= %v, len(s1) = %d, cap(s1) = %d \n", s, len(s), cap(s))
}

func main() {

	fmt.Println("Creating Slice")
	var s []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32) // len 16, cap 32
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying Slice")
	copy(s2, s1) // s1 的内容 copy 到 s2
	printSlice(s2)

	fmt.Println("Deleting elements from Slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	fmt.Println(front)
	s2 = s2[1:]
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	fmt.Println(tail)
	s2 = s2[:len(s2)-1]
	printSlice(s2)

}
