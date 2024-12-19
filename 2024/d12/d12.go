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

func parseRegion(data [][]string, location Location, visited map[Location]string) Region {
	char := data[location.y][location.x]
	locations := explore(data, char, location, visited)
	area, perimeter := analyze(locations)

	return Region{char, area, perimeter, locations}
}

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

func computePrice(regions []Region) int {
	result := 0
	for _, region := range regions {
		pt := region.area*region.perimeter
		result += pt

		fmt.Println(region.symbol, ":", "area", region.area, "perimeter", region.perimeter, "cost", pt)
	}
	return result
}

func Run() {
	data := readData()
	regions := parseRegions(data)

	fmt.Println("[d12.1] Fence pricing:", computePrice(regions))
	// fmt.Println("[d11.2] Number of stones after 75 blinks:", lenAfterNBlinks(stones, 75))
}
