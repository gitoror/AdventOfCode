package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

func get_empty_col(input []string) []int {
	r := []int{}
	for x := 0; x < len(input[0]); x++ {
		add := true
		for y := 0; y < len(input); y++ {
			if input[y][x] == '#' {
				add = false
				break
			}
		}
		if add {
			r = append(r, x)
		}
	}
	return r
}

func get_empty_lines(input []string) []int {
	r := []int{}
	for y := 0; y < len(input); y++ {
		add := true
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == '#' {
				add = false
				break
			}
		}
		if add {
			r = append(r, y)
		}
	}
	return r
}

type Pos struct{ x, y int }

func get_galaxies(input []string) []Pos {
	r := []Pos{}
	for x := 0; x < len(input[0]); x++ {
		for y := 0; y < len(input); y++ {
			if input[y][x] == '#' {
				r = append(r, Pos{x, y})
			}
		}
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(input []string, d int) int {
	empty_col := get_empty_col(input)
	empty_lines := get_empty_lines(input)
	galaxies := get_galaxies(input)
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		g1 := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			sum += abs(g1.x-g2.x) + abs(g1.y-g2.y)
			for _, col := range empty_col {
				if utils.Min(g1.x, g2.x) < col && col < utils.Max(g1.x, g2.x) {
					sum += d - 1
				}
			}
			for _, l := range empty_lines {
				if utils.Min(g1.y, g2.y) < l && l < utils.Max(g1.y, g2.y) {
					sum += d - 1
				}
			}
		}
	}
	return sum
}

func part1(input []string) int {
	return solve(input, 2)
}

func part2(input []string) int {
	return solve(input, 1000000)
}

func main() {
	file := utils.ReadFile("../inputs/11/input.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))

}
