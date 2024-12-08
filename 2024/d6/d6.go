package d6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	iInc int
	jInc int
}

var (
	Up = Direction{-1, 0}
	Down = Direction{1, 0}
	Left = Direction{0, -1}
	Right = Direction{0, 1}
	None = Direction{0, 0}
)

type Location struct {
	x int
	y int
}

func (a Location) Add(direction Direction) Location {
	return Location{a.x + direction.jInc, a.y + direction.iInc}
}

func (a Direction) Turn() Direction {
	if a == Up {
		return Right
	} else if a == Right {
		return Down
	} else if a == Down {
		return Left
	} else if a == Left {
		return Up
	}
	return None
}

func newState(init [][]rune) State {
	obstructions := make(map[Location]struct{})
	var location Location
	for i, row := range init {
		for j, char := range string(row) {
			if char == '#' {
				obstructions[Location{j, i}] = struct{}{}
			} else if char == '^' {
				location = Location{j, i}
			}
		}
	}

	// Initialize the visited location list of the guard;
	// first movement is up
	visited := make(map[Location]map[Direction]struct{}, 0)
	visited[location] = make(map[Direction]struct{})
	addToSet(visited[location], Up)

	return State{
		obstructions,
		visited,
		location,
		Up,
		[]int{len(init), len(init[0])},
	}
}

func addToSet[T comparable](set map[T]struct{}, dir T) {
	set[dir] = struct{}{}
}

func setContains[T comparable](set map[T]struct{}, dir T) bool {
	_, contains := set[dir]
	return contains
}

type State struct {
	obstructions map[Location]struct{}
	visited map[Location]map[Direction]struct{}
	location Location
	direction Direction
	shape []int
}

func (a *State) isInBounds(loc Location) bool {
	return (loc.y >= 0 && loc.y < a.shape[0]) && (loc.x >= 0 && loc.x < a.shape[1])
}

// Evolve Walk the guard around until
//
// 1. They go off the grid (returns false)
// 2. They go in a loop (returns true)
func (a *State) Evolve() bool {
	for {
		nextLoc := a.location.Add(a.direction)
		if setContains(a.obstructions, nextLoc) {
			a.direction = a.direction.Turn()
		} else {
			directions, visited := a.visited[nextLoc]
			if visited && setContains(directions, a.direction) {
				// We've visited already because the next step is on
				// the same path in the same direction as before
				return true
			} else if a.isInBounds(nextLoc) {
				// If we're visiting a site in bounds, add the location
				// and current direction to the visited locations
				if a.visited[nextLoc] == nil {
					a.visited[nextLoc] = make(map[Direction]struct{})
				}
				addToSet(a.visited[nextLoc], a.direction)
				a.location = nextLoc
			} else {
				// If we've left the grid, exit
				return false
			}
		}
	}
}

func (a *State) getVisited() []Location {
	var result []Location
	for loc := range a.visited {
		result = append(result, loc)
	}
	return result
}

func readData() [][]rune {
	file, err := os.Open("d6/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, []rune(line))
	}

	return result
}

func evolve(init [][]rune) State {
	guardState := newState(init)
	guardState.Evolve()
	return guardState
}

func nObstructedLoops(init [][]rune) int {
	guardState := newState(init)
	guardState.Evolve()

	var loopLocs []Location

	visited := guardState.getVisited()
	for i, loc := range visited {
		fmt.Print("\033[G\033[K")
		fmt.Printf("Obstruction position: %5d/%5d\n", i, len(visited))
		fmt.Print("\033[A")

		obstructedState := newState(init)
		addToSet(obstructedState.obstructions, loc)
		if obstructedState.Evolve() {
			loopLocs = append(loopLocs, loc)
		}
	}

	return len(loopLocs)
}

func Run() {
	init := readData()
	state := evolve(init)
	nLoops := nObstructedLoops(init)

	fmt.Println("[d6.1] number visited: ", len(state.visited))
	fmt.Println("[d6.2] number of obstructed loops: ", nLoops)
}
