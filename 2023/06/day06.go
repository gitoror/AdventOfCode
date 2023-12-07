package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func part1(input string) int {
	infos := strings.Split(input, "\n")
	times := utils.SliceInt(strings.Fields(infos[0])[1:])
	distances := utils.SliceInt(strings.Fields(infos[1])[1:])
	product := 1
	for i := 0; i < len(times); i++ {
		n_beatable := 0
		for j := 1; j < times[i]; j++ {
			if j*(times[i]-j) > distances[i] {
				n_beatable++
			}
		}
		product *= n_beatable
	}
	return product
}

func part2(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	infos := strings.Split(input, "\n")
	time, _ := strconv.Atoi(infos[0][5:])
	distance, _ := strconv.Atoi(infos[1][9:])
	n_beatable := 0
	for j := 1; j < time; j++ {
		if j*(time-j) > distance {
			n_beatable++
		}
	}
	return n_beatable
}

func main() {
	file := utils.ReadFile("../inputs/06/input.txt")
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
