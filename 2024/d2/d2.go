package d2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


// Read the data into a ragged array; each value in the
// returned array is a row in the input file.
func readData() [][]int {
	file, err := os.Open("d2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [][]int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		row := make([]int, len(line))
		for col, item := range line {
			value, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			row[col] = value
		}
		result = append(result, row)
	}
	return result
}

func numUnsafe(row []int) int {
	// Single values are safe
	if len(row) < 2 {
		return 0
	}

	unsafe := 0
	lower := 1
	upper := 3

	diff := row[1] - row[0]
	if (diff < 0) {
		lower, upper = -1*upper, -1*lower
	}

	for i := range(len(row) - 1) {
		diff = row[i+1] - row[i]
		if (diff < lower) || (diff > upper) {
			unsafe += 1
		}
	}
	return unsafe
}

func isSafeWithRemoval(row []int, removals int) bool {
	// Single values are safe
	if len(row) < 2 {
		return true
	}

	if numUnsafe(row) == 0 {
		return true
	}

	if removals > 0 {
		for i := range(len(row)) {
			// Can't use append, as it mutates the input :/
			newRow := make([]int, len(row) - 1)
			copy(newRow[:i], row[:i])
			copy(newRow[i:], row[i+1:])

			if isSafeWithRemoval(newRow, removals - 1) {
				return true
			}
		}
	}
	return false
}

func numSafeWithRemoval(data [][]int, removals int) int {
	result := 0
	for _, row := range data {
		if isSafeWithRemoval(row, removals) {
			result += 1
		}
	}
	return result
}

func Run() {
	data := readData()

	fmt.Println("[d2.1] number of safe rows: ", numSafeWithRemoval(data, 0))
	fmt.Println("[d2.2] number of safe rows with removals: ", numSafeWithRemoval(data, 1))
}
