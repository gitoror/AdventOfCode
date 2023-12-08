package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

type Next map[string]string

func part1(input []string) int {
	instructions := input[0]
	network_infos := strings.Split(input[1], "\n")
	network := map[string]Next{}
	for _, n := range network_infos {
		network[n[0:3]] = map[string]string{"L": n[7:10], "R": n[12:15]}
	}
	next := "AAA"
	pos := 0
	step := 0
	for next != "ZZZ" {
		next = network[next][string(instructions[pos])]
		pos = (pos + 1) % len(instructions)
		step++
	}
	return step
}

func part2(input []string) int {
	instructions := input[0]
	network_infos := strings.Split(input[1], "\n")
	network := map[string]Next{}
	for _, n := range network_infos {
		network[n[0:3]] = map[string]string{"L": n[7:10], "R": n[12:15]}
	}
	// Extract init next
	var nexts []string
	for k := range network {
		if string(k[2]) == "A" {
			nexts = append(nexts, k)

		}
	}
	steps := []int{}
	for _, next := range nexts {
		pos := 0
		step := 0
		for string(next[2]) != "Z" {
			next = network[next][string(instructions[pos])]
			pos = (pos + 1) % len(instructions)
			step++
		}
		steps = append(steps, step)
	}
	return utils.LCM(steps[0], steps[1], steps[2:]...)
}

func main() {
	file := utils.ReadFile("../inputs/08/input.txt")
	input := strings.Split(string(file), "\n\n")
	// Part 1
	start := time.Now()
	part1 := part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
