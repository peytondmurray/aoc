package d15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Location struct {
	x int
	y int
}

func (a *Location) add(dx int, dy int) Location {
	return Location{a.x + dx, a.y + dy}
}

type Grid struct {
	arr [][]string
	pos Location
}

func (a *Grid) gpsValue() int {
	result := 0
	for i, row := range a.arr {
		for j, val := range row {
			if val == "O" {
				result += 100*i + j
			}
		}
	}
	return result
}

func (a *Grid) print() {
	for _, row := range a.arr {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
}


func (a *Grid) evolve(steps []string) {
	for _, step := range steps {
		// a.print()
		if step == "^" {
			for j := 1;; j++ {
				char := a.arr[a.pos.y - j][a.pos.x]

				// If we reach a wall, don't move anything
				if char == "#" {
					break
				}

				// If we reach an empty space, push the stack of boxes forward one square
				if char == "." {
					a.arr[a.pos.y - j][a.pos.x] = a.arr[a.pos.y - j + 1][a.pos.x]
					a.arr[a.pos.y - 1][a.pos.x] = "@"
					a.arr[a.pos.y][a.pos.x] = "."
					a.pos = a.pos.add(0, -1)
					break
				}
			}
		} else if step == "v" {
			for j := 1;; j++ {
				char := a.arr[a.pos.y + j][a.pos.x]

				// If we reach a wall, don't move anything
				if char == "#" {
					break
				}

				// If we reach an empty space, push the stack of boxes forward one square
				if char == "." {
					a.arr[a.pos.y + j][a.pos.x] = a.arr[a.pos.y + j - 1][a.pos.x]
					a.arr[a.pos.y + 1][a.pos.x] = "@"
					a.arr[a.pos.y][a.pos.x] = "."
					a.pos = a.pos.add(0, 1)
					break
				}
			}
		} else if step == ">" {
			for j := 1;; j++ {
				char := a.arr[a.pos.y][a.pos.x + j]

				// If we reach a wall, don't move anything
				if char == "#" {
					break
				}

				// If we reach an empty space, push the stack of boxes forward one square
				if char == "." {
					a.arr[a.pos.y][a.pos.x + j] = a.arr[a.pos.y][a.pos.x + j - 1]
					a.arr[a.pos.y][a.pos.x + 1] = "@"
					a.arr[a.pos.y][a.pos.x] = "."
					a.pos = a.pos.add(1, 0)
					break
				}
			}
		} else if step == "<" {
			for j := 1;; j++ {
				char := a.arr[a.pos.y][a.pos.x - j]

				// If we reach a wall, don't move anything
				if char == "#" {
					break
				}

				// If we reach an empty space, push the stack of boxes forward one square
				if char == "." {
					a.arr[a.pos.y][a.pos.x - j] = a.arr[a.pos.y][a.pos.x - j + 1]
					a.arr[a.pos.y][a.pos.x - 1] = "@"
					a.arr[a.pos.y][a.pos.x] = "."
					a.pos = a.pos.add(-1, 0)
					break
				}
			}
		}
	}
}

func readGrid(scanner *bufio.Scanner) Grid {
	var arr [][]string
	var loc Location
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			return Grid{arr, loc}
		}

		arr = append(arr, strings.Split(text, ""))
		start := strings.Index(text, "@")
		if start != -1 {
			loc = Location{i, start}
		}
		i++
	}

	return Grid{arr, loc}
}

func readSteps(scanner *bufio.Scanner) []string {
	var steps []string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			return steps
		}
		steps = append(steps, strings.Split(text, "")...)
	}
	return steps
}

func readData() (Grid, []string) {
	// file, err := os.Open("d15/input2")
	file, err := os.Open("d15/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return readGrid(scanner), readSteps(scanner)
}

func Run() {
	grid, steps := readData()
	grid.evolve(steps)

	fmt.Println("[d15.1] gpsValue:", grid.gpsValue())
	// fmt.Println("[d13.2] tokens:", solve(data, 10000000000000))
}
