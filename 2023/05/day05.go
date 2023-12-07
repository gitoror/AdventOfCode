package main

import (
	"2023/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

func parse_mappings(blocks []string) [][][]int {
	mappings := [][][]int{}
	for i := 1; i < len(blocks); i++ {
		infos := strings.Split(
			strings.Split(blocks[i], ":\n")[1], "\n")
		mapping := [][]int{}
		for _, info := range infos {
			mapping = append(mapping,
				utils.SliceInt(
					strings.Split(info, " ")))
		}
		mappings = append(mappings, mapping)
	}
	return mappings
}

func calc_leaf(seed int, mappings [][][]int) int {
	if len(mappings) == 0 {
		return seed
	}
	for _, rule := range mappings[0] {
		dest, src, d := rule[0], rule[1], rule[2]
		d_seed := seed - src
		if 0 <= d_seed && d_seed <= d {
			return calc_leaf(dest+d_seed, mappings[1:])
		}
	}
	return calc_leaf(seed, mappings[1:])
}

func part1(input string) int {
	// Parse input
	blocks := strings.Split(input, "\n\n")
	seeds := utils.SliceInt(
		strings.Split(
			strings.Split(blocks[0], ": ")[1], " "))
	mappings := parse_mappings(blocks)
	// Compute leafs
	min := math.MaxInt
	for _, seed := range seeds {
		leaf := calc_leaf(seed, mappings)
		if leaf < min {
			min = leaf
		}
	}
	return min
}

func seeds_with_ranges(seeds []int) [][]int {
	// Store [min max]
	seeds_with_ranges := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seeds_with_ranges = append(seeds_with_ranges, []int{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	return seeds_with_ranges
}

func substract(a, b []int) [][]int { // a - b
	// May return 0, 1 or 2 intervals (2 when b includes a)
	r := [][]int{}
	before := []int{a[0], utils.Min(b[0]-1, a[1])}
	after := []int{utils.Max(b[1]+1, a[0]), a[1]}
	if before[1] > before[0] { // If non empty
		r = append(r, before)
	}
	if after[1] > after[0] {
		r = append(r, after)
	}
	return r
}

func calc_next_intervals(A [][]int, mapping [][]int) [][]int {
	r := [][]int{} // Returned intervals
	for _, m := range mapping {
		dest, src, size := m[0], m[1], m[2]
		end := src + size - 1
		shift := dest - src
		deviator := []int{src, end}
		for_later := [][]int{} // Interval that could be deviated by the next m of mapping
		for _, a := range A {
			// 1. Do intersection (deviated parts)
			b := utils.Intersect(a, deviator)
			// Deviate interval
			if len(b) > 0 {
				b = []int{b[0] + shift, b[1] + shift} // shift range
				r = append(r, b)                      // Directly added to return because wont be deviated anymore
			}
			// 2. Look at non deviated parts
			non_deviated_intervals := substract(a, deviator) // a - deviator
			// Add them to the set of intervals that may be deviated by another m of mapping
			for _, non_deviated_interval := range non_deviated_intervals {
				if len(non_deviated_interval) > 0 {
					for_later = append(for_later, non_deviated_interval)
				}
			}
		}
		A = for_later
	}
	return append(r, A...)
}

func part2(input string) int {
	// Parse input
	blocks := strings.Split(input, "\n\n")
	mappings := parse_mappings(blocks)
	seeds_infos := utils.SliceInt(strings.Split(strings.Split(blocks[0], ": ")[1], " "))
	// Answer intervals from where to find the min of starts
	A := seeds_with_ranges(seeds_infos)
	for _, mapping := range mappings {
		A = calc_next_intervals(A, mapping)
	}
	// Min algo
	min := math.MaxInt
	for _, a := range A {
		min = utils.Min(min, a[0])
	}
	return min
}

func main() {
	file := utils.ReadFile("../inputs/05/input.txt")
	input := string(file)
	// Part 1
	start := time.Now()
	part1 := part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))

}
