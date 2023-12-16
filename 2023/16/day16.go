package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

type Pos struct{ x, y int }

func diaplay_energized(input []string, energized map[Pos]bool) {
	lenX, lenY := len(input[0]), len(input)
	for y := 0; y < lenY; y++ {
		line := ""
		for x := 0; x < lenX; x++ {
			if _, ok := energized[Pos{x, y}]; ok {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func beam(input []string, pos Pos, direction string, energized *map[Pos]bool) {
	beam_blocked := false
	lenX, lenY := len(input[0]), len(input)
	D := map[string]Pos{"U": {0, -1}, "R": {1, 0}, "D": {0, 1}, "L": {-1, 0}}
	for !beam_blocked {
		new_pos := Pos{pos.x + D[direction].x, pos.y + D[direction].y}
		if 0 <= new_pos.x && new_pos.x < lenX && 0 <= new_pos.y && new_pos.y < lenY {
			// fmt.Println(new_pos, string(input[new_pos.y][new_pos.x]), direction)
			pos = new_pos
			if input[pos.y][pos.x] == '|' {
				if _, ok := (*energized)[pos]; ok {
					beam_blocked = true
				} else {
					(*energized)[pos] = true
					if direction == "R" || direction == "L" {
						(*energized)[pos] = true
						beam(input, pos, "U", energized)
						beam(input, pos, "D", energized)
						beam_blocked = true
					}
				}
			} else if input[pos.y][pos.x] == '-' {
				if _, ok := (*energized)[pos]; ok {
					beam_blocked = true
				} else {
					(*energized)[pos] = true
					if direction == "U" || direction == "D" {
						beam(input, pos, "R", energized)
						beam(input, pos, "L", energized)
						beam_blocked = true
					}
				}
			} else if input[pos.y][pos.x] == '/' { // energized from what side ?
				// true->UL, false->DR
				energized_from_UL, already_seen := (*energized)[pos]
				if already_seen && ((energized_from_UL && (direction == "U" || direction == "L")) ||
					(!energized_from_UL && (direction == "D" || direction == "R"))) {
					beam_blocked = true
				} else {
					switch direction {
					case "D":
						(*energized)[pos] = false
						direction = "L"
					case "L":
						(*energized)[pos] = true
						direction = "D"
					case "U":
						(*energized)[pos] = true
						direction = "R"
					case "R":
						(*energized)[pos] = false
						direction = "U"
					}
				}
			} else if input[pos.y][pos.x] == '\\' { // true->UR, false->DL
				energized_from_UR, already_seen := (*energized)[pos]
				if already_seen && ((energized_from_UR && (direction == "U" || direction == "R")) ||
					(!energized_from_UR && (direction == "D" || direction == "L"))) {
					beam_blocked = true
				} else {
					switch direction {
					case "U":
						(*energized)[pos] = true
						direction = "L"
					case "R":
						(*energized)[pos] = true
						direction = "D"
					case "D":
						(*energized)[pos] = false
						direction = "R"
					case "L":
						(*energized)[pos] = false
						direction = "U"
					}
				}
			} else if input[pos.y][pos.x] == '.' {
				(*energized)[pos] = true
			}
		} else { // not in the layout anymore
			beam_blocked = true
		}
	}
}

func part2(input []string) int { // -|/\
	lenX, lenY := len(input[0]), len(input)
	energized := map[Pos]bool{}
	max := 0
	for x := 0; x < lenX; x++ {
		energized = map[Pos]bool{}
		beam(input, Pos{x, -1}, "D", &energized)
		v := len(energized)
		if v > max {
			max = v
		}
	}
	for x := 0; x < lenX; x++ {
		energized = map[Pos]bool{}
		beam(input, Pos{x, lenY}, "U", &energized)
		v := len(energized)
		if v > max {
			max = v
		}
	}
	for y := 0; y < lenY; y++ {
		energized = map[Pos]bool{}
		beam(input, Pos{-1, y}, "R", &energized)
		v := len(energized)
		if v > max {
			max = v
		}
	}
	for y := 0; y < lenY; y++ {
		energized = map[Pos]bool{}
		beam(input, Pos{lenX, y}, "L", &energized)
		v := len(energized)
		if v > max {
			max = v
		}
	}
	// diaplay_energized(input, energized)
	// fmt.Println(energized)
	return max
}

func part1(input []string) int {
	energized := map[Pos]bool{}
	beam(input, Pos{-1, 0}, "R", &energized)
	return len(energized)
}

func main() {
	file := utils.ReadFile("../inputs/16/in.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
	// part 2 6762 too low
}
