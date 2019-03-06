package main

import (
	"fmt"
	"GoLearn/src/receiver/mock"
	"GoLearn/src/receiver/real"
	"time"
)

type Receiver interface {
	Get(url string) string
}

func download(r Receiver) string {
	return r.Get("http://www.baidu.com")
}

func main() {

	var r Receiver
	r = mock.Receiver{Contexts: "this is mock Receiver"}
	inspect(r)

	r = &real.Receiver{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// type assertion
	if realReceiver, ok := r.(*real.Receiver); ok {
		fmt.Println(realReceiver.TimeOut)
	} else {
		fmt.Println("not a real Receiver")
	}
	if mockReceiver, ok := r.(*mock.Receiver); ok {
		fmt.Println(mockReceiver.Contexts)
	} else {
		fmt.Println("not a mock Receiver")
	}

	//fmt.Println(download(r))
	//fmt.Println(download(r2))

}

func inspect(r Receiver) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Receiver:
		fmt.Println(v.Contexts)
	case *real.Receiver:
		fmt.Println(v.UserAgent)
	}
}
