package main

import (
	"fmt"
	"math"
)

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s ", op)
	}
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ... int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

// *int 指针
func swap(a, b *int) {
	*b, *a = *a, *b
}
func swap2(a, b int) (int, int) {
	return b, a
}
func main() {

	if result, err := eval(3, 4, "+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(eval(3, 4, "x"))

	q, _ := div(4, 3)
	fmt.Println(div(4, 3))
	fmt.Println(q)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4))

	a, b := 3, 4
	// &a 的内存地址
	//swap(&a, &b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
