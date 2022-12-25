package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("inputs/day22/ex.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n\n")
	fmt.Println(input)
}
