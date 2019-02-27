package main

import (
	"io/ioutil"
	"fmt"
)

func grade(source int) string {
	g := ""
	switch {
	case source < 60:
		g = "F"
	case source < 80:
		g = "C"
	case source < 90:
		g = "B"
	default:
		g = "A"
	}
	return g
}

func main() {

	const filename = "abc.text"
	//contexts, errors := ioutil.ReadFile(filename)

	if contexts, errors := ioutil.ReadFile(filename); errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Printf("%s\n", contexts)
	}

	fmt.Println(grade(1), grade(69))
}
