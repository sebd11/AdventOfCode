package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "input.txt"

func main() {
	inputs := extractInputs()
	firstPart(inputs, 3, 1)
	fmt.Printf("Part 1 - Tree count: %d\n", firstPart(inputs, 3, 1))
	fmt.Printf("Part 2 - Answer: %d\n",
		firstPart(inputs, 1, 1)*firstPart(inputs, 3, 1)*firstPart(inputs, 5, 1)*
			firstPart(inputs, 7, 1)*firstPart(inputs, 1, 2))
}

func extractInputs() [][]string {
	var inputs = [][]string{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		inputs = append(inputs, line)
	}
	return inputs
}

func firstPart(inputs [][]string, right int, down int) int {
	treeCount := 0

	for i := down; i < len(inputs); {
		for j := right; j < len(inputs[1]); {
			if inputs[i][j] == "#" {
				treeCount += 1
			}
			j += right
			i += down
			if j >= len(inputs[1]) {
				j -= len(inputs[1])
			}
			if i >= len(inputs) {
				break
			}
		}
	}
	return treeCount
}
