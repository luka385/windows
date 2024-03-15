package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var s string
	s = 1
	x, err := ioutil.ReadFile("archivvo.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(x))
}
