package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Pos struct{ x, y int }

func buildTower(jetPattern string, nbRocksTot int) (int, map[Pos]uint8) {
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
	for nbRocks < nbRocksTot {
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
				case '>':
					_, ok := tower[Pos{x: pos.x + 4, y: pos.y}]
					if !(pos.x+3+1 > 6) && !ok {
						pos.x++
					}
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 {
						pos.x++
					}
				}
				_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y}]
				_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y}]
				_, ok3 := tower[Pos{x: pos.x, y: pos.y - 1}]
				if ok1 || ok2 || ok3 {
					stopped = true
				} else {
					pos.y--
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x + 2, y: pos.y + 2}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 && !ok3 {
						pos.x++
					}
				}
				_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y - 1}]
				_, ok2 := tower[Pos{x: pos.x, y: pos.y - 1}]
				_, ok3 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
				if ok1 || ok2 || ok3 {
					stopped = true
				} else {
					pos.y--
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 1, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x + 1, y: pos.y + 2}]
					_, ok4 := tower[Pos{x: pos.x + 1, y: pos.y + 3}]
					if !(pos.x+1 > 6) && !ok1 && !ok2 && !ok3 && !ok4 {
						pos.x++
					}
				}
				if _, ok := tower[Pos{x: pos.x, y: pos.y - 1}]; ok {
					stopped = true
				} else {
					pos.y--
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 {
						pos.x++
					}
				}
				_, ok1 := tower[Pos{x: pos.x, y: pos.y - 1}]
				_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
				if ok1 || ok2 {
					stopped = true
				} else {
					pos.y--
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

func buildTower2(jetPattern string, nbRocksTot int, initPattI int, firstRock int) (int, map[Pos]uint8) {
	tower := map[Pos]uint8{
		{x: 0, y: -1}: 1,
		{x: 1, y: -1}: 1,
		{x: 2, y: -1}: 1,
		{x: 3, y: -1}: 1,
		{x: 4, y: -1}: 1,
		{x: 5, y: -1}: 1,
		{x: 6, y: -1}: 1}
	towerLen := 0
	rock := firstRock
	nbRocks := 0
	pattI := initPattI
	fmt.Println("ini pattI", pattI)
	jetDirection := jetPattern[pattI]
	flatLvls := map[int][]int{}
	foundCycle := false
	fmt.Println(flatLvls)
	fmt.Println(nbRocksTot)
	for nbRocks < nbRocksTot && !foundCycle {
		switch rock {
		case 0: // -
			pos := Pos{x: 2, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok := tower[Pos{x: pos.x - 1, y: pos.y}]
					if !(pos.x-1 < 0) && !ok {
						pos.x--
					}
				case '>':
					_, ok := tower[Pos{x: pos.x + 4, y: pos.y}]
					if !(pos.x+3+1 > 6) && !ok {
						pos.x++
					}
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
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x + 1, y: pos.y}] = 1
			tower[Pos{x: pos.x + 2, y: pos.y}] = 1
			tower[Pos{x: pos.x + 3, y: pos.y}] = 1
			towerLen = max(towerLen, towerLen+1-(towerLen-pos.y))
		case 1: // +
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 {
						pos.x++
					}
				}
				_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y}]
				_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y}]
				_, ok3 := tower[Pos{x: pos.x, y: pos.y - 1}]
				if ok1 || ok2 || ok3 {
					stopped = true
				} else {
					pos.y--
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
		case 2: // L
			pos := Pos{x: 3, y: towerLen + 3}
			stopped := false
			for !stopped {
				switch jetDirection {
				case '<':
					_, ok := tower[Pos{x: pos.x - 2, y: pos.y}]
					if !(pos.x-2 < 0) && !ok {
						pos.x--
					}
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x + 2, y: pos.y + 2}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 && !ok3 {
						pos.x++
					}
				}
				_, ok1 := tower[Pos{x: pos.x - 1, y: pos.y - 1}]
				_, ok2 := tower[Pos{x: pos.x, y: pos.y - 1}]
				_, ok3 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
				if ok1 || ok2 || ok3 {
					stopped = true
				} else {
					pos.y--
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
		case 3: // |
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 1, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y + 1}]
					_, ok3 := tower[Pos{x: pos.x + 1, y: pos.y + 2}]
					_, ok4 := tower[Pos{x: pos.x + 1, y: pos.y + 3}]
					if !(pos.x+1 > 6) && !ok1 && !ok2 && !ok3 && !ok4 {
						pos.x++
					}
				}
				if _, ok := tower[Pos{x: pos.x, y: pos.y - 1}]; ok {
					stopped = true
				} else {
					pos.y--
				}
				pattI = (pattI + 1) % len(jetPattern)
				jetDirection = jetPattern[pattI]
			}
			tower[Pos{x: pos.x, y: pos.y}] = 1
			tower[Pos{x: pos.x, y: pos.y + 1}] = 1
			tower[Pos{x: pos.x, y: pos.y + 2}] = 1
			tower[Pos{x: pos.x, y: pos.y + 3}] = 1
			towerLen = max(towerLen, towerLen+4-(towerLen-pos.y))
		case 4: // ::
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
				case '>':
					_, ok1 := tower[Pos{x: pos.x + 2, y: pos.y}]
					_, ok2 := tower[Pos{x: pos.x + 2, y: pos.y + 1}]
					if !(pos.x+2 > 6) && !ok1 && !ok2 {
						pos.x++
					}
				}
				_, ok1 := tower[Pos{x: pos.x, y: pos.y - 1}]
				_, ok2 := tower[Pos{x: pos.x + 1, y: pos.y - 1}]
				if ok1 || ok2 {
					stopped = true
				} else {
					pos.y--
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
		// Part 2
		if isFlatLvl(towerLen-1, tower) {
			flatLvls[towerLen-1] = []int{rock, pattI, nbRocks}
		}
		for k, v := range flatLvls {
			if v[0] == 0 && v[1] == pattI && k != towerLen-1 {
				fmt.Println("nbRocks", nbRocks, k)
				foundCycle = true
				//cycleBeginLvl := k
				lenCycle := towerLen - k
				nbRocksCycle := nbRocks - v[2]
				fmt.Println(flatLvls)
				fmt.Println("lent", towerLen-1, k)
				fmt.Println(nbRocksCycle, nbRocks, v[2])
				nbRocksBeforeCycle := v[2]
				nbCycle := (nbRocksTot - nbRocksBeforeCycle) / nbRocksCycle
				nbRocksRest := (nbRocksTot - nbRocksBeforeCycle) % nbRocksCycle
				fmt.Println("rest", nbRocksRest)
				fmt.Println(nbRocksBeforeCycle + nbCycle*nbRocksCycle + nbRocksRest)
				fmt.Println("In restH")
				restH, _ := buildTower2(jetPattern, nbRocksRest, pattI, (rock+1)%5)
				fmt.Println("lenCycle", lenCycle, nbCycle, nbRocksCycle)
				towerLen = k + lenCycle*nbCycle + restH
				break
			}
		}
		//
		rock = (rock + 1) % 5
		nbRocks++
		// fmt.Println(nbRocks)
		// displayTower(tower, towerLen)
		// fmt.Println()
	}
	fmt.Println("RETURN", foundCycle, flatLvls)
	return towerLen, tower
}

