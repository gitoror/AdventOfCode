package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

type Pos struct{ x, y int }

func connected(d Pos, pipe string, next_pipe string) bool { // -|7FLJ
	if d.y == -1 && strings.Contains("S|LJ", pipe) && strings.Contains("S|7F", next_pipe) ||
		d.y == 1 && strings.Contains("S|7F", pipe) && strings.Contains("S|LJ", next_pipe) ||
		d.x == -1 && strings.Contains("S-7J", pipe) && strings.Contains("S-LF", next_pipe) ||
		d.x == 1 && strings.Contains("S-LF", pipe) && strings.Contains("S-7J", next_pipe) {
		return true
	}
	return false
}

func in_map(pos Pos, X int, Y int) bool {
	return 0 <= pos.x && pos.x < X && 0 <= pos.y && pos.y < Y
}

func retrieve_start(input []string) Pos {
	for y, line := range input {
		for x, r := range line {
			if r == 'S' {
				return Pos{x, y}
			}
		}
	}
	return Pos{-1, -1}
}

func retrieve_loop(input []string) (map[Pos]string, []Pos) {
	start := retrieve_start(input)
	loop := map[Pos]string{start: "S"}
	keys := []Pos{start}
	// follow path
	path_len := 1
	pos := start
	pipe := string(input[pos.y][pos.x])
	prev_pos := pos // init
	len_X, len_Y := len(input[0]), len(input)
	// look around pos while pos != S again
	// if connected to pos and not previous pos
	// pos = this new pos; len ++
	// if pos == S: break
	D := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	back_to_start := false
	for !back_to_start {
		for _, d := range D {
			next := Pos{x: pos.x + d.x, y: pos.y + d.y}
			if next != prev_pos &&
				in_map(next, len_X, len_Y) &&
				connected(d, pipe, string(input[next.y][next.x])) {
				if string(input[next.y][next.x]) == "S" {
					back_to_start = true
				} else {
					prev_pos = pos
					pos = next
					pipe = string(input[pos.y][pos.x])
					path_len++
					loop[pos] = pipe
					keys = append(keys, pos)
				}
				break
			}
		}
	}
	return loop, keys
}

func retrieve_maze(input []string) map[Pos]string {
	maze := map[Pos]string{}
	for y, line := range input {
		for x, s := range line {
			maze[Pos{x, y}] = string(s)
		}
	}
	return maze
}

func part1(input []string) int {
	loop, _ := retrieve_loop(input)
	return len(loop) / 2
}

func print_maze(maze map[Pos]string, input []string) {
	len_X, len_Y := len(input[0]), len(input)
	for y := 0; y < len_Y; y++ {
		line := ""
		for x := 0; x < len_X; x++ {
			line += maze[Pos{x, y}]
		}
		fmt.Println(line)
	}
}

func write_maze(maze map[Pos]string, loop map[Pos]string, input []string) {
	// for pos := range loop {
	// 	if strings.Contains("JF7L", maze[pos]) {
	// 		maze[pos] = " "
	// 	}
	// }
	// for pos, s := range maze {
	// 	if _, ok := loop[pos]; !ok && !strings.Contains("01", s) {
	// 		maze[pos] = "X"
	// 	}
	// }
	for pos, s := range maze { // only draw loop
		if _, ok := loop[pos]; !ok {
			if strings.Contains("01", s) {
				maze[pos] = " "
			} else {
				maze[pos] = "X"
			}
		}
	}
	len_X, len_Y := len(input[0]), len(input)
	data_string := ""
	for y := 0; y < len_Y; y++ {
		line := ""
		for x := 0; x < len_X; x++ {
			line += maze[Pos{x, y}]
		}
		data_string += line + "\n"
	}
	utils.WriteFile("lol.txt", []byte(data_string))
}

func add_oz(pos Pos, s string, t *[]Pos, loop map[Pos]string, maze map[Pos]string, len_X, len_Y int) {
	if _, ok := loop[pos]; !ok && !strings.Contains("10", maze[pos]) && in_map(pos, len_X, len_Y) {
		*t = append(*t, pos)
		maze[pos] = "0"
	}
}

