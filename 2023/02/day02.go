package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func get_cube_infos(cube string) (int, string) {
	cube_infos := strings.Split(cube, " ")
	nb, _ := strconv.Atoi(string(cube_infos[0]))
	color := cube_infos[1]
	return nb, color
}

func sum_possible_ids(input []string) int {
	sum := 0
	for _, game := range input {
		game_infos := strings.Split(game, ": ")
		game_id, _ := strconv.Atoi(string(strings.Split(game_infos[0], " ")[1]))
		cubes := strings.Split(strings.ReplaceAll(game_infos[1], ",", ";"), "; ")
		is_possible := true
		for _, cube := range cubes {
			nb, color := get_cube_infos(cube)
			if color == "red" && nb > 12 ||
				color == "green" && nb > 13 ||
				color == "blue" && nb > 14 {
				is_possible = false
				break
			}
		}
		if is_possible {
			sum += game_id
		}
	}
	return sum
}

func find_power(cubes []string) int {
	nb_max_red := -1
	nb_max_green := -1
	nb_max_blue := -1
	for _, cube := range cubes {
		nb, color := get_cube_infos(cube)
		if color == "red" && nb > nb_max_red {
			nb_max_red = nb
		}
		if color == "green" && nb > nb_max_green {
			nb_max_green = nb
		}
		if color == "blue" && nb > nb_max_blue {
			nb_max_blue = nb
		}
	}
	power := nb_max_red * nb_max_green * nb_max_blue
	return power
}

func power_sum(input []string) int {
	sum := 0
	for _, game := range input {
		game_infos := strings.Split(game, ": ")
		cubes := strings.Split(strings.ReplaceAll(game_infos[1], ",", ";"), "; ")
		power := find_power(cubes)
		sum += power
	}
	return sum
}

func main() {
	file := utils.ReadFile("../inputs/02/input.txt")
	input := strings.Split(string(file), "\n")
	// Part 1
	start := time.Now()
	part1 := sum_possible_ids(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := power_sum(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
