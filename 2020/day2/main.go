package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"

type Input struct {
	minOcc   int
	maxOcc   int
	letter   string
	password string
}

func main() {
	inputs := extractPasswords()
	firstPart(inputs)
	secondPart(inputs)
}

func firstPart(inputs []Input) {
	validPasswords := 0

	for _, input := range inputs {
		count := strings.Count(input.password, input.letter)
		if count >= input.minOcc && count <= input.maxOcc {
			validPasswords += 1
		}
	}
	fmt.Printf("Part 1 - Number of valid passwords: %d\n", validPasswords)
}

func secondPart(inputs []Input) {
	validPasswords := 0

	for _, input := range inputs {
		password := strings.Split(input.password, "")
		isValid := password[input.minOcc-1] == input.letter && password[input.maxOcc-1] != input.letter ||
			password[input.minOcc-1] != input.letter && password[input.maxOcc-1] == input.letter
		if isValid {
			validPasswords += 1
		}
	}
	fmt.Printf("Part 2 - Number of valid passwords: %d\n", validPasswords)
}

func extractPasswords() []Input {
	var inputs = []Input{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		occ := strings.Split(line[0], "-")
		minOcc, _ := strconv.Atoi(occ[0])
		maxOcc, _ := strconv.Atoi(occ[1])
		input := Input{
			minOcc:   minOcc,
			maxOcc:   maxOcc,
			letter:   strings.TrimSuffix(line[1], ":"),
			password: line[2],
		}
		inputs = append(inputs, input)
	}
	return inputs
}
