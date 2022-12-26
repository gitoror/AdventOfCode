package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Pos struct{ x, y, z int }

func (p Pos) neighbors() []Pos {
	return []Pos{
		{p.x + 1, p.y, p.z},
		{p.x - 1, p.y, p.z},
		{p.x, p.y + 1, p.z},
		{p.x, p.y - 1, p.z},
		{p.x, p.y, p.z + 1},
		{p.x, p.y, p.z - 1},
	}
}

func area(input []string) int {
	mapp := make(map[Pos]uint8)
	for _, line := range input {
		coord := strings.Split(line, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		z, _ := strconv.Atoi(coord[2])
		mapp[Pos{x, y, z}] = 0
	}
	area := 0
	for p := range mapp {
		for _, neighbor := range p.neighbors() {
			if _, ok := mapp[neighbor]; ok {
				area -= 1
			}
		}
		area += 6
	}
	return area
}

func extremBounds(mapp map[Pos]uint8) (int, int, int, int, int, int) {
	minX := 50
	maxX := 0
	minY := 50
	maxY := 0
	minZ := 50
	maxZ := 0
	for p := range mapp {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
		if p.z > maxZ {
			maxZ = p.z
		}
		if p.z < minZ {
			minZ = p.z
		}
	} // -1 and +1 to make sure the lava is not touching the edges so we can go around it
	return minX - 1, maxX + 1, minY - 1, maxY + 1, minZ - 1, maxZ + 1
}

func coutourCubes(mapp map[Pos]uint8) map[Pos]uint8 {
	minX, maxX, minY, maxY, minZ, maxZ := extremBounds(mapp)
	start := Pos{x: minX, y: minY, z: minZ}
	contour := map[Pos]uint8{}
	explored := map[Pos]uint8{}
	toExplore := []Pos{start}
	explored[start] = 0
	for len(toExplore) > 0 {
		exploredPos := toExplore[len(toExplore)-1]
		toExplore = toExplore[:len(toExplore)-1]
		for _, p := range exploredPos.neighbors() {
			if _, ok := explored[p]; ok {
				continue
			}
			if p.x < minX || p.x > maxX || p.y < minY || p.y > maxY || p.z < minZ || p.z > maxZ {
				continue
			}
			if _, ok := mapp[p]; ok {
				if _, ok := contour[exploredPos]; !ok {
					contour[exploredPos] = 0
				}
				continue
			}
			toExplore = append(toExplore, p)
		}
		explored[exploredPos] = 0
	}
	return contour
}

func areaContour(input []string) int {
	mapp := make(map[Pos]uint8)
	for _, line := range input {
		coord := strings.Split(line, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		z, _ := strconv.Atoi(coord[2])
		mapp[Pos{x, y, z}] = 0
	}
	contour := coutourCubes(mapp)
	//fmt.Println(contour)
	//fmt.Println("len", len(contour))
	area := 0
	for p := range contour {
		for _, neighbor := range p.neighbors() {
			if _, ok := mapp[neighbor]; ok {
				area++
			}
		}
	}
	return area
}

func main() {
	file, err := ioutil.ReadFile("inputs/day18/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n")
	// Part 1
	start := time.Now()
	part1 := area(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := areaContour(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
