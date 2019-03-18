package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {

	fmt.Println("atomicInt increment")

	// 利用匿名函数来实现 代码块 锁定
	func() {
		a.lock.Lock()
		// 确保在函数结束时发生
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {

	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Second)
	fmt.Println(a.get())
}
