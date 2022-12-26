package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := ioutil.ReadFile("inputs/day20/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n")
	// Part 1
	start := time.Now()
	part1 := Part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	// Part 2
	start = time.Now()
	part2 := Part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}

func buildList(input []string) []int {
	var list []int
	for _, line := range input {
		num, _ := strconv.Atoi(line)
		list = append(list, num)
	}
	return list
}

func move(initIndex int, circularList []int, initList []int, corespIndex map[int]int) []int {
	V := initList[initIndex]
	//println("move", V)
	I := corespIndex[initIndex]
	n := len(circularList)
	if V%n != n {
		if V > 0 {
			newI := (I + V) % (n - 1)
			if newI != I {
				circularList = delete(I, circularList)
				//println("delete", I)
				circularList = insert(V, newI, circularList)
				//println("insert", newI)
				if newI > I {
					for k := 0; k < len(circularList); k++ {
						if I+1 <= corespIndex[k] && corespIndex[k] <= newI {
							corespIndex[k] = (corespIndex[k] - 1) % n
						}
					}
				}
				if newI < I {
					for k := 0; k < len(circularList); k++ {
						if newI <= corespIndex[k] && corespIndex[k] <= I-1 {
							corespIndex[k] = (corespIndex[k] + 1) % n
						}
					}
				}
				corespIndex[initIndex] = newI
			}
		}
		if V < 0 {
			newI := ((I+V)%(n-1) + n - 1) % (n - 1)
			if newI != I {
				circularList = delete(I, circularList)
				//println("delete", I)
				circularList = insert(V, newI, circularList)
				//println("insert", newI)
				if newI > I {
					for k := 0; k < len(circularList); k++ {
						if I+1 <= corespIndex[k] && corespIndex[k] <= newI {
							corespIndex[k] = (corespIndex[k] - 1) % n
						}
					}
				}
				if newI < I {
					for k := 0; k < len(circularList); k++ {
						if newI <= corespIndex[k] && corespIndex[k] <= I-1 {
							corespIndex[k] = (corespIndex[k] + 1) % n
						}
					}
				}
				corespIndex[initIndex] = newI
			}
		}
	}
	return circularList
}

func insert(v int, pos int, circularList []int) []int {
	circularList = append(circularList[:pos], append([]int{v}, circularList[pos:]...)...)
	return circularList
}

func delete(pos int, circularList []int) []int {
	circularList = append(circularList[:pos], circularList[pos+1:]...)
	return circularList
}

func initCorespIndex(list []int) map[int]int {
	corespIndex := map[int]int{}
	for i, _ := range list {
		corespIndex[i] = i
	}
	return corespIndex
}

func Part1(input []string) int {
	sl := buildList(input)
	initList := buildList(input)
	corespIndex := initCorespIndex(initList)
	//fmt.Println("init", sl, corespIndex)
	for k := 0; k < len(initList); k++ {
		sl = move(k, sl, initList, corespIndex)
		//fmt.Println(k, sl, corespIndex)
	}
	// Find 0
	index0 := 0
	for sl[index0] != 0 {
		index0++
	}
	println(sl[(index0+1000)%len(sl)], sl[(index0+2000)%len(sl)], sl[(index0+3000)%len(sl)])
	return sl[(index0+1000)%len(sl)] + sl[(index0+2000)%len(sl)] + sl[(index0+3000)%len(sl)]
}

func Part2(input []string) int {
	sl := buildList(input)
	initList := buildList(input)
	for i := 0; i < len(initList); i++ {
		initList[i] = initList[i] * 811589153
	}
	corespIndex := initCorespIndex(initList)
	//fmt.Println("init", sl, corespIndex)
	for step := 0; step < 10; step++ {
		for k := 0; k < len(initList); k++ {
			sl = move(k, sl, initList, corespIndex)
			//fmt.Println(k, sl, corespIndex)
		}
	}
	// Find 0
	index0 := 0
	for sl[index0] != 0 {
		index0++
	}
	println(sl[(index0+1000)%len(sl)], sl[(index0+2000)%len(sl)], sl[(index0+3000)%len(sl)])
	return sl[(index0+1000)%len(sl)] + sl[(index0+2000)%len(sl)] + sl[(index0+3000)%len(sl)]
}
