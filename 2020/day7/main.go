package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"
const shinyBag = "shiny gold"

type bagCount struct {
	color string
	qty   int
}

func main() {
	inputs := extractInputs()
	contains, containedBy := createMaps(inputs)
	firstPart(containedBy)
	secondPart(contains)
}

func firstPart(containedBy map[string][]string) {
	canContain := containedBy[shinyBag]
	seen := make(map[string]bool)
	seen[shinyBag] = true
	for len(canContain) > 0 {
		curr := canContain[0]
		canContain = canContain[1:]
		if seen[curr] {
			continue
		}
		seen[curr] = true
		canContain = append(canContain, containedBy[curr]...)
	}

	fmt.Printf("Part 1 - The number of bag: %d\n", len(seen)-1)
}

func secondPart(contains map[string][]bagCount) {
	fmt.Printf("Part 2 - The sum is: %d\n", countContent(shinyBag, contains))
}

func cleanBagString(bag *string) {
	*bag = strings.TrimSuffix(*bag, " bags")
	*bag = strings.TrimSuffix(*bag, " bag")
	*bag = strings.TrimSuffix(*bag, " bags.")
	*bag = strings.TrimSuffix(*bag, " bag.")
}

func countContent(color string, contains map[string][]bagCount) int {
	sum := 0
	for _, bag := range contains[color] {
		sum += bag.qty
		sum += bag.qty * countContent(bag.color, contains)
	}
	return sum
}

func createMaps(inputs []string) (map[string][]bagCount, map[string][]string) {
	contains, containedBy := make(map[string][]bagCount), make(map[string][]string)

	for _, input := range inputs {
		subStrings := strings.Split(input, ", ")
		mainBag := strings.Split(subStrings[0], " contain ")
		container := mainBag[0]
		cleanBagString(&container)

		for index, ss := range subStrings {
			if index == 0 {
				ss = mainBag[1]
			}
			cleanBagString(&ss)
			if ss != "no other" {
				bag := bagCount{}
				content := strings.Split(ss, " ")
				qty, _ := strconv.Atoi(content[0])
				color := content[1] + " " + content[2]
				bag.color = color
				bag.qty = qty
				contains[container] = append(contains[container], bag)
				containedBy[bag.color] = append(containedBy[bag.color], container)
			}
		}
	}
	return contains, containedBy
}

func extractInputs() []string {
	inputs := []string{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	return inputs
}
