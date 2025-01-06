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

type BigGrid struct {
	Grid
}

type Grid struct {
	arr [][]string
	pos Location
}

type IGrid interface {
	getarr() [][]string
	getpos() Location
}

func (a BigGrid) getarr() [][]string {
	return a.arr
}
func (a BigGrid) getpos() Location {
	return a.pos
}
func (a Grid) getarr() [][]string {
	return a.arr
}
func (a Grid) getpos() Location {
	return a.pos
}

// gpsValue Calculate the GPS value of the grid.
func gpsValue(a IGrid) int {
	result := 0
	for i, row := range a.getarr() {
		for j, val := range row {
			if val == "O" || val == "[" {
				result += 100*i + j
			}
		}
	}
	return result
}

// printGrid Print the grid to stdout.
func printGrid(a IGrid) {
	fmt.Println(a.getpos())
	for _, row := range a.getarr() {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
}


// evolve Evolve the grid using the given steps.
func (a *Grid) evolve(steps []string) {
	for _, step := range steps {
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
				if char == "#" {
					break
				}
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
				if char == "#" {
					break
				}
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
				if char == "#" {
					break
				}
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

// move Attempt to move the robot.
//
// dx : +1 to go right, -1 to go left
// dy : +1 to go down, -1 to go up
func (a *BigGrid) move(dx int, dy int) {
	if dy == 0 {
		// Horizontal moves are the same as they were in part 1
		i := 1
		for ;; i++ {
			char := a.arr[a.pos.y][a.pos.x + i*dx]
			if char == "#" {
				return
			} else if char == "." {
				break
			}
		}

		for j := i; j > 0; j-- {
			a.arr[a.pos.y][a.pos.x + j*dx] = a.arr[a.pos.y][a.pos.x + (j - 1)*dx]
		}
		a.arr[a.pos.y][a.pos.x] = "."
		a.pos.x += dx
		return
	} else {
		// Need to check the entire tree if it can be pushed
		// before attempting to move it; each leaf checking
		// its own children is not enough
		if a.canPush(a.pos, dy, make(map[Location]struct{})) {
			a.push(a.pos, dy, make(map[Location]struct{}))
			a.pos.y += dy
		}
	}
}

// push Push the robot at the current location.
// Only used for the y-direction because the x-direction is the same
// as in part 1.
func (a *BigGrid) push(at Location, dy int, visited map[Location]struct{}) {
	if _, exists := visited[at]; exists {
		return
	}

	visited[at] = struct{}{}
	switch char := a.arr[at.y][at.x]; char {
	case "#":
		// Should never happen
	case "[":
		a.push(Location{at.x, at.y+dy}, dy, visited)
		a.push(Location{at.x+1, at.y}, dy, visited)
		a.arr[at.y][at.x] = "."
		a.arr[at.y][at.x+1] = "."
		a.arr[at.y+dy][at.x] = "["
		a.arr[at.y+dy][at.x+1] = "]"
	case "]":
		a.push(Location{at.x, at.y+dy}, dy, visited)
		a.push(Location{at.x-1, at.y}, dy, visited)
		a.arr[at.y][at.x] = "."
		a.arr[at.y][at.x-1] = "."
		a.arr[at.y+dy][at.x-1] = "["
		a.arr[at.y+dy][at.x] = "]"
	case ".":
		// Just move the square from behind you forward
	case "@":
		a.push(Location{at.x, at.y+dy}, dy, visited)
		a.arr[at.y][at.x] = "."
		a.arr[at.y+dy][at.x] = "@"
		return
	}
}

// canPush Check whether the robot can move at the current location.
// Only used for the y-direction because the x-direction is the same
// as in part 1.
//
// at : Location to check whether a push can be made
// dy : +1 is a push downward, -1 is a push in the upward
// visited : A cache of visited locations; should be empty unless the call
// comes from canPush().
func (a *BigGrid) canPush(at Location, dy int, visited map[Location]struct{}) bool {
	if _, exists := visited[at]; exists {
		// If it can't be pushed, that'll be handled wherever this was
		// put into the visited cache
		return true
	}

	visited[at] = struct{}{}
	var canPush bool
	switch char := a.arr[at.y][at.x]; char {
	case "#":
		canPush = false
	case "[":
		canPush = a.canPush(Location{at.x, at.y+dy}, dy, visited)
		canPush = canPush && a.canPush(Location{at.x+1, at.y}, dy, visited)
	case "]":
		canPush = a.canPush(Location{at.x, at.y+dy}, dy, visited)
		canPush = canPush && a.canPush(Location{at.x-1, at.y}, dy, visited)
	case ".":
		canPush = true
	case "@":
		canPush = a.canPush(Location{at.x, at.y+dy}, dy, visited)
	}
	return canPush
}

// evolve Evolve the BigGrid forward in time. If progress == true,
// print the grid at each step (and at the end)
func (a *BigGrid) evolve(steps []string, progress bool) {
	for _, step := range steps {
		if progress {
			printGrid(a)
		}
		switch step {
		case "^":
			a.move(0, -1)
		case ">":
			a.move(1, 0)
		case "v":
			a.move(0, 1)
		case "<":
			a.move(-1, 0)
		}
	}
	if progress {
		printGrid(a)
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
	// file, err := os.Open("d15/input3")
	// file, err := os.Open("d15/input2")
	file, err := os.Open("d15/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return readGrid(scanner), readSteps(scanner)
}

// double Generate the part 2 grid from the part 1 grid
func (a *Grid) double() BigGrid {
	newArr := make([][]string, len(a.arr))
	var newLoc Location
	for i, row := range a.arr {
		newArr[i] = make([]string, 2*len(a.arr[0]))
		for j, char := range row {
			if char == "#" {
				newArr[i][2*j] = "#"
				newArr[i][2*j+1] = "#"
			} else if char == "O" {
				newArr[i][2*j] = "["
				newArr[i][2*j+1] = "]"
			} else if char == "." {
				newArr[i][2*j] = "."
				newArr[i][2*j+1] = "."
			} else if char == "@" {
				newArr[i][2*j] = "@"
				newArr[i][2*j+1] = "."
				newLoc = Location{2*j, i}
			}
		}
	}
	return BigGrid{Grid{newArr, newLoc}}
}

func Run() {
	grid, steps := readData()

	bgrid := grid.double()
	grid.evolve(steps)
	bgrid.evolve(steps, false)
	fmt.Println("[d15.1] gpsValue:", gpsValue(grid))
	fmt.Println("[d15.2] gpsValue:", gpsValue(bgrid))
}
