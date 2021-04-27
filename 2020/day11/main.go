package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "input.txt"
const floor = '.'
const empty = 'L'
const occupied = '#'

func main() {
	seats := extractSeats()
	fmt.Printf("First part - There is %d seats occupied.\n", firstPart(seats))
	seats = extractSeats()
	fmt.Printf("Second part - There is %d seats occupied.\n", secondPart(seats))
}

func firstPart(seats [][]rune) int {
	prevSeats := copyArray(seats)
	maxRow := len(prevSeats) - 1
	maxCol := len(prevSeats[0]) - 1
	changed := false

	for i := 1; i < maxRow; i++ {
		for j := 1; j < maxCol; j++ {
			if prevSeats[i][j] == empty && verifyAdjacentNeighbors(i, j, 8, prevSeats, []rune {floor, empty}) {
				seats[i][j] = occupied
				changed = true
			}
			if prevSeats[i][j] == occupied && verifyAdjacentNeighbors(i, j, 4, prevSeats, []rune {occupied}) {
				seats[i][j] = empty
				changed = true
			}
		}
	}

	if changed {
		firstPart(seats)
	}
	return calculateOccupiedSeats(maxRow, maxCol, seats)
}

func secondPart(seats [][]rune) int {
	prevSeats := copyArray(seats)
	maxRow := len(prevSeats) - 1
	maxCol := len(prevSeats[0]) - 1
	changed := false

	for i := 1; i < maxRow; i++ {
		for j := 1; j < maxCol; j++ {
			if prevSeats[i][j] == empty && verifyVisibleNeighbors(i, j, 8, prevSeats, []rune {empty, floor}) {
				seats[i][j] = occupied
				changed = true
			}
			if prevSeats[i][j] == occupied && verifyVisibleNeighbors(i, j, 5, prevSeats, []rune {occupied}) {
				seats[i][j] = empty
				changed = true
			}
		}
	}

	if changed {
		secondPart(seats)
	}
	return calculateOccupiedSeats(maxRow, maxCol, seats)
}

func calculateOccupiedSeats(maxRow, maxCol int, seats [][]rune) int {
	seatsOccupied := 0
	for i := 1; i < maxRow; i++ {
		for j := 1; j < maxCol; j++ {
			if seats[i][j] == occupied {
				seatsOccupied++
			}
		}
	}

	return seatsOccupied
}

func verifyAdjacentNeighbors(i, j, minOcc int, seats [][]rune, validStates []rune) bool {
	validSeats := 0

	for row := i-1; row < i+2; row++ {
		for col :=j-1; col < j+2; col++ {
			if row == i && col == j {
				continue
			}
			if runeInArray(seats[row][col], validStates) {
				validSeats++
			}
		}
	}
	return validSeats >= minOcc
}

func verifyVisibleNeighbors(i, j, minOcc int, seats [][]rune, validStates []rune) bool {
	validSeats := 0
	
	for row := i-1; row < i+2; row++ {
		for col :=j-1; col < j+2; col++ {
			if row == i && col == j {
				continue
			}

			seat := findClosestNeighbor(i, j, row-i, col-j, seats)
			if runeInArray(seat, validStates) {
				validSeats++
			}
		}
	}
	return validSeats >= minOcc
}

func findClosestNeighbor(i, j, modI, modJ int, seats [][]rune) rune{
	row := i + modI
	col := j + modJ

	for seats[row][col] == floor && !isOutOfBound(row, col, len(seats)-2, len(seats[0])-2) {
		row += modI
		col += modJ
	}
	return seats[row][col]
}

func isOutOfBound(i,j, maxRow, maxCol int) bool {
	return i < 1 || i > maxRow || j < 1 || j > maxCol
}

// Returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Returns if a rune is in the list of runes
func runeInArray(a rune, list []rune) bool {
	for _, b := range list {
        if b == a {
            return true
        }
    }
	return false
}

// Returns a copy of a 2D rune array
func copyArray(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))

	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func extractSeats() [][]rune {
	seats := [][]rune{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		// Adds a column of floor at the beginning and the end
		runes = append([]rune {floor}, runes...)
		runes = append(runes, floor)

		seats = append(seats, runes)
	}
	lineLength := len(seats[0])
	border := []rune {}
	for i :=0; i< lineLength; i++ {
		border = append(border, floor)
	}

	// Adds a row of floor at the beginning and the end
	seats = append([][]rune {border}, seats...)
	seats = append(seats, border)
	return seats
}