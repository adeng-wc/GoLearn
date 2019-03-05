package main

import "fmt"

// 循环队列
type MyCircularQueue struct {
	size   int
	arrays []int
	head   int // 队列头指针
	tail   int // 队列尾指针
}

/** Initialize your data structure here. Set the size of the queue to be size. */
func Constructor(k int) MyCircularQueue {
	var result = MyCircularQueue{size: k, head: -1, tail: -1}
	result.arrays = make([]int, k)
	return result
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		fmt.Println("array is full")
		return false
	}

	if this.head == this.tail && this.tail == -1 {
		this.head = 0
		this.tail = 0
		this.arrays[0] = value
		return true
	} else {
		this.tail = (this.tail + 1) % this.size
		this.arrays[this.tail] = value
	}

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		fmt.Println("array is empty")
		return false
	}

	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	} else {
		this.head = (this.head + 1) % this.size
	}

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head == -1{
		return -1
	}

	return this.arrays[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.tail == -1{
		return -1
	}

	return this.arrays[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == this.tail && -1 == this.tail {
		return true
	} else {
		return false
	}
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if index := (this.tail + 1) % len(this.arrays); index == this.head {
		return true
	} else {
		return false
	}
}

func main() {

	//circularQueue := Constructor(3) //将大小设置为3
	//fmt.Println(circularQueue.EnQueue(1))                     //返回true
	//fmt.Println(circularQueue.EnQueue(2))                       //返回true
	//fmt.Println(circularQueue.EnQueue(3))                       //返回true
	//fmt.Println(circularQueue.EnQueue(4))                       //返回false，队列已满
	//fmt.Println(circularQueue.Rear())                           //返回3
	//fmt.Println(circularQueue.IsFull())                         //返回true
	//fmt.Println(circularQueue.DeQueue())                        //返回true
	//fmt.Println(circularQueue.EnQueue(4))                       //返回true
	//fmt.Println(circularQueue.Rear())                           //返回4

	circularQueue := Constructor(6) //将大小设置为3
	fmt.Println(circularQueue.EnQueue(6))
	fmt.Println(circularQueue.Rear())
	fmt.Println(circularQueue.Rear())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.EnQueue(5))
	fmt.Println(circularQueue.Rear())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.Front())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.DeQueue())
}
