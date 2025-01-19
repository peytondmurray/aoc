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
	file, err := os.Open("d16/input2")
	// file, err := os.Open("d16/input")
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

func (a *Grid) solve(loc Location, dir Direction, visited map[Location]struct{}) int {

	// Reached the end
	if loc.x == a.end.x && loc.y == a.end.y {
		return 0
	}

	visited[loc] = struct{}{}

	var costs []int
	leftLoc := loc.step(dir.left())
	if a.at(leftLoc) != "#" {
		if _, exists := visited[leftLoc]; !exists {
				costs = append(costs, 1000 + a.solve(loc, dir.left(), visited))
		}
	}

	rightLoc := loc.step(dir.right())
	if a.at(rightLoc) != "#" {
		if _, exists := visited[rightLoc]; !exists {
			costs = append(costs, 1000 + a.solve(loc, dir.right(), visited))
		}
	}

	forwardLoc := loc.step(dir)
	if a.at(forwardLoc) != "#" {
		if _, exists := visited[forwardLoc]; !exists {
			costs = append(costs, 1 + a.solve(loc.step(dir), dir, visited))
		}
	}

	return slices.Min(costs)
}

func Run() {
	grid := readData()

	fmt.Println("[d16.1] min cost:", grid.solve(grid.start, East, make(map[Location]struct{})))
}
