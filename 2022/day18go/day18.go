package main

import (
	"io/ioutil"
	"strconv"
	"strings"
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

func area2(input []string) int {
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
	bb := airBubbles(mapp)
	var bubbleArea int = 0
	for p := range bb {
		for _, neighbor := range p.neighbors() {
			if _, ok := mapp[neighbor]; ok {
				bubbleArea += 1
			}
		}
	}
	return area - bubbleArea
}

func airBubbles(mapp map[Pos]uint8) map[Pos]uint8 {
	bubbles := map[Pos]uint8{}
	max := 50
	// Find bubble points
	for p := range mapp {
		for _, n := range p.neighbors() {
			if _, ok := mapp[n]; ok {
				continue
			}
			blockedXU := false
			blockedXD := false
			blockedYU := false
			blockedYD := false
			blockedZU := false
			blockedZD := false
			exploreXU := Pos{x: n.x + 1, y: n.y, z: n.z}
			exploreXD := Pos{x: n.x - 1, y: n.y, z: n.z}
			exploreYU := Pos{x: n.x, y: n.y + 1, z: n.z}
			exploreYD := Pos{x: n.x, y: n.y - 1, z: n.z}
			exploreZU := Pos{x: n.x, y: n.y, z: n.z + 1}
			exploreZD := Pos{x: n.x, y: n.y, z: n.z - 1}
			for step := 0; step <= max; step++ {
				if _, ok := mapp[exploreXU]; ok {
					blockedXU = true
				}
				exploreXU.x++
				if _, ok := mapp[exploreXD]; ok {
					blockedXD = true
				}
				exploreXD.x--
				if _, ok := mapp[exploreYU]; ok {
					blockedYU = true
				}
				exploreYU.y++
				if _, ok := mapp[exploreYD]; ok {
					blockedYD = true
				}
				exploreYD.y--
				if _, ok := mapp[exploreZU]; ok {
					blockedZU = true
				}
				exploreZU.z++
				if _, ok := mapp[exploreZD]; ok {
					blockedZD = true
				}
				exploreZD.z--
				if blockedXD && blockedXU && blockedYD && blockedYU && blockedZD && blockedZU {
					//fmt.Println("bubble", n)
					bubbles[n] = 0
				} else {
					//fmt.Println("not bubble", n, blockedXD, blockedXU, blockedYD, blockedYU, blockedZD, blockedZU)
				}
			}
		}
	}
	return bubbles
}

func main() {
	file, err := ioutil.ReadFile("inputs/day18/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n")
	println(area(input))
	println(area2(input))
}
