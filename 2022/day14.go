package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//go:embed inputs/day14/ex.txt
var input string

func main(){
	fmt.Println(input)
	input := strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	grid := map[Pos]byte{}

	for _, line := range lines{
		fmt.Println("ligne", line)
	}
	fmt.Println(reflect.TypeOf(lines))

}

type Pos struct{
	X int
	Y int
}

func buildPair(pair string) Pos {
	xy := strings.Split(pair, ",")
	X, _ := strconv.Atoi(xy[0])
	Y, _ := strconv.Atoi(xy[1])
	return Pos{X: X, Y: Y}
}



