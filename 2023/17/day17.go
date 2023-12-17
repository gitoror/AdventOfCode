package main

import (
	"2023/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

type Pos struct{ x, y int }

func min_pos(paths map[Pos]Info) Pos {
	min := math.MaxInt
	min_pos := Pos{-1, -1}
	for pos, info := range paths {
		if info.dst < min {
			min_pos = pos
			min = info.dst
		}
	}
	return min_pos
}

func get_neighbors(pos Pos, plan Plan) []Pos {
	neighbors := []Pos{}
	D := []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range D {
		neigh := Pos{pos.x + d.x, pos.y + d.y}
		if neigh.x >= 0 && neigh.x < plan.lenX && neigh.y >= 0 && neigh.y < plan.lenY {
			neighbors = append(neighbors, neigh)
		}
	}
	return neighbors
}

type Info struct {
	prev                                Pos
	dst, three_blocks_x, three_blocks_y int
}

func get_shortest_path(paths map[Pos]Info, src, dest Pos) []Pos {
	path := []Pos{}
	pos := dest
	for pos != src {
		path = append(path, pos)
		pos = paths[pos].prev
		fmt.Println(pos)
	}
	return path
}

func containsPos(path []Pos, pos Pos) bool {
	for _, p := range path {
		if p == pos {
			return true
		}
	}
	return false
}

func display_map_with_path(plan Plan, path []Pos) {
	for y := 0; y < plan.lenY; y++ {
		for x := 0; x < plan.lenX; x++ {
			pos := Pos{x, y}
			if containsPos(path, pos) {
				fmt.Print("X")
			} else {
				fmt.Print(plan.M[Pos{x, y}])
			}
		}
		fmt.Println()
	}
}

func dijkstra(plan Plan, src, dest Pos) int {
	M := plan.M
	// paths := map[Pos]int{src: 0}
	paths := map[Pos]Info{src: {Pos{0, 0}, 0, 0, 0}}
	already_explored := map[Pos]bool{}
	found := false
	min_dst := -1
	p := map[Pos]Info{}
	for !found {
		pos := min_pos(paths)
		if pos == dest {
			paths[pos] = Info{paths[pos].prev, paths[pos].dst, paths[pos].three_blocks_x, paths[pos].three_blocks_y}
			min_dst = paths[pos].dst
			found = true
			break
		}
		neighbors := get_neighbors(pos, plan) // those inside map bounds
		for _, neigh := range neighbors {
			if _, ok := already_explored[neigh]; !ok { // if not already explored
				same_direction_x, same_direction_y := (pos.x == neigh.x), (pos.y == neigh.y)
				new_three_blocks_x, new_three_blocks_y, new_dst := -1, -1, -1
				if same_direction_x {
					new_three_blocks_x = paths[pos].three_blocks_x + 1
				} else {
					new_three_blocks_x = 0
				}
				if same_direction_y {
					new_three_blocks_y = paths[pos].three_blocks_y + 1
				} else {
					new_three_blocks_y = 0
				}
				if new_three_blocks_x > 3 || new_three_blocks_y > 3 {
					new_dst = math.MaxInt
				} else {
					new_dst = paths[pos].dst + M[neigh]
				}
				if _, ok := paths[neigh]; ok { // if there is already a value for that neigh
					if new_dst < paths[neigh].dst {
						paths[neigh] = Info{pos, new_dst, new_three_blocks_x, new_three_blocks_y} // then replace it only if smaller
					}
				} else {
					paths[neigh] = Info{pos, new_dst, new_three_blocks_x, new_three_blocks_y}
				}
			}
		}
		p[pos] = paths[pos] //debug

		already_explored[pos] = true
		delete(paths, pos)
		fmt.Println(pos, paths)
	}
	fmt.Println(p)
	p[dest] = paths[dest]
	pa := get_shortest_path(p, src, dest)
	fmt.Println(pa)
	display_map_with_path(plan, pa)
	return min_dst
}

type Plan struct {
	M          map[Pos]int
	lenX, lenY int
}

func get_map(input []string) Plan {
	M := map[Pos]int{}
	lenX, lenY := len(input[0]), len(input)
	for x := 0; x < len(input[0]); x++ {
		for y := 0; y < len(input); y++ {
			M[Pos{x, y}] = int(input[y][x] - '0')
		}
	}
	return Plan{M, lenX, lenY}
}

func part1(input []string) int {
	plan := get_map(input)
	src, dest := Pos{0, 0}, Pos{len(input[0]) - 1, len(input) - 1}
	return dijkstra(plan, src, dest)
}

func main() {
	file := utils.ReadFile("../inputs/17/ex2.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// start, part2 := time.Now(), part2(input)
	// fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
