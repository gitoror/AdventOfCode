package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func is_digit(x string) bool {
	digits := "0123456789"
	return strings.Contains(digits, x)
}

func is_spelled_digit(i int, line string, reverse bool) (bool, int) {
	spelled_digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for digit, spelled_digit := range spelled_digits {
		if !reverse {
			if line[i] == spelled_digit[0] {
				if i+len(spelled_digit) < len(line) &&
					line[i:i+len(spelled_digit)] == spelled_digit {
					return true, digit
				}
			}
		} else {
			if line[i] == spelled_digit[len(spelled_digit)-1] {
				if i-len(spelled_digit)+1 >= 0 &&
					line[i-len(spelled_digit)+1:i+1] == spelled_digit {
					return true, digit
				}
			}
		}

	}
	return false, -1
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		var (
			x, y string
		)
		// Find the first number from left side
		for i := range line {
			x = string(line[i])
			if is_digit(x) {
				break
			}
		}
		// Find the first number from right side
		for i := len(line) - 1; i >= 0; i-- {
			y = string(line[i])
			if is_digit(y) {
				break
			}
		}
		dsum, _ := strconv.Atoi(x + y)
		sum += dsum
	}
	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		var (
			x, y string
		)
		// Find the first number from left side
		for i := range line {
			x = string(line[i])
			if is_digit(x) {
				break
			}
			is_spelled_digit, digit := is_spelled_digit(i, line, false)
			if is_spelled_digit {
				x = strconv.Itoa(digit)
				break
			}
		}
		// Find the first number from right side
		for i := len(line) - 1; i >= 0; i-- {
			y = string(line[i])
			if is_digit(y) {
				break
			}
			is_spelled_digit, digit := is_spelled_digit(i, line, true)
			if is_spelled_digit {
				y = strconv.Itoa(digit)
				break
			}
		}
		dsum, _ := strconv.Atoi(x + y)
		sum += dsum
	}
	return sum
}

func main() {
	file := utils.ReadFile("../inputs/01/input.txt")
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
