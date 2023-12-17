package main

import (
	"fmt"
)

func alo(a *[]int) {
	*a = append(*a, 99)
}

type Pos struct {
	x int
	y int
}

func main() {
	a := []string{"123"}
	x := int(a[0][0] - '0')
	fmt.Println(x)
}
