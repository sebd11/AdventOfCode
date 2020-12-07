package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const filename = "input.txt"

func main() {
	seats := extractInputs()
	sortedSeats := firstPart(seats)
	secondPart(sortedSeats)
}

func firstPart(seats []string) []int {
	sortedSeats := []int{}
	for _, seat := range seats {
		row, col := getRowAndCol(seat)
		seatID := getSeatId(row, col)
		sortedSeats = insertSorted(sortedSeats, seatID)
	}
	fmt.Printf("Part 1 - Highest seat ID is: %d\n", sortedSeats[len(sortedSeats)-1])
	return sortedSeats
}

func secondPart(seats []int) {
	mySeat := 0
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] == 2 {
			mySeat = seats[i] + 1
			break
		}
	}
	fmt.Printf("Part 2 - My seat ID is: %d\n", mySeat)
}

func insertSorted(s []int, e int) []int {
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func getRowAndCol(seat string) (int, int) {
	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7

	letters := strings.Split(seat, "")

	for i := 0; i < 7; i++ {
		diff := maxRow - minRow
		if letters[i] == "F" {
			maxRow -= diff/2 + 1
		} else {
			minRow += diff/2 + 1
		}
		if i == 6 && letters[i] == "B" {
			maxRow = minRow
		}
	}

	for i := 7; i < len(letters); i++ {
		diff := maxCol - minCol
		if letters[i] == "L" {
			maxCol -= diff/2 + 1
		} else {
			minCol += diff/2 + 1
		}
		if i == len(letters)-1 && letters[i] == "R" {
			maxCol = minCol
		}
	}

	return maxRow, maxCol
}

func getSeatId(row int, col int) int {
	return row*8 + col
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
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}
