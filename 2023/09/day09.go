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
	last_values := []int{h[len(h)-1]}
	for !all_zeros(h) {
		h = diff(h)
		last_values = append(last_values, h[len(h)-1])
	}
	return utils.SliceSum(last_values)
}

func part1(input []string) int {
	sum := 0
	for _, history := range input {
		sum += predict(utils.SliceInt(strings.Split(history, " ")))
	}
	return sum
}

func predict_backwards(h []int) int {
	first_values := []int{h[0]}
	for !all_zeros(h) {
		h = diff(h)
		first_values = append(first_values, h[0])
	}
	a := first_values[0]
	for k, v := range first_values[1:] {
		if k%2 == 0 {
			a -= v
		} else {
			a += v
		}
	}
	return a
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
	// Part 1
	start := time.Now()
	part1 := part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
