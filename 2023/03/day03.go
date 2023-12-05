package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Pos struct{ left, right, y int }

func is_symbol(s string) bool {
	if strings.Contains("0123456789", s) || s == "." {
		return false
	}
	return true
}

func numbers_pos(input []string) map[Pos]int {
	numbers := map[Pos]int{}
	for y, line := range input {
		number := ""
		pos := -1
		for i, x := range line {
			if strings.Contains("0123456789", string(x)) {
				if number == "" {
					pos = i
				}
				number += string(x)
			} else {
				if number != "" {
					n, _ := strconv.Atoi(number)
					numbers[Pos{left: pos, right: pos + len(number) - 1, y: y}] = n
					number = ""
				}
			}
		}
		// In case of we reached the end of the line with a number
		if number != "" {
			n, _ := strconv.Atoi(number)
			numbers[Pos{left: pos, right: pos + len(number) - 1, y: y}] = n
			number = ""
		}
	}
	return numbers
}

func is_adjacent(n_pos Pos, input []string) bool {
	d := []int{-1, 0, 1}
	y := n_pos.y
	for x := n_pos.left; x <= n_pos.right; x++ {
		for _, dx := range d {
			for _, dy := range d {
				if 0 <= y+dy && y+dy < len(input) &&
					0 <= x+dx && x+dx < len(input[0]) &&
					is_symbol(string(input[y+dy][x+dx])) {
					return true
				}
			}
		}
	}
	return false
}

func part1(input []string) int {
	sum := 0
	numbers := numbers_pos(input)
	for number, n := range numbers {
		if is_adjacent(number, input) {
			sum += n
		}
	}
	return sum
}

func get_left(input []string, x int, y int) string {
	number := ""
	for i := x; i >= 0 &&
		strings.Contains("0123456789", string(input[y][i])); i-- {
		number = string(input[y][i]) + number
	}
	return number
}

func get_right(input []string, x int, y int) string {
	number := ""
	for i := x; i < len(input[0]) &&
		strings.Contains("0123456789", string(input[y][i])); i++ {
		number = number + string(input[y][i])
	}
	return number
}

func find_adj_number(xg int, yg int, input []string) []int {
	adj_numbers := []int{}
	for _, y := range []int{yg - 1, yg, yg + 1} {
		if 0 <= y && y < len(input) {
			if string(input[y][xg]) == "." || y == yg {
				if string(input[y][xg-1]) != "." {
					number := get_left(input, xg-1, y)
					n, _ := strconv.Atoi(number)
					adj_numbers = append(adj_numbers, n)
				}
				if string(input[y][xg+1]) != "." {
					number := get_right(input, xg+1, y)
					n, _ := strconv.Atoi(number)
					adj_numbers = append(adj_numbers, n)
				}
			} else {
				number := get_left(input, xg-1, y) +
					string(input[y][xg]) +
					get_right(input, xg+1, y)
				n, _ := strconv.Atoi(number)
				adj_numbers = append(adj_numbers, n)
			}
		}
	}
	return adj_numbers
}

type XY struct{ x, y int }

func find_gears(input []string) map[XY]int {
	gears := map[XY]int{}
	for y, line := range input {
		for x, char := range line {
			s := string(char)
			if s == "*" {
				adj_numbers := find_adj_number(x, y, input)
				if len(adj_numbers) == 2 {
					gears[XY{x: x, y: y}] = adj_numbers[0] * adj_numbers[1] // gear ratio
				}
			}
		}
	}
	return gears
}

func sum_gear_ratios(input []string) int {
	sum := 0
	gears := find_gears(input)
	for _, gear_ratio := range gears {
		sum += gear_ratio
	}
	return sum
}
func main() {
	file := utils.ReadFile("../inputs/03/input.txt")
	input := strings.Split(string(file), "\n")

	// Part 1
	start := time.Now()
	part1 := part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := sum_gear_ratios(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