func part2(input []string) []int { // -|7FLJ
	len_X, len_Y := len(input[0]), len(input)
	maze := retrieve_maze(input)
	loop, keys := retrieve_loop(input)
	ones := []Pos{} // count on the right
	zeros := []Pos{}
	start := retrieve_start(input)
	prev_pos := start
	for _, pos := range keys {
		right := Pos{pos.x + 1, pos.y}
		left := Pos{pos.x - 1, pos.y}
		down := Pos{pos.x, pos.y + 1}
		up := Pos{pos.x, pos.y - 1}
		if pos.y < prev_pos.y {
			add_oz(right, "0", &zeros, loop, maze, len_X, len_Y)
			add_oz(left, "1", &ones, loop, maze, len_X, len_Y)
			// if _, ok := loop[right]; !ok && !strings.Contains("10", maze[right]) && right.x < len_X {
			// 	zeros = append(zeros, right)
			// 	maze[right] = "0"
			// }
			// if _, ok := loop[left]; !ok && !strings.Contains("10", maze[left]) && 0 <= left.x {
			// 	ones = append(ones, left)
			// 	maze[left] = "1"
			// }
		}
		if pos.y > prev_pos.y {
			add_oz(right, "1", &ones, loop, maze, len_X, len_Y)
			add_oz(left, "0", &zeros, loop, maze, len_X, len_Y)
			// if _, ok := loop[right]; !ok && !strings.Contains("10", maze[right]) && right.x < len_X {
			// 	ones = append(ones, right)
			// 	maze[right] = "1"
			// }
			// if _, ok := loop[left]; !ok && !strings.Contains("10", maze[left]) && 0 <= left.x {
			// 	zeros = append(zeros, left)
			// 	maze[left] = "0"
			// }
		}
		if pos.x < prev_pos.x {
			add_oz(down, "1", &ones, loop, maze, len_X, len_Y)
			add_oz(up, "0", &zeros, loop, maze, len_X, len_Y)
			// if _, ok := loop[down]; !ok && !strings.Contains("10", maze[down]) && down.y < len_Y {
			// 	ones = append(ones, down)
			// 	maze[down] = "1"
			// }
			// if _, ok := loop[up]; !ok && !strings.Contains("10", maze[up]) && 0 <= up.y {
			// 	zeros = append(zeros, up)
			// 	maze[up] = "0"
			// }
		}
		if pos.x > prev_pos.x {
			add_oz(down, "0", &zeros, loop, maze, len_X, len_Y)
			add_oz(up, "1", &ones, loop, maze, len_X, len_Y)
			// if _, ok := loop[down]; !ok && !strings.Contains("10", maze[down]) && down.y < len_Y {
			// 	zeros = append(zeros, down)
			// 	maze[down] = "0"
			// }
			// if _, ok := loop[up]; !ok && !strings.Contains("10", maze[up]) && 0 <= up.y {
			// 	ones = append(ones, up)
			// 	maze[up] = "1"
			// }
		}
		prev_pos = pos
	}
	// print_maze(maze, input)
	D := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	nb_ones := len(ones)
	nb_zeros := len(zeros)

	fmt.Println("ones", nb_ones, "zeros", nb_zeros, "    z+o", nb_ones+nb_zeros)
	fmt.Println("loop", len(loop), "    z+o+loop", nb_ones+nb_zeros+len(loop))
	fmt.Println("lenX*lenY", len_X*len_Y, "     reste", len_X*len_Y-(nb_ones+nb_zeros+len(loop)))
	fmt.Println()
	// Find other ones and zeros
	// fmt.Println(ones)
	// fmt.Println(zeros)
	write_maze(maze, loop, input)
	for nb_ones+nb_zeros+len(loop) < len_X*len_Y && (len(ones) != 0 || len(zeros) != 0) {
		// ONES
		if len(ones) > 0 {
			pos := ones[0]
			ones_to_add := []Pos{}
			for _, d := range D { // explore
				next := Pos{x: pos.x + d.x, y: pos.y + d.y}
				if _, ok := loop[next]; !ok && !strings.Contains("10", maze[next]) && in_map(next, len_X, len_Y) {
					ones_to_add = append(ones_to_add, next)
					maze[next] = "1"
					nb_ones++
				}
			}
			ones = ones[1:] // pop
			ones = append(ones, ones_to_add...)
		}
		// ZEROS
		if len(zeros) > 0 {
			pos := zeros[0]
			zeros_to_add := []Pos{}
			for _, d := range D { // explore
				next := Pos{x: pos.x + d.x, y: pos.y + d.y}
				if _, ok := loop[next]; !ok && !strings.Contains("10", maze[next]) && in_map(next, len_X, len_Y) {
					zeros_to_add = append(zeros_to_add, next)
					maze[next] = "0"
					nb_zeros++
				}
			}
			zeros = zeros[1:] // pop
			zeros = append(zeros, zeros_to_add...)
		}

		// fmt.Println("ok", len(ones), len(zeros))
	}
	// print_maze(maze, input)

	fmt.Println("ones", nb_ones, "zeros", nb_zeros, "    z+o", nb_ones+nb_zeros)
	fmt.Println("loop", len(loop), "    z+o+loop", nb_ones+nb_zeros+len(loop))
	fmt.Println("lenX*lenY", len_X*len_Y, "     reste", len_X*len_Y-(nb_ones+nb_zeros+len(loop)))
	fmt.Println()
	return []int{nb_ones, nb_zeros}
}

func main() {
	file := utils.ReadFile("../inputs/10/input1_ex.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
	// between  [405 5906]    len_X*len_Y-9
}
