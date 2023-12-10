package main

import (
	"2023/utils"
)

func alo(a []int) {
	a = append(a, 99)
}

type Pos struct {
	x int
	y int
}

func main() {
	s := "edvjchzeichzovf"
	utils.WriteFile("lol.txt", []byte(s))
}
