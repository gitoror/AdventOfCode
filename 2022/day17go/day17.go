package main

import (
	"fmt"
	"io/ioutil"
)

type Pos struct{ x, y int }

func buildTower(jetPattern string) (int, map[Pos]uint8) {
	tower := map[Pos]uint8{
		{x: 0, y: -1}: 1,
		{x: 1, y: -1}: 1,
		{x: 2, y: -1}: 1,
		{x: 3, y: -1}: 1,
		{x: 4, y: -1}: 1,
		{x: 5, y: -1}: 1,
		{x: 6, y: -1}: 1}
	towerLen := 0
	rock := 0
	nbRocks := 0
	pattI := 0
	jetDirection := jetPattern[pattI]
	for nbRocks < 2022 {
		switch rock {
		case 0:
			pos := Pos{x: 2, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok := tower[Pos{x: pos.x - 1, y: pos.y}]
					if !(pos.x-1 < 0) && !ok {
						pos.x--
					}
					_, ok1 := tower[Pos{x: pos.x, y: pos.y - 1}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
					_, ok3 := tower[Pos{x: pos.x + 2, y: pos.y - 1}]
					_, ok4 := tower[Pos{x: pos.x + 3, y: pos.y - 1}]
					if ok1 || ok2 || ok3 || ok4 {
						stopped = true
					} else {
						pos.y--
					}
				case '>':
					_, ok := tower[Pos{x: pos.x + 4, y: pos.y}]
					if !(pos.x+3+1 > 6) && !ok {
						pos.x++
					}
					_, ok1 := tower[Pos{x: pos.x, y: pos.y - 1}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
					_, ok3 := tower[Pos{x: pos.x + 2, y: pos.y - 1}]
					_, ok4 := tower[Pos{x: pos.x + 3, y: pos.y - 1}]
					if ok1 || ok2 || ok3 || ok4 {
						stopped = true
					} else {
						pos.y--
					}
				}
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y}] = 1
			tower[Pos{x: pos.x + 2, y: pos.y}] = 1
			tower[Pos{x: pos.x + 3, y: pos.y}] = 1
			towerLen = max(towerLen, towerLen+1-(towerLen-pos.y))
		case 1:
			pos := Pos{x: 3, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok1 := tower[Pos{x: pos.x - 2, y: pos.y + 1}]
					_, ok2 := tower[Pos{x: pos.x - 1, y: pos.y}]
					if !(pos.x-2 < 0) && !ok1 && !ok2 {
						pos.x--
					}
					_, ok1 = tower[Pos{x: pos.x - 1, y: pos.y}]
					_, ok2 = tower[Pos{x: pos.x + 1, y: pos.y}]
					_, ok3 := tower[Pos{x: pos.x, y: pos.y - 1}]
					if ok1 || ok2 || ok3 {
						stopped = true
					} else {
						pos.y--
					}
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 {
						pos.x++
					}
					_, ok1 = tower[Pos{x: pos.x - 1, y: pos.y}]
					_, ok2 = tower[Pos{x: pos.x + 1, y: pos.y}]
					_, ok3 := tower[Pos{x: pos.x, y: pos.y - 1}]
					if ok1 || ok2 || ok3 {
						stopped = true
					} else {
						pos.y--
					}
				}
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x, y: pos.y + 1}] = 1
			tower[Pos{x: pos.x, y: pos.y + 2}] = 1
			tower[Pos{x: pos.x - 1, y: pos.y + 1}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y + 1}] = 1
			towerLen = max(towerLen, towerLen+3-(towerLen-pos.y))
		case 2:
			pos := Pos{x: 3, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok := tower[Pos{x: pos.x - 2, y: pos.y}]
					if !(pos.x-2 < 0) && !ok {
						pos.x--
					}
					_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y - 1}]
					_, ok2 := tower[Pos{x: pos.x, y: pos.y - 1}]
					_, ok3 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
					if ok1 || ok2 || ok3 {
						stopped = true
					} else {
						pos.y--
					}
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x + 2, y: pos.y + 2}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 && !ok3 {
						pos.x++
					}
					_, ok1 = tower[Pos{x: pos.x - 1, y: pos.y - 1}]
					_, ok2 = tower[Pos{x: pos.x, y: pos.y - 1}]
					_, ok3 = tower[Pos{x: pos.x + 1, y: pos.y - 1}]
					if ok1 || ok2 || ok3 {
						stopped = true
					} else {
						pos.y--
					}
				}
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x - 1, y: pos.y}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y + 1}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y + 2}] = 1
			towerLen = max(towerLen, towerLen+3-(towerLen-pos.y))
		case 3:
			pos := Pos{x: 2, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x - 1, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x - 1, y: pos.y + 2}]
					_, ok4 := tower[Pos{x: pos.x - 1, y: pos.y + 3}]
					if !(pos.x-1 < 0) && !ok1 && !ok2 && !ok3 && !ok4 {
						pos.x--
					}
					if _, ok := tower[Pos{x: pos.x, y: pos.y - 1}]; ok {
						stopped = true
					} else {
						pos.y--
					}
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 1, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x + 1, y: pos.y + 2}]
					_, ok4 := tower[Pos{x: pos.x + 1, y: pos.y + 3}]
					if !(pos.x+1 > 6) && !ok1 && !ok2 && !ok3 && !ok4 {
						pos.x++
					}
					if _, ok := tower[Pos{x: pos.x, y: pos.y - 1}]; ok {
						stopped = true
					} else {
						pos.y--
					}
				}
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x, y: pos.y + 1}] = 1
			tower[Pos{x: pos.x, y: pos.y + 2}] = 1
			tower[Pos{x: pos.x, y: pos.y + 3}] = 1
			towerLen = max(towerLen, towerLen+4-(towerLen-pos.y))
		case 4:
			pos := Pos{x: 2, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x - 1, y: pos.y + 1}]
					if !(pos.x-1 < 0) && !ok1 && !ok2 {
						pos.x--
					}
					_, ok1 = tower[Pos{x: pos.x, y: pos.y - 1}]
					_, ok2 = tower[Pos{x: pos.x + 1, y: pos.y - 1}]
					if ok1 || ok2 {
						stopped = true
					} else {
						pos.y--
					}
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 {
						pos.x++
					}
					_, ok1 = tower[Pos{x: pos.x, y: pos.y - 1}]
					_, ok2 = tower[Pos{x: pos.x + 1, y: pos.y - 1}]
					if ok1 || ok2 {
						stopped = true
					} else {
						pos.y--
					}
				}
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y}] = 1
			tower[Pos{x: pos.x, y: pos.y + 1}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y + 1}] = 1
			towerLen = max(towerLen, towerLen+2-(towerLen-pos.y))
		}
		rock = (rock + 1) % 5
		nbRocks++
		// fmt.Println(nbRocks)
		// displayTower(tower, towerLen)
		// fmt.Println()
	}
	return towerLen, tower
}

func displayTower(tower map[Pos]uint8, towerLen int) {
	for y := towerLen - 1; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if _, ok := tower[Pos{x: x, y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := ioutil.ReadFile("inputs/day17/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	//input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	fmt.Println(len(input))

	H, _ := buildTower(input)

	fmt.Println(H)
}
