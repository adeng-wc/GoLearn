package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(id int, c chan int) {
	for n := range c {
		// 演示 生产者 > 消费者
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)

			out <- i

			//fmt.Printf("generator send %d\n", i)
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWork(0)

	var values []int

	after := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		// activeWorker, c1, c2 都是 channel
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		// 通过 通信 来共享数据。 CSP 模型
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
			// activeWorker 是 nil 的情况下不会循环到
		case activeWorker <- activeValue:
			values = values[1:]
			// 循环的时候超过 800 毫秒 就会提示超时
		case <-time.After(200 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-after:
			fmt.Println("byb")
			return
		}
	}
}
