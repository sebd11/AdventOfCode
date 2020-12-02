package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := extractInputs()
	inputs[1] = 12
	inputs[2] = 2

	firstPartInputs := make([]int, len(inputs))
	copy(firstPartInputs, inputs)

	secondPartInputs := make([]int, len(inputs))
	copy(secondPartInputs, inputs)

	firstPart(firstPartInputs, true)
	secondPart(secondPartInputs)
}

func extractInputs() []int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	numbers := strings.Split(string(content), ",")
	var inputs = []int{}
	for _, value := range numbers {
		if value != "" {
			value = strings.TrimSuffix(value, "\n")
			input, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			inputs = append(inputs, input)
		}
	}
	return inputs
}

func firstPart(inputs []int, printResult bool) {
	for i := 0; i < len(inputs); i++ {
		switch inputs[i] {
		case 1:
			sum := inputs[inputs[i+1]] + inputs[inputs[i+2]]
			inputs[inputs[i+3]] = sum
			i += 3
		case 2:
			prod := inputs[inputs[i+1]] * inputs[inputs[i+2]]
			inputs[inputs[i+3]] = prod
			i += 3
		case 99:
			i = len(inputs)
		default:
			break
		}
	}
	if printResult {
		fmt.Printf("Part 1 result: %d\n", inputs[0])
	}
}

func secondPart(inputs []int) {
	tmp := make([]int, len(inputs))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(tmp, inputs)
			tmp[1] = noun
			tmp[2] = verb

			firstPart(tmp, false)
			if tmp[0] == 19690720 {
				fmt.Printf("Part 2 result: %d\n", 100*noun+verb)
				break
			}
		}
	}
}
