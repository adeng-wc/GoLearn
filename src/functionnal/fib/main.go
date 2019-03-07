package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
	"os"
)

func main() {

	//f := fibonacci()
	//
	////fmt.Println(f()) //1
	////fmt.Println(f()) //1
	////fmt.Println(f()) //2
	////fmt.Println(f()) //3
	////fmt.Println(f()) //5
	////fmt.Println(f()) //8
	////fmt.Println(f()) //13
	//
	//printFileContents(f)

	writeFile("fib.txt")
}

// 使用 defer 定义 函数结束时发生
func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 实现 /usr/local/go/src/errors/errors.go:18 接口
	//err = errors.New("this is a custom error")

	if err != nil {
		//panic(err)
		// 处理知道的error
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}

		return
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
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
