package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	totalFuel := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		for input > 0 {
			input = int(input/3) - 2
			if input > 0 {
				totalFuel += input
			}
		}
	}
	fmt.Printf("Rounded number: %d", totalFuel)
}
