package main

import (
	"fmt"
)

func alo(a []int) {
	a = append(a, 99)
}

func main() {
	a := []int{}
	alo(a)
	fmt.Println(a)
}
