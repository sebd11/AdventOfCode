package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "input.txt"

type Input struct {
	op  string
	num int
}

func main() {
	fileContent, _ := ioutil.ReadFile("input.txt")
	inputs := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	acc, _ := firstPart(inputs)
	fmt.Printf("Part 1 - The value in the accumulator is: %d\n", acc)
	fmt.Printf("Part 2 - The value in the accumulator is: %d\n", secondPart(inputs))
}

func firstPart(inputs []string) (int, error) {
	acc := 0
	var err error
	seen := map[int]bool{}

	for i := 0; i < len(inputs); {
		if seen[i] {
			err = errors.New("Infinite Loop!")
			break
		}
		seen[i] = true
		var op string
		var num int

		fmt.Sscanf(inputs[i], "%s %d", &op, &num)

		switch op {
		case "acc":
			acc += num
		case "jmp":
			i += num - 1
		}
		i += 1
	}
	return acc, err
}

func secondPart(inputs []string) int {
	for i, input := range inputs {
		tmp := make([]string, len(inputs))
		copy(tmp, inputs)
		tmp[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(input)
		if acc, err := firstPart(tmp); err == nil {
			return acc
		}
	}
	return 0
}
