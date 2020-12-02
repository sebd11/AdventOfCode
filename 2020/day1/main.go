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
	firstResult := 0
	secondResult := 0

	for i, v := range inputs {
		for j, vv := range inputs {
			// Part 1
			if i != j {
				if v+vv == 2020 {
					firstResult = v * vv
				}
			}
			// Part 2
			for k, vvv := range inputs {
				if i != j && i != k && j != k {
					if v+vv+vvv == 2020 {
						secondResult = v * vv * vvv
					}
				}
			}
		}
	}
	fmt.Printf("Answer for part 1 is: %d\n", firstResult)
	fmt.Printf("Answer for part 2 is: %d\n", secondResult)
}

func extractInputs() []int {
	var inputs = []int{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, input)
	}
	return inputs
}
