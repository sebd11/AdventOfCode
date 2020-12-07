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
	firstPart(inputs)
	secondPart(inputs)
}

func firstPart(inputs []string) {
	sum := 0

	for _, input := range inputs {
		m := map[string]int{}
		substrings := strings.Split(input, ",")
		letters := strings.Split(strings.Join(substrings, ""), "")
		for _, letter := range letters {
			m[letter] = 1
		}
		sum += len(m)
	}
	fmt.Printf("Part 1 - The sum is: %d\n", sum)
}

func secondPart(inputs []string) {
	sum := 0

	for _, input := range inputs {
		m := map[string]int{}
		substrings := strings.Split(input, ",")
		if substrings[len(substrings)-1] == "" {
			substrings = substrings[:len(substrings)-1]
		}
		letters := strings.Split(strings.Join(substrings, ""), "")
		for _, letter := range letters {
			if _, ok := m[letter]; ok {
				m[letter] += 1
			} else {
				m[letter] = 1
			}
		}
		for _, v := range m {
			if v == len(substrings) {
				sum += 1
			}
		}
	}

	fmt.Printf("Part 2 - The sum is: %d\n", sum)
}

func extractInputs() []string {
	inputs := []string{}
	var sb strings.Builder

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputs = append(inputs, strings.TrimSuffix(sb.String(), ","))
			sb.Reset()
		} else {
			sb.WriteString(line + ",")
		}
	}
	inputs = append(inputs, sb.String())

	return inputs
}
