package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func ABS(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := ioutil.ReadFile("inputs/day21/input.txt")
	if err != nil {
		panic(err)
	}

	//file := "root: pppw + sjmn\nsjmn: a + b\na: 1\nb: 1\npppw: cczh + humn\ncczh: 1\nhumn: 5"
	input := strings.Split(string(file), "\n")
	// Part 1
	start := time.Now()
	monkeyMap := parseInput(input)
	//fmt.Println(parseInput(input))
	//part1 := yell("root", parseInput(input))
	part1 := []int{yell("root", parseInput(input)), yell(monkeyMap["root"][0], parseInput(input)), yell(monkeyMap["root"][2], parseInput(input))}
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))

	// Part 2
	fmt.Println("Part 2 :")
	start = time.Now()
	//fmt.Println(parseInput(input))
	//part2 := dicho(3*10^12, 4*10^12, parseInput(input))
	//part2 := brutasse(parseInput(input))
	//fmt.Println(monkeyMap["root"][0], monkeyMap["root"][2])
	//fmt.Println(yell2(5, "ptdq", monkeyMap))
	//fmt.Println(yell("ptdq", monkeyMap))
	part2 := []int{yell2(3403989691757, monkeyMap["root"][0], monkeyMap), yell2(3403989691757, monkeyMap["root"][2], monkeyMap)}
	//part2 := []int{yell(monkeyMap["root"][0], monkeyMap), yell(monkeyMap["root"][2], monkeyMap)}

	//part2 := part2(monkeyMap)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}

func parseInput(input []string) map[string][]string {
	monkeyMap := map[string][]string{}
	for _, line := range input {
		L := strings.Split(line, ": ")
		name := L[0]
		content := strings.Split(L[1], " ")
		if len(content) == 1 {
			monkeyMap[name] = []string{content[0]}
		} else {
			monkeyMap[name] = []string{content[0], content[1], content[2]}
		}
	}
	return monkeyMap
}

func yell(name string, monkeyMap map[string][]string) int {
	if len(monkeyMap[name]) == 1 {
		n, _ := strconv.Atoi(monkeyMap[name][0])
		return n
	} else {
		op := monkeyMap[name][1]
		if op == "+" {
			return yell(monkeyMap[name][0], monkeyMap) + yell(monkeyMap[name][2], monkeyMap)
		}
		if op == "*" {
			return yell(monkeyMap[name][0], monkeyMap) * yell(monkeyMap[name][2], monkeyMap)
		}
		if op == "-" {
			return yell(monkeyMap[name][0], monkeyMap) - yell(monkeyMap[name][2], monkeyMap)
		}
		if op == "/" {
			return yell(monkeyMap[name][0], monkeyMap) / yell(monkeyMap[name][2], monkeyMap)
		}
		return 0
	}
}

func yell2(x int, name string, monkeyMap map[string][]string) int {
	if len(monkeyMap[name]) == 1 {
		n, _ := strconv.Atoi(monkeyMap[name][0])
		return n
	} else {
		op := monkeyMap[name][1]
		if op == "+" {
			if monkeyMap[name][0] == "humn" {
				return x + yell2(x, monkeyMap[name][2], monkeyMap)
			}
			if monkeyMap[name][2] == "humn" {
				return yell2(x, monkeyMap[name][0], monkeyMap) + x
			}
			return yell2(x, monkeyMap[name][0], monkeyMap) + yell2(x, monkeyMap[name][2], monkeyMap)
		}
		if op == "*" {
			if monkeyMap[name][0] == "humn" {
				return x * yell2(x, monkeyMap[name][2], monkeyMap)
			}
			if monkeyMap[name][2] == "humn" {
				return yell2(x, monkeyMap[name][0], monkeyMap) * x
			}
			return yell2(x, monkeyMap[name][0], monkeyMap) * yell2(x, monkeyMap[name][2], monkeyMap)
		}
		if op == "-" {
			//println(name, monkeyMap[name][0], monkeyMap[name][2])
			//fmt.Println(yell2(x, monkeyMap[name][0], monkeyMap), yell2(x, monkeyMap[name][2], monkeyMap))
			if monkeyMap[name][0] == "humn" {
				return x - yell2(x, monkeyMap[name][2], monkeyMap)
			}
			if monkeyMap[name][2] == "humn" {
				return yell2(x, monkeyMap[name][0], monkeyMap) - x
			}
			return yell2(x, monkeyMap[name][0], monkeyMap) - yell2(x, monkeyMap[name][2], monkeyMap)
		}
		if op == "/" {
			if monkeyMap[name][0] == "humn" {
				return x / yell2(x, monkeyMap[name][2], monkeyMap)
			}
			if monkeyMap[name][2] == "humn" {
				return yell2(x, monkeyMap[name][0], monkeyMap) / x
			}
			return yell2(x, monkeyMap[name][0], monkeyMap) / yell2(x, monkeyMap[name][2], monkeyMap)
		}
		return 0
	}
}

