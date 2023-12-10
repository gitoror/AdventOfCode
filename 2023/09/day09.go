package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

func all_zeros(t []int) bool {
	for _, v := range t {
		if v != 0 {
			return false
		}
	}
	return true
}

func diff(t []int) []int {
	r := []int{}
	for i := 0; i < len(t)-1; i++ {
		r = append(r, t[i+1]-t[i])
	}
	return r
}

func predict(h []int) int {
	p := h[len(h)-1]
	for !all_zeros(h) {
		h = diff(h)
		p += h[len(h)-1]
	}
	return p
}

func part1(input []string) int {
	sum := 0
	for _, history := range input {
		sum += predict(utils.SliceInt(strings.Split(history, " ")))
	}
	return sum
}

func predict_backwards(h []int) int {
	p, k := h[0], 1
	for !all_zeros(h) {
		h = diff(h)
		if k%2 == 1 {
			p -= h[0]
		} else {
			p += h[0]
		}
		k++
	}
	return p
}

func part2(input []string) int {
	sum := 0
	for _, history := range input {
		sum += predict_backwards(utils.SliceInt(strings.Split(history, " ")))
	}
	return sum
}

func main() {
	file := utils.ReadFile("../inputs/09/input.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
