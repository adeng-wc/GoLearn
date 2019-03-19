package main

import (
	"fmt"
	"regexp"
)

const text = `My email is wengcheng.no1@gmail.com
email is abf1@gmail.com
email is kkk@gmail.com
email is ddd@gmail.com

`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9.]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9.]+)`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)

	submatch := re.FindAllStringSubmatch(text, -1)
	fmt.Println(submatch)

	for _, m := range submatch {
		fmt.Println(m)
	}
}
