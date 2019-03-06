package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

func main() {

	f := fibonacci()

	//fmt.Println(f()) //1
	//fmt.Println(f()) //1
	//fmt.Println(f()) //2
	//fmt.Println(f()) //3
	//fmt.Println(f()) //5
	//fmt.Println(f()) //8
	//fmt.Println(f()) //13

	printFileContents(f)
}
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 是类型就能实现接口
type intGen func() int

// 实现 /usr/local/go/src/io/io.go:77 接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)

	//TODO : incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
