package main

import (
	"fmt"
)

func main() {
	tryRecover()
}
func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occured:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do:%v \n", r))
		}
	}()
	//panic(errors.New("this is an error"))

	b := 0
	a := 5 / b
	fmt.Println(a)

	//panic(123)

}
