package d16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Location struct {
	x int
	y int
}

type Direction struct {
	dx int
	dy int
}

var East = Direction{1, 0}
var South = Direction{0, 1}
var West = Direction{-1, 0}
var North = Direction{0, -1}

func (a *Direction) right() Direction {
	switch *a {
		case East:
			return South
		case South:
			return West
		case West:
			return North
		default:
			return East
	}
}

func (a *Direction) left() Direction {
	switch *a {
		case East:
			return North
		case North:
			return West
		case West:
			return South
		default:
			return East
	}
}

func (a *Location) step(b Direction) Location {
	return Location{a.x + b.dx, a.y + b.dy}
}

type LocDir struct {
	loc Location
	dir Direction
}


type Grid struct {
	arr [][]string
	start Location
	end Location
}

func (a *Grid) at(loc Location) string {
	return a.arr[loc.y][loc.x]
}

func newGrid(arr [][]string) Grid {
	var start Location
	var end Location
	for i, row := range arr {
		for j, char := range row {
			if char == "E" {
				end = Location{j, i}
			} else if char == "S" {
				start = Location{j, i}
			}
		}
	}
	return Grid{arr, start, end}
}

func readData() Grid {
	// file, err := os.Open("d16/input2")
	file, err := os.Open("d16/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr [][]string
	for scanner.Scan() {
		text := scanner.Text()

		arr = append(arr, strings.Split(text, ""))
		if len(arr) == 0 {
			return newGrid(arr)
		}
	}

	return newGrid(arr)
}

func (a *Grid) solve(loc Location, dir Direction, visited map[LocDir]int, currentCost int, paths map[Location]struct{}) (int, bool) {

	// Reached the end
	if loc == a.end {
		paths[loc] = struct{}{}
		return currentCost, true
	}

	// If you've come upon the same location and direction,
	// don't bother continuing if the cost was lower in the other code path
	state := LocDir{loc, dir}
	if cachedCost, exists := visited[state]; exists && cachedCost < currentCost {
		return 0, false
	}


	// Store the cost of getting to the current location
	visited[state] = currentCost

	var possiblePaths []map[Location]struct{}

	var costs []int
	leftDir := dir.left()
	leftLoc := loc.step(leftDir)
	if a.at(leftLoc) != "#" {
		nextCost := 1000 + currentCost
		cachedCost, exists := visited[LocDir{loc, leftDir}]
		if !exists || cachedCost >= nextCost {

			restPath := make(map[Location]struct{})
			restCost, okay := a.solve(loc, leftDir, visited, nextCost, restPath)
			if okay {
				costs = append(costs, restCost)
				possiblePaths = append(possiblePaths, restPath)
			}
		}
	}

	rightDir := dir.right()
	rightLoc := loc.step(rightDir)
	if a.at(rightLoc) != "#" {
		nextCost := 1000 + currentCost
		cachedCost, exists := visited[LocDir{loc, rightDir}]
		if !exists || cachedCost >= nextCost {

			restPath := make(map[Location]struct{})
			restCost, okay := a.solve(loc, rightDir, visited, nextCost, restPath)
			if okay {
				costs = append(costs, restCost)
				possiblePaths = append(possiblePaths, restPath)
			}
		}
	}

	forwardLoc := loc.step(dir)
	if a.at(forwardLoc) != "#" {
		nextCost := 1 + currentCost
		cachedCost, exists := visited[LocDir{forwardLoc, dir}]
		if !exists || cachedCost >= nextCost {

			restPath := make(map[Location]struct{})
			restCost, okay := a.solve(forwardLoc, dir, visited, nextCost, restPath)
			if okay {
				costs = append(costs, restCost)
				possiblePaths = append(possiblePaths, restPath)
			}
		}
	}

	if len(costs) == 0 {
		return 0, false
	}

	// Copy the minimum paths into the paths map
	for _, i := range idxmin(costs) {
		for location := range possiblePaths[i] {
			paths[location] = struct{}{}
		}
	}
	// Include the current position as well
	paths[loc] = struct{}{}
	return slices.Min(costs), true
}

func idxmin(arr []int) []int {
	minimum := arr[0]
	iMin := []int{0}

	for i, value := range arr {
		if value < minimum {
			minimum = value
			iMin = []int{i}
		} else if value == minimum {
			iMin = append(iMin, i)
		}
	}

	return iMin
}

func (a *Grid) print(paths map[Location]struct{}) {

	// Copy the grid into a new array so that we can draw the path on it
	arr := make([][]string, len(a.arr))
	for i, row := range a.arr {
		arr[i] = make([]string, len(row))
		for j, char := range row {
			arr[i][j] = char
		}
	}

	// Iterate through the locations visited and insert them into the array
	for loc := range paths {
		arr[loc.y][loc.x] = "O"
	}

	// Print the array
	fmt.Println()
	for _, row := range arr {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func Run() {
	grid := readData()

	paths := make(map[Location]struct{})
	minCost, _ := grid.solve(grid.start, East, make(map[LocDir]int), 0, paths)
	grid.print(paths)
	fmt.Println("[d16.1] min cost:", minCost)
	fmt.Println("[d16.2] optimal seats:", len(paths))
}
