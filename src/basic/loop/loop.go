package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// 转2进制
func conventToBin(n int) interface{} {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}
}
// 死循环
func forever() {
	for {
		println("forever")
	}
}

func main() {

	fmt.Println(
		conventToBin(5),
		conventToBin(11),
	)

	printFile("src/basic/loop/abc.text")
}
