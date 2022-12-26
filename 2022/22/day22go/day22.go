package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := ioutil.ReadFile("inputs/day22/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n\n")
	//board := input[0]
	//display(parse(board))
	start := time.Now()
	part1, newBoard := followPath(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	display(newBoard)
}

type Pos struct{ x, y int }

func parse(board string) map[Pos]string {
	boardLines := strings.Split(board, "\n")
	mapp := map[Pos]string{}
	for i := range boardLines {
		for j, c := range boardLines[i] {
			if c == '.' {
				mapp[Pos{x: j + 1, y: i + 1}] = "."
			}
			if c == '#' {
				mapp[Pos{x: j + 1, y: i + 1}] = "#"
			}
		}
	}
	return mapp
}

func display(mapp map[Pos]string) {
	board := ""
	// Find max min
	maxX := 0
	maxY := 0
	for k := range mapp {
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	// write in the board
	for i := 1; i <= maxY; i++ {
		for j := 1; j <= maxX; j++ {
			if v, ok := mapp[Pos{x: j, y: i}]; ok {
				board += v
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	fmt.Println(board)
}

func followPath(input []string) (int, map[Pos]string) {
	board := parse(input[0])
	path := input[1]
	minX_Y1 := math.MaxInt
	for k := range board {
		if k.y == 1 && k.x < minX_Y1 {
			minX_Y1 = k.x
		}
	}
	pos := Pos{x: minX_Y1, y: 1} // starting pos
	dir := 0                     // facing to the right initially
	board[pos] = ">"
	dst := string(path[0]) // has to be an int bcs start facing R
	for i := 1; i < len(path); i++ {
		v := path[i]
		if v == 'R' || v == 'L' {
			dstInt, _ := strconv.Atoi(dst)
			pos, board = move(pos, dir, dstInt, board)
			dst = ""
			if v == 'R' {
				dir = (dir + 1) % 4
			}
			if v == 'L' {
				dir = (dir - 1 + 4) % 4
			}
		} else { // build dst
			dst += string(v)
			if i == len(path)-1 {
				dstInt, _ := strconv.Atoi(dst)
				pos, board = move(pos, dir, dstInt, board)
			}
		}
		board[pos] = symbol[dir]
	}
	return 1000*pos.y + 4*pos.x + dir, board
}

var symbol []string = []string{">", "v", "<", "^"}

func move(pos Pos, dir int, dst int, board map[Pos]string) (Pos, map[Pos]string) {
	// move dir, dst
	for step := 0; step < dst; step++ {
		if dir == 0 { // >
			if v, ok := board[Pos{x: pos.x + 1, y: pos.y}]; ok {
				if v != "#" {
					pos.x = pos.x + 1
				} else {
					break
				}
			} else { // wrap around
				// find opposite edge of the board
				temp := Pos{x: pos.x - 1, y: pos.y}
				for _, ok := board[temp]; ok; _, ok = board[temp] {
					temp.x--
				}
				temp.x++
				if board[Pos{x: temp.x, y: temp.y}] != "#" {
					pos = temp
				} else {
					break
				}
			}
		} else if dir == 1 { // v
			if v, ok := board[Pos{x: pos.x, y: pos.y + 1}]; ok {
				if v != "#" {
					pos.y = pos.y + 1
				} else {
					break
				}
			} else {
				temp := Pos{x: pos.x, y: pos.y - 1}
				for _, ok := board[temp]; ok; _, ok = board[temp] {
					temp.y--
				}
				temp.y++
				if board[Pos{x: temp.x, y: temp.y}] != "#" {
					pos = temp
				} else {
					break
				}
			}
		} else if dir == 2 { // <
			if v, ok := board[Pos{x: pos.x - 1, y: pos.y}]; ok {
				if v != "#" {
					pos.x = pos.x - 1
				} else {
					break
				}
			} else {
				temp := Pos{x: pos.x + 1, y: pos.y}
				for _, ok := board[temp]; ok; _, ok = board[temp] {
					temp.x++
				}
				temp.x--
				if board[Pos{x: temp.x, y: temp.y}] != "#" {
					pos = temp
				} else {
					break
				}
			}
		} else if dir == 3 { // ^
			if v, ok := board[Pos{x: pos.x, y: pos.y - 1}]; ok {
				if v != "#" {
					pos.y = pos.y - 1
				} else {
					break
				}
			} else {
				temp := Pos{x: pos.x, y: pos.y + 1}
				for _, ok := board[temp]; ok; _, ok = board[temp] {
					temp.y++
				}
				temp.y--
				if board[Pos{x: temp.x, y: temp.y}] != "#" {
					pos = temp
				} else {
					break
				}
			}
		}
		board[pos] = symbol[dir]
	}
	return pos, board
}
