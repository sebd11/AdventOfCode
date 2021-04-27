package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const filename = "input.txt"

func main() {
	inputs := extractInputs()
	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)

	fmt.Printf("Part 1 - The produce is: %d\n", firstPart(inputs))

	fmt.Printf("Part 2 - The number of distinct ways is: %d\n", secondPart(0, inputs, make(map[int]int)))
}

func firstPart(inputs []int) int {
	var diffOne, diffThree int
	for i := range inputs {
		if i == 0 {
			continue
		}
		prev, curr := inputs[i-1], inputs[i]
		if prev+1 == curr {
			diffOne++
		}
		if prev+3 == curr {
			diffThree++
		}
	}
	return diffOne * diffThree
}

func secondPart(fromIndex int, nums []int, visited map[int]int) int {
	if fromIndex >= len(nums)-3 {
		return 1
	}

	num := nums[fromIndex]
	prevNum := num
	if fromIndex > 0 {
		prevNum = nums[fromIndex -1]
	}

	if res, ok := visited[num]; ok {
		return res
	}

	var count int
	for i := fromIndex + 1; i < fromIndex+4; i++ {
		n := nums[i]
		if areCompatible(num, n) {
			count += secondPart(i, nums, visited)
		}
	}
	visited[num] = count

	if num != prevNum {
		visited[num] += visited[prevNum]
	}
	return visited[num]
}

func areCompatible(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}

func extractInputs() []int {
	inputs := []int{}
	inputs = append(inputs, 0)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, num)
	}
	return inputs
}
