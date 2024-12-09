package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(input []byte) ([]int, []int) {
	var arr1 []int
	var arr2 []int
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		nums := strings.Fields(line)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}
	return arr1, arr2
}

func getDistances(arr1, arr2 []int) []int {
	var distances []int
	for len(arr1) > 0 {
		min1 := arr1[0]
		min2 := arr2[0]
		if min1 < min2 {
			distances = append(distances, min2-min1)
			arr1 = arr1[1:]
			arr2 = arr2[1:]
		} else {
			distances = append(distances, min1-min2)
			arr1 = arr1[1:]
			arr2 = arr2[1:]
		}
	}
	return distances
}

func main() {
	var result int
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	arr1, arr2 := parse(input)
	slices.Sort(arr1)
	slices.Sort(arr2)
	distances := getDistances(arr1, arr2)

	for _, distance := range distances {
		result += distance
	}
	fmt.Println("The total distance is: ", result)
}
