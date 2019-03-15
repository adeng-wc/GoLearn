package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel as first-calss citizen")
	channelDemo()
	bufferedChannel()
	channelClose()
}

func channelClose() {
	// 3个缓冲区
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	// 3个缓冲区
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

/*
	close 之后还能收到 空值
 */
func work(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}

	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func channelDemo() {

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	//var c chan int  // c = nil
	//c := make(chan int)

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
