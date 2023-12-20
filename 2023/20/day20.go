package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

const (
	BROADCASTER = "broadcaster"
	FLIPFLOP    = "flipflop"
	CONJUNCTION = "conjunction"
	UNTYPED     = "untyped"
)

type Module struct {
	Id         string
	Category   string          // flipflop, conjuction, broadcaster
	OutModules []*Module       // flipflop, conjuction, broadcaster
	On         int             // flipflop
	Memory     map[*Module]int // conjuction
}

type Instruction struct {
	From  *Module
	Pulse int
	To    *Module
}

type Instructions []Instruction

func createModules(input []string) map[string]*Module {
	modules := map[string]*Module{}
	temp := map[string][]string{}
	// Populate modules and temp
	for _, line := range input {
		moduleInfos := strings.Split(line, " -> ")
		var id, category string
		if moduleInfos[0] == BROADCASTER {
			id, category = BROADCASTER, BROADCASTER
		} else if moduleInfos[0][0] == '%' {
			id, category = moduleInfos[0][1:], FLIPFLOP
		} else {
			id, category = moduleInfos[0][1:], CONJUNCTION
		}
		outModules := strings.Split(moduleInfos[1], ", ")
		temp[id] = outModules
		modules[id] = &Module{id, category, []*Module{}, 0, map[*Module]int{}}
	}
	// Populate OutModules and Memory
	for id, outModules := range temp {
		for _, outModule := range outModules {
			if _, ok := modules[outModule]; ok {
				if modules[outModule].Category == CONJUNCTION {
					modules[outModule].Memory[modules[id]] = 0
				}
			} else {
				modules[outModule] = &Module{outModule, UNTYPED, []*Module{}, 0, map[*Module]int{}}
			}
			modules[id].OutModules = append(modules[id].OutModules, modules[outModule])
		}
	}
	return modules
}

func (instructions *Instructions) add(instruction Instruction) {
	*instructions = append(*instructions, instruction)
}

func (instructions *Instructions) process() (int, int) {
	instruction := (*instructions)[0]
	lowPulses, highPulses := 0, 0
	if instruction.Pulse == 0 {
		lowPulses = 1
	} else {
		highPulses = 1
	}
	switch instruction.To.Category {
	case BROADCASTER:
		for _, outModule := range instruction.To.OutModules {
			instructions.add(Instruction{instruction.To, 0, outModule})
		}
	case FLIPFLOP:
		if instruction.Pulse == 0 {
			instruction.To.On = (instruction.To.On + 1) % 2
			for _, outModule := range instruction.To.OutModules {
				instructions.add(Instruction{instruction.To, instruction.To.On, outModule})
			}
		}
	case CONJUNCTION:
		instruction.To.Memory[instruction.From] = instruction.Pulse
		sum := 0
		for _, value := range instruction.To.Memory {
			sum += value
		}
		if sum == len(instruction.To.Memory) {
			for _, outModule := range instruction.To.OutModules {
				instructions.add(Instruction{instruction.To, 0, outModule})
			}
		} else {
			for _, outModule := range instruction.To.OutModules {
				instructions.add(Instruction{instruction.To, 1, outModule})
			}
		}
	}
	*instructions = (*instructions)[1:]
	return lowPulses, highPulses
}

func part1(input []string) int {
	modules := createModules(input)
	lowPulses, highPulses := 0, 0
	nOrders := 0
	for nOrders < 1000 {
		instructions := Instructions{Instruction{&Module{}, 0, modules[BROADCASTER]}} // button
		for len(instructions) != 0 {
			dLow, dHigh := instructions.process()
			lowPulses += dLow
			highPulses += dHigh
		}
		nOrders++
	}
	fmt.Println("lowPulses:", lowPulses, "highPulses:", highPulses)
	return lowPulses * highPulses
}

func part2(input []string) int {
	modules := createModules(input)
	lowPulses, highPulses := 0, 0
	nOrders := 1
	rxOn := false
	for !rxOn {
		instructions := Instructions{Instruction{&Module{}, 0, modules[BROADCASTER]}} // button
		for len(instructions) != 0 {
			// if instructions[0].To.Id == "rx" && instructions[0].Pulse == 0 {
			// 	rxOn = true
			// 	fmt.Println("rxOn:", rxOn)
			// 	break
			// } // too long
			for _, prev_id := range []string{"vd", "ns", "bh", "dl"} { // oui c'est sale .. flemme de le coder proprement
				if instructions[0].From.Id == prev_id && instructions[0].Pulse == 1 {
					fmt.Println(prev_id, ":", nOrders)
				}
			}
			dLow, dHigh := instructions.process()
			lowPulses += dLow
			highPulses += dHigh
		}
		nOrders++
	}
	fmt.Println("lowPulses:", lowPulses, "highPulses:", highPulses)
	return nOrders
}

func main() {
	file := utils.ReadFile("../inputs/20/in.txt")
	input := strings.Split(string(file), "\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
	a := []int{3761, 3767, 3779, 3881}
	fmt.Println(utils.LCM(a[0], a[1], a[2:]...))
}
