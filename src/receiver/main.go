package main

import (
	"fmt"
	real2 "GoLearn/src/receiver/real"
	"GoLearn/src/receiver/mock"
)

type Receiver interface {
	Get(url string) string
}

func download(r Receiver) string {
	return r.Get("http://www.baidu.com")
}

func main() {

	var r Receiver = mock.Receiver{Contexts: "mock Receiver"}
	var r2 Receiver = real2.Receiver{}

	fmt.Println(download(r))
	fmt.Println(download(r2))
}