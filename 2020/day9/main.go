package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const filename = "input.txt"

func main() {
	inputs := extractInputs()
	num := firstPart(inputs, 25)
	fmt.Printf("Part 1 - The invalid number is: %d\n", num)
	sum := secondPart(inputs, num)
	fmt.Printf("Part 2 - The sum of min and max values is: %d\n", sum)
}

func firstPart(inputs []int, preamble int) int {
out:
	for i := preamble; i < len(inputs); i++ {
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if inputs[k]+inputs[j] == inputs[i] {
					continue out
				}
			}
		}
		return inputs[i]
	}
	return 0
}

func secondPart(inputs []int, invalidNum int) int {
	for i := 0; i < len(inputs); i++ {
		sum := inputs[i]
		max := inputs[i]
		min := inputs[i]

		for j := i + 1; j < len(inputs); j++ {
			sum += inputs[j]
			if inputs[j] > max {
				max = inputs[j]
			}
			if inputs[j] < min {
				min = inputs[j]
			}
			if sum > invalidNum {
				break
			}
			if sum == invalidNum {
				return min + max
			}
		}
	}

	return 0
}

func extractInputs() []int {
	inputs := []int{}

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
