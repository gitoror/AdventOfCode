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

func seeds_with_ranges1(seeds []int) []int {
	seeds_with_ranges := []int{}
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		seed_range := seeds[i+1]
		for k := 0; k < seed_range; k++ {
			seeds_with_ranges = append(seeds_with_ranges, seed+k)
		}
	}
	return seeds_with_ranges
} // Trop long, faire intersection des intervalles de seed

func merge_ranges(ranges [][]int) [][]int {
	new_ranges := [][]int{}
	// Merge
	for i := 0; i < len(ranges); i++ {
		r1 := ranges[i]
		merged := false // merged at least once
		for j := i + 1; j < len(ranges); j++ {
			r2 := ranges[j]
			if r1[1] < r2[0] || r2[1] < r1[0] {
				break
			}
			if r1[0] <= r2[0] && r2[1] <= r1[1] {
				merged = true
				new_ranges = append(new_ranges, r1)
			}
			if r2[0] <= r1[0] && r1[1] <= r2[1] {
				merged = true
				new_ranges = append(new_ranges, r2)
			}
			if r1[0] < r2[0] && r1[1] <= r2[1] {
				merged = true
				new_ranges = append(new_ranges, []int{r1[0], r2[1]})
			}
			if r2[0] <= r1[0] && r2[1] < r1[1] {
				merged = true
				new_ranges = append(new_ranges, []int{r2[0], r1[1]})
			}
		}
		if !merged {
			new_ranges = append(new_ranges, []int{r1[0], r1[1]})
		}
	}
	return new_ranges
} // Still too long

func seeds_with_ranges(seeds []int) [][]int {
	// Store [min max]
	seeds_with_ranges := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seeds_with_ranges = append(seeds_with_ranges, []int{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	seeds_with_ranges = merge_ranges(seeds_with_ranges)
	return seeds_with_ranges
}

func part2(input string) int {
	// Parse input
	blocks := strings.Split(input, "\n\n")
	seeds := utils.SliceInt(
		strings.Split(
			strings.Split(blocks[0], ": ")[1], " "))
	mappings := parse_mappings(blocks)
	seeds_ranges := seeds_with_ranges(seeds)
	print("ok")

	// Compute leafs
	min := math.MaxInt
	for _, r := range seeds_ranges {
		for seed := r[0]; seed <= r[1]; seed++ {
			leaf := calc_leaf(seed, mappings)
			if leaf < min {
				min = leaf
			}
		}
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
