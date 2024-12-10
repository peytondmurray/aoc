package d8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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

func pairs(locs []Location) [][2]Location {
	var results [][2]Location
	for i := 0; i < len(locs); i++ {
		for j := i+1; j < len(locs); j++ {
			results = append(results, [2]Location{locs[i], locs[j]})
		}
	}

	return results
}

func (a *Grid) isInBounds(loc Location) bool {
	return loc.x >= 0 && loc.x < a.shape[1] && loc.y >= 0 && loc.y < a.shape[0]
}

func (a *Grid) findAllAntinodes() []Location {
	result := map[Location]struct{}{}
	for char, _ := range a.locations {

	}

}

func (a *Grid) findAntinodes(chars ...rune) []Location {
	result := map[Location]struct{}{}

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

func (a *Grid) print() {
	str := a.getBaseGrid()

	for _, line := range str {
		fmt.Printf("%v\n", line)
	}
	fmt.Println()
}

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
	// text, err := os.ReadFile("d8/input2")
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

			if _, exists := result[char]; !exists {
				result[char] = make([]Location, 0)
			}
			result[char] = append(result[char], Location{j, i})
		}
	}
	return Grid{result, [2]int{len(lines), len(lines[0])}}
}

func Run() {
	grid := readData()
	grid.print()

	nodes := grid.findAntinodes('a')
	grid.printWithAntinodes(nodes)

	// fmt.Println("[d7.1] : ", )
	// fmt.Println("[d7.2] : ", nLoops)
}
