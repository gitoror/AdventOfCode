package main

import (
	"2023/utils"
	"fmt"
	"math"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func part1(input []string) int {
	sum := 0
	for _, card := range input {
		n_matches := 0
		card = strings.ReplaceAll(card, "  ", " ")
		card_infos := strings.Split(strings.Split(card, ": ")[1], " | ")
		winning_numbers := utils.SliceInt(strings.Split(card_infos[0], " "))
		my_numbers := utils.SliceInt(strings.Split(card_infos[1], " "))
		for _, n := range my_numbers {
			if slices.Contains(winning_numbers, n) {
				n_matches++
			}
		}
		sum += int(math.Pow(2, float64(n_matches-1)))
	}
	return sum
}

func part2(input []string) int {
	n_matches := map[int]int{}
	for id, card := range input {
		matches := 0
		card = strings.ReplaceAll(card, "  ", " ")
		card_infos := strings.Split(strings.Split(card, ": ")[1], " | ")
		winning_numbers := utils.SliceInt(strings.Split(card_infos[0], " "))
		my_numbers := utils.SliceInt(strings.Split(card_infos[1], " "))
		for _, n := range my_numbers {
			if slices.Contains(winning_numbers, n) {
				matches++
			}
		}
		n_matches[id] = matches
	}
	//
	won_cards := map[int]int{}
	for id := 0; id < len(input); id++ {
		won_cards[id] = 1
	}
	//
	for id := 1; id < len(input); id++ {
		for k := 0; k < id; k++ {
			if n_matches[k] >= id-k {
				won_cards[id] += won_cards[k]
			}
		}
	}
	//
	n_cards := 0
	for _, n := range won_cards {
		n_cards += n
	}
	return n_cards
}

func main() {
	file := utils.ReadFile("../inputs/04/input.txt")
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