func isFlatLvl(y int, tower map[Pos]uint8) bool {
	flat := true
	for x := 0; x < 7; x++ {
		if _, ok := tower[Pos{x: x, y: y}]; !ok {
			flat = false
		}
	}
	return flat
}

func findSameSituation(tower map[Pos]uint8,  maxLevel int, h int) bool {
	for y := h; y < maxLevel-h; y++ {
		flat := true
		for x := 0; x < 7; x++ {
			if _, ok := tower[Pos{x: x, y: y}]; !ok {
				flat = false
			}
		}
		if flat == true {
			flatLvls = append(flatLvls, y)
		}
	}
}

func findFlat(tower map[Pos]uint8, maxLevel int) []int {
	flatLvls := []int{}
	for y := -1; y < maxLevel; y++ {
		flat := true
		for x := 0; x < 7; x++ {
			if _, ok := tower[Pos{x: x, y: y}]; !ok {
				flat = false
			}
		}
		if flat == true {
			flatLvls = append(flatLvls, y)
		}
	}
	return flatLvls
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
	start := time.Now()
	H, _ := buildTower2(input, 100000, 0, 0)
	//H, _ := buildTower(input, 2022000)
	//fmt.Println(findFlat(tower, H))
	fmt.Println("Part 1 :", H, "- Time :", time.Since(start))
}
