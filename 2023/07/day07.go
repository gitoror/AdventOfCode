package main

import (
	"2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Hand struct {
	Cards string
	Bid   int
}

func find_category(cards string) int {
	// cat 1 is five of a kind
	labels_count := map[rune]int{'A': 0, 'K': 0, 'Q': 0, 'J': 0, 'T': 0, '9': 0, '8': 0, '7': 0, '6': 0, '5': 0, '4': 0, '3': 0, '2': 0} // same order
	for _, card := range cards {
		labels_count[card]++
	}
	n_uplets := [6]int{} // number of "nothing", singletons, pairs, triplets, 4-uplets, 5-uplets
	for _, v := range labels_count {
		n_uplets[v]++
	}
	if n_uplets[5] == 1 {
		return 1 // five of a kind
	}
	if n_uplets[4] == 1 {
		return 2 // four of a kind
	}
	if n_uplets[3] == 1 {
		if n_uplets[2] == 1 {
			return 3 // full house
		}
		return 4 // three of a kind
	}
	if n_uplets[2] == 2 {
		return 5 //two pairs
	}
	if n_uplets[2] == 1 {
		return 6 // one pair
	}
	return 7 // high card
}

func find_category2(cards string) int {
	labels_count := map[rune]int{'A': 0, 'K': 0, 'Q': 0, 'J': 0, 'T': 0, '9': 0, '8': 0, '7': 0, '6': 0, '5': 0, '4': 0, '3': 0, '2': 0} // same order
	for _, card := range cards {
		labels_count[card]++
	}
	if labels_count['J'] == 0 {
		return find_category(cards)
	}
	prev_cat := find_category(cards)
	switch prev_cat {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 1
	case 4:
		return 2
	case 5:
		if labels_count['J'] == 1 {
			return 3
		}
		return 2 // labels_count['J'] == 2
	case 6:
		return 4
	case 7:
		return 6
	default:
		fmt.Println("Problem: previous category wrong")
		return -1 // impossible
	}
}

func greater_than(c1 string, c2 string, joker bool) bool { // c1 stronger than c2 ?
	var cat1, cat2 int
	var labels map[rune]int
	if joker {
		cat1, cat2 = find_category2(c1), find_category2(c2)
		labels = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 0, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}
	} else {
		cat1, cat2 = find_category(c1), find_category(c2)
		labels = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}
	}
	if cat1 < cat2 {
		return true
	}
	if cat1 == cat2 {
		for i := 0; i < 5; i++ {
			if labels[rune(c1[i])] > labels[rune(c2[i])] {
				return true
			}
			if labels[rune(c1[i])] < labels[rune(c2[i])] {
				return false
			}
		}
		fmt.Println("Problem: 2 hands are equals, cannot compare them")
	}
	return false
}

// Sort interface ==> need to implement some methods
// Ref: https://pkg.go.dev/sort
type ByCards []Hand

func (h ByCards) Len() int           { return len(h) }
func (h ByCards) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByCards) Less(i, j int) bool { return greater_than(h[j].Cards, h[i].Cards, false) }

func part1(input []string) int {
	hands := []Hand{}
	for _, line := range input {
		bid, _ := strconv.Atoi(line[6:])
		hands = append(hands, Hand{Cards: line[0:5], Bid: bid})
	}
	sort.Sort(ByCards(hands))
	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += (i + 1) * hands[i].Bid
	}
	return sum
}

var wantMe []string

func part2(input []string) int {
	hands := []Hand{}
	for _, line := range input {
		bid, _ := strconv.Atoi(line[6:])
		hands = append(hands, Hand{Cards: line[0:5], Bid: bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		return greater_than(hands[j].Cards, hands[i].Cards, true)
	})

	for i := 0; i < len(hands); i++ {
		wantMe = append(wantMe, hands[i].Cards)
	}

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += (i + 1) * hands[i].Bid
	}
	return sum
}

func main() {
	file := utils.ReadFile("../inputs/07/input.txt")
	input := strings.Split(string(file), "\n")
	// Part 1
	start := time.Now()
	part1 := part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))

	fmt.Println(find_category2("QJJQ2"))
	// Part 2
	start = time.Now()
	part2 := part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))

}
