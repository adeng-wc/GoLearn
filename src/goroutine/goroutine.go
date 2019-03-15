package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int

	for i := 0; i < 1000; i++ {
		// 并发执行
		go func(ii int) { // WARNING: DATA RACE
			for {
				// IO 操作,有协程之间的切换
				fmt.Printf("Hello from goroutine %d\n", i)
				//a[ii]++
				// 释放控制器
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
