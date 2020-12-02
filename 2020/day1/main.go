package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	var inputs = []int{}
	for _, value := range lines {
		if value != "" {
			input, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			inputs = append(inputs, input)
		}
	}

	for i, value := range inputs {
		for j, value2 := range inputs {
			// Part 1
			if i != j {
				if value+value2 == 2020 {
					fmt.Printf("Answer for part 1 is: %d, ", value*value2)
				}
			}
			// Part 2
			for k, value3 := range inputs {
				if i != j && i != k && j != k {
					if value+value2+value3 == 2020 {
						fmt.Printf("Answer for part 2 is: %d, ", value*value2*value3)
					}
				}
			}
		}
	}
}
