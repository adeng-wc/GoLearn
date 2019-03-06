package main

import "fmt"

func main() {
	//a := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(a(i))
	//}

	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Println(i, s)
	}
}

// 返回 闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// "正统" 闭包
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}
