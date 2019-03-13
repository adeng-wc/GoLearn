package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 11
	bb = "bb"
)

func variableZeroValue() {
	var a int
	var s string

	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, s, b)
	fmt.Printf("%d %q\n", a, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, true, "sss", 1.1
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 1, true, "sss", 1.1
	fmt.Println(a, b, c, d)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// e的 i * pi 次方  指数  + 1
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	// 百分号 3 位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}
func triangles() {
	var a, b = 3, 4
	fmt.Println(calcTriangles(a, b))
}

func calcTriangles(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		str  = "aa"
		a, b = 3, 4
	)

	var c = int(math.Sqrt(a*a + b*b))

	fmt.Println(a, b, str, c)
}
func enums() {

	const (
		java = iota
		_
		python
		goload
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(java, python, goload)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("hello world")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	euler()
	triangles()
	consts()
	enums()
}
