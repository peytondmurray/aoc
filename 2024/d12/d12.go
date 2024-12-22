package d12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readData() [][]string {
	file, err := os.Open("d12/input")
	// file, err := os.Open("d12/input2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		row := make([]string, len(text))
		for i, char := range text {
			row[i] = string(char)
		}
		result = append(result, row)
	}

	return result
}

type Region struct {
	symbol string
	area int
	perimeter int
	locations map[Location]struct{}
}

type Location struct {
	x int
	y int
}

// explore Generate a map of locations containing the same letter as and connected to `location`.
func explore(data [][]string, char string, location Location, visited map[Location]string) map[Location]struct{} {
	if location.x < 0 || location.x >= len(data[0]) || location.y < 0 || location.y >= len(data) {
		return make(map[Location]struct{})
	}

	if data[location.y][location.x] != char {
		return make(map[Location]struct{})
	}
	visited[location] = char

	up := Location{location.x, location.y-1}
	right := Location{location.x+1, location.y}
	down := Location{location.x, location.y+1}
	left := Location{location.x-1, location.y}

	locations := make(map[Location]struct{})
	locations[location] = struct{}{}
	if up.y >= 0 {
		if _, exists := visited[up]; !exists {
			for loc := range explore(data, char, up, visited) {
				locations[loc] = struct{}{}
			}
		}
	}

	if right.x < len(data[0]) {
		if _, exists := visited[right]; !exists {
			for loc := range explore(data, char, right, visited) {
				locations[loc] = struct{}{}
			}
		}
	}

	if down.y < len(data) {
		if _, exists := visited[down]; !exists {
			for loc := range explore(data, char, down, visited) {
				locations[loc] = struct{}{}
			}
		}
	}
	if left.x >= 0 {
		if _, exists := visited[left]; !exists {
			for loc := range explore(data, char, left, visited) {
				locations[loc] = struct{}{}
			}
		}
	}
	return locations
}

// analyze Cound the number of fences for a set of locations;
// used in d12p1.
func analyze(locations map[Location]struct{}) (int, int) {
	perimeter := 0
	for loc := range locations {
		up := Location{loc.x, loc.y-1}
		right := Location{loc.x+1, loc.y}
		down := Location{loc.x, loc.y+1}
		left := Location{loc.x-1, loc.y}

		perimeter += 4
		if _, exists := locations[up]; exists {
			perimeter--
		}
		if _, exists := locations[down]; exists {
			perimeter--
		}
		if _, exists := locations[left]; exists {
			perimeter--
		}
		if _, exists := locations[right]; exists {
			perimeter--
		}
	}

	return len(locations), perimeter
}

// printRegion Print a region, including highlighted letters.
func printRegion(data [][]string, highlight []Location) {

	locs := make(map[Location]struct{})
	for _, loc := range highlight {
		locs[loc] = struct{}{}
	}

	reset := "\033[0m"
	red := "\033[31m"

	fmt.Print("\n")
	for i, row := range data {
		for j, char := range row {
			if _, exists := locs[Location{j, i}]; exists {
				fmt.Print(red, char, reset)
			} else {
				fmt.Print(char)
			}
		}
		fmt.Print("\n")
	}

	fmt.Println()
}


// parseRegion Fully explore all connected sites of a region on the grid.
func parseRegion(data [][]string, location Location, visited map[Location]string) Region {
	char := data[location.y][location.x]
	locations := explore(data, char, location, visited)
	area, perimeter := analyze(locations)

	return Region{char, area, perimeter, locations}
}

// parseRegions Parse the data into regions.
func parseRegions(data [][]string) []Region {
	visited := make(map[Location]string)
	var regions []Region

	for i, row := range data {
		for j := range row {
			loc := Location{j, i}
			if _, exists := visited[loc]; !exists {
				regions = append(regions, parseRegion(data, loc, visited))
			}
		}
	}

	return regions
}

// computePrice Compute the price for d12p1.
func computePrice(regions []Region) int {
	result := 0
	for _, region := range regions {
		pt := region.area*region.perimeter
		result += pt
	}
	return result
}

// raster "Raster" across the dataset to find the edges of a region.
// Use these edges to find the number of sides of each region, which is
// the perimeter in d12p2.
func raster(data [][]string, region Region) int {

	// Find the edges of the dataset
	lEdges := make([][]int, len(data))
	rEdges := make([][]int, len(data))
	uEdges := make([][]int, len(data))
	dEdges := make([][]int, len(data))
	for i := range data {
		lEdges[i] = make([]int, len(data[0]))
		rEdges[i] = make([]int, len(data[0]))
		uEdges[i] = make([]int, len(data[0]))
		dEdges[i] = make([]int, len(data[0]))
	}

	for loc := range region.locations {
		left := Location{loc.x-1, loc.y}
		right := Location{loc.x+1, loc.y}
		up := Location{loc.x, loc.y-1}
		down := Location{loc.x, loc.y+1}

		if _, exists := region.locations[left]; !exists {
			lEdges[loc.y][loc.x] = 1
		}

		if _, exists := region.locations[right]; !exists {
			rEdges[loc.y][loc.x] = 1
		}

		if _, exists := region.locations[up]; !exists {
			uEdges[loc.y][loc.x] = 1
		}

		if _, exists := region.locations[down]; !exists {
			dEdges[loc.y][loc.x] = 1
		}
	}

	// Looking at the left and right edges, these from top
	// to bottom; consecutive flips don't count.
	sum := 0
	for j := range data[0] {
		lastL := 0
		lastR := 0
		for i := range data {
			if lastL == 0 && lEdges[i][j] == 1 {
				sum++
			}

			if lastR == 0 && rEdges[i][j] == 1 {
				sum++
			}

			lastL = lEdges[i][j]
			lastR = rEdges[i][j]
		}
	}

	// Looking at the top and bottom edges, sum left to
	// right; consecutive flips don't count.
	for i := range data {
		lastU := 0
		lastD := 0
		for j := range data[0] {
			if lastU == 0 && uEdges[i][j] == 1 {
				sum++
			}

			if lastD == 0 && dEdges[i][j] == 1 {
				sum++
			}
			lastU = uEdges[i][j]
			lastD = dEdges[i][j]
		}
	}

	return sum
}

// printData Print the dataset.
func printData[T comparable](data [][]T) {
	for _, row := range data {
		fmt.Println(row)
	}
	fmt.Println()
}

// computeAdjustedPrice Compute the price for d12p2.
func computeAdjustedPrice(regions []Region, data [][]string) int {
	result := 0
	for _, region := range regions {
		result += region.area*raster(data, region)
	}
	return result
}

func Run() {
	data := readData()
	regions := parseRegions(data)

	fmt.Println("[d12.1] Fence pricing:", computePrice(regions))
	fmt.Println("[d12.2] Fence pricing with joined edges:", computeAdjustedPrice(regions, data))
}
