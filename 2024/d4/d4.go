package d4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData() []string {
	text, err := os.ReadFile("d4/input")
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Split(string(text), "\n")
	return result[:len(result)-1]
}

// 	pattern := [][]rune {
// 		{'S', '.', '.', 'S', '.', '.', 'S'},
// 		{'.', 'A', '.', 'A', '.', 'A', '.'},
// 		{'.', '.', 'M', 'M', 'M', '.', '.'},
// 		{'S', 'A', 'M', 'X', 'M', 'A', 'S'},
// 		{'.', '.', 'M', 'M', 'M', '.', '.'},
// 		{'.', 'A', '.', 'A', '.', 'A', '.'},
// 		{'S', '.', '.', 'S', '.', '.', 'S'},
// 	}

func isMas(a byte, b byte, c byte) bool {
	return (a == 'M') && (b == 'A') && (c == 'S')
}
func isCrossed(upleft byte, upright byte, downright byte, downleft byte) bool {
	uldr := ((upleft == 'M') && (downright == 'S') || (upleft == 'S') && (downright == 'M'))
	urdl := ((upright == 'M') && (downleft == 'S')) || ((upright == 'S') && (downleft == 'M'))
	return uldr && urdl
}

// printTarget Print the entire dataset with the character
// under examination highlighted in red
func printTarget(data []string, i int, j int) {
	reset := "\033[0m"
	red := "\033[31m"

	fmt.Printf("\n(%d, %d)\n", i, j)
	for ii, line := range data {
		if ii == i {
			fmt.Printf("%3d: %v%v%v%v%v\n", ii, line[:j], red, string(line[j]), reset, line[j+1:])
		} else {
			fmt.Printf("%3d: %v\n", ii, line)
		}
	}
}

func nCrossed(data []string, i int, j int) int {
	result := 0

	up := i >= 1
	down := i <= len(data)-2
	left := j >= 1
	right := j <= len(data[i])-2

	if up && down && left && right {
		if isCrossed(data[i-1][j-1], data[i-1][j+1], data[i+1][j+1], data[i+1][j-1]) {
			result++
		}
	}
	return result
}
func nSurrounding(data []string, i int, j int) int {
	result := 0

	up := i >= 3
	down := i <= len(data)-4
	left := j >= 3
	right := j <= len(data[i])-4

	if up {
		// up
		if isMas(data[i-1][j], data[i-2][j], data[i-3][j]) {
			result += 1
		}

		// up left
		if left && isMas(data[i-1][j-1], data[i-2][j-2], data[i-3][j-3]) {
			result += 1
		}

		// up right
		if right && isMas(data[i-1][j+1], data[i-2][j+2], data[i-3][j+3]) {
			result += 1
		}
	}
	if down {
		// down
		if isMas(data[i+1][j], data[i+2][j], data[i+3][j]) {
			result += 1
		}

		// down left
		if left && isMas(data[i+1][j-1], data[i+2][j-2], data[i+3][j-3]) {
			result += 1
		}

		// down right
		if right && isMas(data[i+1][j+1], data[i+2][j+2], data[i+3][j+3]) {
			result += 1
		}
	}

	// left
	if left && isMas(data[i][j-1], data[i][j-2], data[i][j-3]) {
		result += 1
	}
	// right
	if right && isMas(data[i][j+1], data[i][j+2], data[i][j+3]) {
		result += 1
	}

	return result
}

func findAllCrossed(data []string) int {
	result := 0
	for i, line := range data {
		for j, char := range line {
			if char == 'A' {
				result += nCrossed(data, i, j)
			}
		}
	}
	return result
}

func findAllXmas(data []string) int {
	result := 0
	for i, line := range data {
		for j, char := range line {
			if char == 'X' {
				result += nSurrounding(data, i, j)
			}
		}
	}
	return result
}

func Run() {
	data := readData()

	fmt.Println("[d4.1] nXmas: ", findAllXmas(data))
	fmt.Println("[d4.2] nCrossed: ", findAllCrossed(data))
}
