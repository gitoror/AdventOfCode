package main

import (
	"2023/utils"
	"fmt"
	"strings"
	"time"
)

type Rule struct {
	cat       string
	comp      string
	v         int
	next_name string
}

type Part map[string]int

func parse(input []string) (map[string][]Rule, []Part) {
	workflows_infos := strings.Split(input[0], "\n")
	parts_infos := strings.Split(input[1], "\n")
	workflows := map[string][]Rule{}
	for _, w := range workflows_infos { // qqz{s>2770:qs,m<1801:hdj,R}
		rules := []Rule{}
		infos := strings.Split(w[:len(w)-1], "{")
		name := infos[0]
		rules_infos := strings.Split(infos[1], ",")
		for _, r := range rules_infos {
			rule := Rule{}
			a := strings.Split(r, ":")
			if len(a) == 1 {
				rule.next_name = a[0]
			} else {
				rule.cat = a[0][0:1]
				rule.comp = a[0][1:2]
				rule.v = utils.ToInt(a[0][2:])
				rule.next_name = a[1]
			}
			rules = append(rules, rule)
		}
		workflows[name] = rules
	}
	//
	parts := []Part{}
	for _, p := range parts_infos {
		part := Part{}
		infos := strings.Split(p[1:len(p)-1], ",")
		part["x"] = utils.ToInt(infos[0][2:])
		part["m"] = utils.ToInt(infos[1][2:])
		part["a"] = utils.ToInt(infos[2][2:])
		part["s"] = utils.ToInt(infos[3][2:])
		parts = append(parts, part)
	}
	return workflows, parts
}

func is_accepted(workflows map[string][]Rule, name string, part Part) bool {
	wf := workflows[name]
	if len(wf) == 0 {
		return name == "A"
	}
	for _, rule := range wf {
		if rule.comp == ">" {
			if part[rule.cat] > rule.v {
				return is_accepted(workflows, rule.next_name, part)
			}
		} else if rule.comp == "<" {
			if part[rule.cat] < rule.v {
				return is_accepted(workflows, rule.next_name, part)
			}
		} else { // last rule
			return is_accepted(workflows, rule.next_name, part)
		}
	}
	return false
}

func part1(input []string) int {
	workflows, parts := parse(input)
	sum := 0
	for _, part := range parts {
		if is_accepted(workflows, "in", part) {
			sum += part["x"] + part["m"] + part["a"] + part["s"]
		}
	}
	return sum
}

func nb_arr(ranges map[string][]int) int {
	count := 1
	for _, t := range ranges {
		count *= t[1] - t[0] + 1
	}
	return count
}

func copy_map(a map[string][]int) map[string][]int {
	b := make(map[string][]int)
	for key, value := range a {
		b[key] = append([]int{}, value...)
	}
	return b
}

func count2(workflows map[string][]Rule, name string, ranges map[string][]int) int {
	for _, l := range []string{"x", "m", "a", "s"} {
		if ranges[l][1] < ranges[l][0] {
			return 0
		}
	}
	if name == "A" {
		// fmt.Println(nb_arr(ranges))
		return nb_arr(ranges)
	} else if name == "R" {
		return 0
	} else {
		sum := 0
		rules := workflows[name]
		new_ranges := copy_map(ranges)
		for _, rule := range rules {
			if rule.comp == "<" {
				new_ranges[rule.cat] = []int{ranges[rule.cat][0], rule.v - 1} // hyp toujours bien ordonné (j'ai test avec Min ca change rien)
				// fmt.Println("1", name, rule, new_ranges)
				sum += count2(workflows, rule.next_name, new_ranges)
				new_ranges[rule.cat] = []int{rule.v, ranges[rule.cat][1]}
				// fmt.Println("2", name, rule, new_ranges)
			} else if rule.comp == ">" {
				new_ranges[rule.cat] = []int{rule.v + 1, ranges[rule.cat][1]}
				// fmt.Println("1", name, rule, new_ranges)
				sum += count2(workflows, rule.next_name, new_ranges)
				new_ranges[rule.cat] = []int{ranges[rule.cat][0], rule.v}
				// fmt.Println("2", name, rule, new_ranges)
			} else { // last rule
				sum += count2(workflows, rule.next_name, new_ranges)
			}
		}
		return sum
	}
}

func part2(input []string) int {
	workflows, _ := parse(input)
	ranges := map[string][]int{}
	for _, l := range []string{"x", "m", "a", "s"} {
		ranges[l] = []int{1, 4000}
	}
	return count2(workflows, "in", ranges)
}

func main() {
	file := utils.ReadFile("../inputs/19/ex2.txt")
	input := strings.Split(string(file), "\n\n")
	start, part1 := time.Now(), part1(input)
	fmt.Println("Part 1 :", part1, "- Time :", time.Since(start))
	start, part2 := time.Now(), part2(input)
	fmt.Println("Part 2 :", part2, "- Time :", time.Since(start))
}