func transformMap(name string, monkeyMap map[string][]string) int {
	println("name", name)
	if name == monkeyMap["root"][0] {
		monkeyMap[monkeyMap["root"][0]] = monkeyMap[monkeyMap["root"][2]]
	}
	if len(monkeyMap[name]) == 1 {
		println("value number")
		return 0
	} else {
		println("inverser eq")
		a := 0
		b := 0
		// Find another relation where name is in the rhs
		for otherName, content := range monkeyMap {
			if len(content) == 3 {
				if content[0] == name {
					if content[1] == "+" {
						monkeyMap[name] = []string{otherName, "-", content[2]}
					}
					if content[1] == "-" {
						monkeyMap[name] = []string{otherName, "+", content[2]}
					}
					if content[1] == "*" {
						monkeyMap[name] = []string{otherName, "/", content[2]}
					}
					if content[1] == "/" {
						monkeyMap[name] = []string{otherName, "*", content[2]}
					}
					a = transformMap(content[0], monkeyMap)
				}
				if content[2] == name {
					if content[1] == "+" {
						monkeyMap[name] = []string{otherName, "-", content[0]}
					}
					if content[1] == "-" {
						monkeyMap[name] = []string{content[0], "-", otherName}
					}
					if content[1] == "*" {
						monkeyMap[name] = []string{otherName, "/", content[0]}
					}
					if content[1] == "/" {
						monkeyMap[name] = []string{content[0], "/", otherName}
					}
					b = transformMap(content[2], monkeyMap)
				}
			}
		}
		return a + b
	}
}

func part2(monkeyMap map[string][]string) int {
	monkeyMap["humn"] = []string{"0", "0", "0"} // good size
	monkeyMap[monkeyMap["root"][2]] = []string{strconv.Itoa(yell(monkeyMap["root"][2], monkeyMap))}

	delete(monkeyMap, "root")
	_ = transformMap("humn", monkeyMap)
	//fmt.Println(monkeyMap)
	return yell("humn", monkeyMap)
}

func dicho(a int, b int, monkeyMap map[string][]string) int {
	found := false
	for !found {
		m := (a + b) / 2
		f_a := yell2(a, monkeyMap["root"][0], monkeyMap) - yell2(a, monkeyMap["root"][2], monkeyMap)
		f_m := yell2(m, monkeyMap["root"][0], monkeyMap) - yell2(m, monkeyMap["root"][2], monkeyMap)
		if f_a == 0 {
			fmt.Println("a", a)
			return a
		}
		if f_m == 0 {
			fmt.Println("m", m, "et", a, b)
			return m
		}
		if f_a*f_m <= 0 {
			b = m
		} else {
			a = m
		}
	}
	return (a + b) / 2
}

func brutasse(monkeyMap map[string][]string) int {
	for i := 100000; i < 1000000; i++ {
		if yell2(i, monkeyMap["root"][0], monkeyMap) == yell2(i, monkeyMap["root"][2], monkeyMap) {
			return i
		}
	}
	return 0
}

// /*
// def dichotomie(f,a,b,e):
//     delta = 1
//     while delta > e:
//         m = (a + b) / 2
//         delta = abs(b - a)
//         if f(m) == 0:
//             return m
//         elif f(a) * f(m)  > 0:
//             a = m
//         else:
//             b = m
//     return a, b

// */
// /*
// --- Part Two ---

// Due to some kind of monkey-elephant-human mistranslation, you seem to have misunderstood a few key details about the riddle.

// First, you got the wrong job for the monkey named root; specifically, you got the wrong math operation. The correct operation for monkey root should be =, which means that it still listens for two numbers (from the same two monkeys as before), but now checks that the two numbers match.

// Second, you got the wrong monkey for the job starting with humn:. It isn't a monkey - it's you. Actually, you got the job wrong, too: you need to figure out what number you need to yell so that root's equality check passes. (The number that appears after humn: in your input is now irrelevant.)

// In the above example, the number you need to yell to pass root's equality test is 301. (This causes root to get the same number, 150, from both of its monkeys.)

// What number do you yell to pass root's equality test?

// while (monkeyMap[name][0] =)

// */
