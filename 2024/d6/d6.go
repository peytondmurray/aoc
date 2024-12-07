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

type State struct {
	state [][]rune
	location Location
	direction Direction
	nNewSteps int
}

func (a *State) At(loc Location) rune {
	if a.isInBounds(loc) {
		return a.state[loc.y][loc.x]
	}
	return -1
}

func (a *State) Set(loc Location, value rune) {
	a.state[loc.y][loc.x] = value
}


func (a *State) isInBounds(loc Location) bool {
	return (loc.y >= 0 && loc.y < len(a.state)) && (loc.x >= 0 && loc.x < len(a.state[0]))
}

func (a *State) print() {
	for i, row := range a.state {
		fmt.Printf("%4d: %v\n", i, string(row))
	}
}

func (a *State) Evolve() {
	for {
		nextLoc := a.location.Add(a.direction)
		nextRune := a.At(nextLoc)
		if nextRune == '.' {
			a.Set(nextLoc, 'X')
			a.nNewSteps = a.nNewSteps + 1
			a.location = nextLoc
		} else if nextRune == '#' {
			a.direction = a.direction.Turn()
		} else if nextRune == 'X' {
			a.location = nextLoc
		} else if nextRune == -1 {
			// We're off the map; return
			a.Set(a.location, 'X')
			a.nNewSteps = a.nNewSteps + 1
			return
		}
	}
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

func findGuard(state [][]rune) Location {
	for i, row := range state {
		for j, char := range string(row) {
			if char == '^' {
				return Location{j, i}
			}
		}
	}
	return Location{-1, -1}
}

func evolve(init [][]rune) State {

	location := findGuard(init)
	if location.x == -1 {
		log.Fatal("Error finding guard.")
	}


	// Initialize the state of the guard;
	// first movement is up
	guardState := State{
		init,
		location,
		Up,
		0,
	}

	// Run the simulation
	guardState.Evolve()
	return guardState
}

func Run() {
	state := evolve(readData())

	fmt.Println("[d6.1] number visited: ", state.nNewSteps)
	// fmt.Println("[d6.2] sum of middle values of fixed updates: ", sumAllMiddles(fixed))
}
