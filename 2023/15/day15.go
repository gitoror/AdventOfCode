package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

func hash(code string) int {
	h := 0
	for _, s := range code {
		h += int(s)
		h *= 17
		h = h % 256
	}
	// fmt.Println(code, h)
	return h
}

func part1(input []string) int {
	sum := 0
	for _, code := range input {
		sum += hash(code)
	}
	return sum
}

func main() {
	file := utils.ReadFile("../inputs/15/in.txt")
	input := strings.Split(string(file), ",")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// start, part2 := time.Now(), part2(input)
	// fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))

}
