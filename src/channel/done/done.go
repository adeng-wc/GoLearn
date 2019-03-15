package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, done func()) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 往 done 发
		//go func() { done <- true }()
		done()
	}
}

type worker struct {
	in chan int
	//done chan bool

	//wg *sync.WaitGroup
	done func()
}

func createWork(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: make(chan bool),
		done: func() {
			wg.Done()
		},
	}

	//go doWork(id, w.in, w.done)
	go doWork(id, w.in, w.done)
	return w
}

func channelDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i, &wg)
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// 或者 <-worker.done ， doWork 中就不需要再创建 goroute

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	//for _, worker := range workers {
	//	// done 收
	//	<-worker.done
	//	<-worker.done
	//}

	wg.Wait()
}

func main() {
	channelDemo()
}
