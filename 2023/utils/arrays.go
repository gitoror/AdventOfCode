package utils

import (
	"strconv"
)

func SliceInt(slice []string) []int {
	var slice_int = []int{}
	for _, i := range slice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		slice_int = append(slice_int, j)
	}
	return slice_int
}

func SliceSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
