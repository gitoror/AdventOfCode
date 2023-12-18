package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	dir string
	dst int
}

type Pos struct{ x, y int }

func part1(input []string) int {
	dig_plan := []Instruction{}
	for _, line := range input {
		infos := strings.Split(line, " ")
		dig_plan = append(dig_plan, Instruction{infos[0], utils.ToInt(infos[1])})
	}
	// retrieve vertices (xi,yi)
	vertices := []Pos{}
	x, y := 0, 0
	for _, instr := range dig_plan {
		switch instr.dir {
		case "U":
			y -= instr.dst
		case "D":
			y += instr.dst
		case "R":
			x += instr.dst
		case "L":
			x -= instr.dst
		}
		vertices = append(vertices, Pos{x, y})
	}
	// len loop
	loop_len := 0
	for i := 0; i < len(dig_plan); i++ {
		loop_len += dig_plan[i].dst
	}
	// calc the area of the polygon
	area := 0
	for i := 0; i < len(vertices); i++ {
		area += vertices[i].x * vertices[(i+1)%len(vertices)].y
		area -= vertices[i].y * vertices[(i+1)%len(vertices)].x
	}
	area = utils.AbsInt(area)/2 + loop_len/2 + 1
	return area
}

type Instruction2 struct {
	dir int
	dst int
}

func part2(input []string) int {
	dig_plan := []Instruction2{}
	for _, line := range input {
		infos := strings.Split(line, " ") // infos[2] is like #12ab20 (hexa)
		// Each hexadecimal code is six hexadecimal digits long. The first five hexadecimal digits encode the distance in meters as a five-digit hexadecimal number. The last hexadecimal digit encodes the direction to dig: 0 means R, 1 means D, 2 means L, and 3 means U.
		dir := int(infos[2][len(infos[2])-2] - '0')
		dst, _ := strconv.ParseInt(infos[2][2:len(infos[2])-2], 16, 64)
		dig_plan = append(dig_plan, Instruction2{dir, int(dst)})
	}
	// retrieve vertices (xi,yi)
	vertices := []Pos{}
	x, y := 0, 0
	// 0 means R, 1 means D, 2 means L, and 3 means U.
	for _, instr := range dig_plan {
		switch instr.dir {
		case 3:
			y -= instr.dst
		case 1:
			y += instr.dst
		case 0:
			x += instr.dst
		case 2:
			x -= instr.dst
		}
		vertices = append(vertices, Pos{x, y})
	}
	// len loop
	loop_len := 0
	for i := 0; i < len(dig_plan); i++ {
		loop_len += dig_plan[i].dst
	}
	// calc the area of the polygon
	area := 0
	for i := 0; i < len(vertices); i++ {
		area += vertices[i].x * vertices[(i+1)%len(vertices)].y
		area -= vertices[i].y * vertices[(i+1)%len(vertices)].x
	}
	area = utils.AbsInt(area)/2 + loop_len/2 + 1
	return area
}

func main() {
	file := utils.ReadFile("../inputs/18/in.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))

}
