package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := load()
	if err != nil {
		log.Fatalf("fatal error loading file: %v", err)
	}

	// convert to int slice for convenience
	nums := make([]int, len(input))
	for i, e := range input {
		n, _  := strconv.Atoi(e) // ignoring error
		nums[i] = n
	}

	n1, n2, err := twoSum(nums, 2020)
	if err != nil {
		log.Fatalf("error in part 1: %v", err)
	}
	fmt.Printf("%d * %d = %d\n", n1, n2, n1 * n2)

	n1, n2, n3, err := threeSum(nums, 2020)
	if err != nil {
		log.Fatalf("error in part 1: %v", err)
	}
	fmt.Printf("%d * %d * %d = %d\n", n1, n2, n3, n1 * n2 * n3)
}

// part 1: use hash set
func twoSum(input []int, target int) (int, int, error) {
	var n2 int
	inputSet := make(map[int]struct{}, len(input))
	for _, n1 := range input {
		n2 = target - n1
		if _, ok := inputSet[n2]; ok {
			return n1, n2, nil
		}
		inputSet[n1] = struct{}{}
	}

	return 0, 0, fmt.Errorf("no solution found")
}

// part 2: same as part 1 but nested loops
func threeSum(input []int, target int) (int, int, int, error) {
	inputSet := make(map[int]struct{}, len(input))
	var newTarget, n3 int
	for i1, n1 := range input {
		for i2, n2 := range input {
			if i1 == i2 {
				continue
			}
			newTarget = target - n1
			n3 = newTarget - n2
			if _, ok := inputSet[n3]; ok {
				return n1, n2, n3, nil
			}
			inputSet[n2] = struct{}{}
		}
		inputSet = make(map[int]struct{}, len(input))
	}

	return 0, 0, 0, fmt.Errorf("no solution found")
}


func load() ([]string, error) {
	inputFile, err := os.Open("day1/day1.txt")
	if err != nil {
		return nil, err
	}

	input, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(input), "\n"), nil
}