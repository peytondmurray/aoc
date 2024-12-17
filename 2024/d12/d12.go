package d12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readData() [][]string {
	file, err := os.Open("d12/input2")
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

// func parseRegions(data [][]string) map[string]*Region {
//
// 	regions := make(map[string]*Region)
//
// 	// Iterating across the array from left to right, top to bottom.
// 	// Only need to check left and up to adjust perimeter of each
// 	// region.
// 	for i, row := range data {
// 		for j, char := range row {
// 			if region, exists := regions[char]; exists {
// 				region.area++
//
// 				// Up, down, left, right of current location
// 				region.perimeter += 4
//
// 				// If there's a square to the left, remove one fence for the
// 				// current square, and another for the previous square too
// 				if _, locExists := region.locations[Location{j-1, i}]; locExists {
// 					region.perimeter -= 2
// 				}
//
// 				// Similar for the neighbor to the top
// 				if _, locExists := region.locations[Location{j, i-1}]; locExists {
// 					region.perimeter -= 2
// 				}
// 				region.locations[Location{j, i}] = struct{}{}
// 			} else {
// 				locs := make(map[Location]struct{})
// 				locs[Location{j, i}] = struct{}{}
// 				regions[char] = &Region{char, 1, 4, locs}
// 			}
// 		}
// 	}
// 	return regions
// }

func explore(data [][]string, location Location, visited map[Location]string) []Location {
	char := data[location.y][location.x]
	visited[location] = char

	up := Location{location.y - 1, location.x}
	right := Location{location.y, location.x + 1}
	down := Location{location.y+1, location.x}
	left := Location{location.y, location.x-1}

	var locations []Location
	if up.y >= 0 {
		if _, exists := visited[up]; !exists {
			for _, loc := range explore(data, up, visited) {
				locations = append(locations, loc)
			}
		}
	}

	if right.x < len(data[0]) {
		if _, exists := visited[right]; !exists {
			for _, loc := range explore(data, right, visited) {
				locations = append(locations, loc)
			}
		}
	}

	if down.y < len(data) {
		if _, exists := visited[down]; !exists {
			for _, loc := range explore(data, down, visited) {
				locations = append(locations, loc)
			}
		}
	}
	if left.x >= 0 {
		if _, exists := visited[left]; !exists {
			for _, loc := range explore(data, left, visited) {
				locations = append(locations, loc)
			}
		}
	}
	return locations
}

func parseRegions(data [][]string) []Region {
	visited := map[Location]string
	var regions []Region

	for i, row := range data {
		for j := range row {
			loc := Location{j, i}

			if _, exists := visited[loc]; !exists {
				regions = append(regions, explore(data, loc, visited))
			}
		}
	}

	return regions
}

func computePrice(regions map[string]*Region) int {
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
