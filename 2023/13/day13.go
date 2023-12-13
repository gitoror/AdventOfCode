package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

func reflect(note []string, i int, mode_col bool, max_smudges int) int {
	if mode_col {
		j := 0
		smudges := 0
		for 0 <= i-j && i+1+j < len(note[0]) {
			for k := 0; k < len(note); k++ {
				if note[k][i-j] != note[k][i+1+j] {
					smudges++
					if smudges > max_smudges {
						return 0
					}
				}
			}
			j++
		}
		if smudges == max_smudges {
			return i + 1
		} else {
			return 0
		}
	} else {
		j := 0
		smudges := 0
		for 0 <= i-j && i+1+j < len(note) {
			for k := 0; k < len(note[0]); k++ {
				if note[i-j][k] != note[i+1+j][k] {
					smudges++
					if smudges > max_smudges {
						return 0
					}
				}
			}
			j++
		}
		if smudges == max_smudges {
			return (i + 1) * 100
		} else {
			return 0
		}
	}
}

func mirror(note []string, mode_col bool, max_smudges int) int {
	N := 0
	var size int
	if mode_col {
		size = len(note[0])
	} else {
		size = len(note)
	}
	for i := 0; i < size-1; i++ {
		N = utils.Max(N, reflect(note, i, mode_col, max_smudges))
	}
	return N
}

func solve(input []string, max_smudges int) int {
	sum := 0
	for _, block := range input {
		note := strings.Split(block, "\n")
		c := mirror(note, true, max_smudges) // col
		l := 0
		if c == 0 {
			l = mirror(note, false, max_smudges) // line
		}
		sum += c + l
	}
	return sum
}

func part1(input []string) int {
	return solve(input, 0)
}

func part2(input []string) int {
	return solve(input, 1)
}

func main() {
	file := utils.ReadFile("../inputs/13/input.txt")
	input := strings.Split(string(file), "\n\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
