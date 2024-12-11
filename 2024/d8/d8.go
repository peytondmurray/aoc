package d8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Add an item to the set
func addToSet[T comparable](set map[T]struct{}, item T) {
	set[item] = struct{}{}
}

// setToSlice Convert a set to a slice
func setToSlice[T comparable](set map[T]struct{}) []T {
	var result []T
	for item := range set {
		result = append(result, item)
	}
	return result
}

type Grid struct {
	locations map[rune][]Location
	shape [2]int // ysize, xsize
}

type Location struct {
	x int
	y int
}

func (a *Location) subtract(x int, y int) Location {
	return Location{a.x - x, a.y - y}
}

func (a Location) distanceTo(b Location) (int, int) {
	return b.x - a.x, b.y - a.y
}

// pairs Generate a slice containing every pair of inputs
func pairs(locs []Location) [][2]Location {
	var results [][2]Location
	for i := 0; i < len(locs); i++ {
		for j := i+1; j < len(locs); j++ {
			results = append(results, [2]Location{locs[i], locs[j]})
		}
	}

	return results
}

// isInBounds True if the location is in bounds, false otherwise
func (a *Grid) isInBounds(loc Location) bool {
	return loc.x >= 0 && loc.x < a.shape[1] && loc.y >= 0 && loc.y < a.shape[0]
}

// findAllAntinodes Find all antinodes for every antenna on the grid
func (a *Grid) findAllAntinodes() []Location {
	nodes := map[Location]struct{}{}

	for char := range a.locations {
		for _, loc := range a.findAntinodes(char) {
			if _, exists := nodes[loc]; !exists {
				nodes[loc] = struct{}{}
			}
		}
	}

	return setToSlice(nodes)
}


// findAntinodes Find all antinodes for the given characters
func (a *Grid) findAntinodes(chars ...rune) []Location {
	nodes := map[Location]struct{}{}

	for _, char := range chars {
		locations := a.locations[char]
		for _, pair := range pairs(locations) {

			distX, distY := pair[0].distanceTo(pair[1])
			for i := 0;; i++ {
				location := pair[0].subtract(i*distX, i*distY)
				if !a.isInBounds(location) {
					break
				}
				addToSet(nodes, location)
			}
			for i := -1;; i-- {
				location := pair[0].subtract(i*distX, i*distY)
				if !a.isInBounds(location) {
					break
				}
				addToSet(nodes, location)
			}
		}
	}

	return setToSlice(nodes)
}

// findAllNearestAntinodes Find the nearest antinodes for every
// antenna on the grid
func (a *Grid) findAllNearestAntinodes() []Location {
	// Use a hashmap cause we don't care about duplicates
	nodes := map[Location]struct{}{}
	for char := range a.locations {
		for _, loc := range a.findNearestAntinodes(char) {
			if _, exists := nodes[loc]; !exists {
				nodes[loc] = struct{}{}
			}
		}
	}

	return setToSlice(nodes)
}

// findNearestAntinodes Find the nearest antinodes for the given characters (part 1)
func (a *Grid) findNearestAntinodes(chars ...rune) []Location {
	var result []Location

	for _, char := range chars {
		locations := a.locations[char]
		for _, pair := range pairs(locations) {
			distX, distY := pair[0].distanceTo(pair[1])

			location := pair[0].subtract(distX, distY)
			if a.isInBounds(location) {
				result = append(result, location)
			}

			location = pair[1].subtract(-distX, -distY)
			if a.isInBounds(location) {
				result = append(result, location)
			}
		}
	}

	return result
}

// print Print the grid with the locations of the antennae
func (a *Grid) print() {
	str := a.getBaseGrid()

	for _, line := range str {
		fmt.Printf("%v\n", line)
	}
	fmt.Println()
}

// getBaseGrid Get the actual 2D map of the world
// from the spase list of antennae
func (a *Grid) getBaseGrid() [][]string {
	var str [][]string
	for i := 0; i<a.shape[0]; i++ {
		var row []string
		for j := 0; j<a.shape[1]; j++ {
			row = append(row, ".")
		}
		str = append(str, row)
	}

	for char, locs := range a.locations {
		for _, loc := range locs {
			str[loc.y][loc.x] = string(char)
		}
	}
	return str
}

// printWithAntinodes Print the grid with the locations of the antenne
// and of the antinodes
func (a *Grid) printWithAntinodes(nodes []Location) {
	str := a.getBaseGrid()

	for _, node := range nodes {
		str[node.y][node.x] = "#"
	}

	for _, line := range str {
		fmt.Printf("%v\n", line)
	}
	fmt.Println()
}


func readData() Grid {
	text, err := os.ReadFile("d8/input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(text), "\n")

	// Strip the trailing newline
	lines = lines[:len(lines)-1]
	result := make(map[rune][]Location)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			char := rune(line[j])

			if char != '.' {
				if _, exists := result[char]; !exists {
					result[char] = make([]Location, 0)
				}
				result[char] = append(result[char], Location{j, i})
			}
		}
	}
	return Grid{result, [2]int{len(lines), len(lines[0])}}
}

func Run() {
	grid := readData()

	nodes := grid.findAllNearestAntinodes()
	grid.printWithAntinodes(nodes)

	fmt.Println("[d8.1] number of nearest antinodes: ", len(grid.findAllNearestAntinodes()))
	fmt.Println("[d8.2] number of antinodes: ", len(grid.findAllAntinodes()))
}
