package d14

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readData() Grid {
	text, err := os.ReadFile("d14/input")
	// text, err := os.ReadFile("d14/input2")
	if err != nil {
		log.Fatal(err)
	}

	var robots []Robot
	for _, line := range strings.Split(string(text), "\n") {
		if len(line) > 0 {
			robots = append(robots, newRobot(line))
		}
	}
	return Grid{robots, 101, 103}
	// return Grid{robots, 11, 7} // If using "d14/input2" above
}

type Robot struct {
	x int
	y int
	vx int
	vy int
}

// newRobot Generate a new robot from the input text.
func newRobot(text string) Robot {
	re := regexp.MustCompile(`p=(?P<x>\d+),(?P<y>\d+) v=(?P<vx>-?\d+),(?P<vy>-?\d+)`)
	match := re.FindStringSubmatch(text)

	x, err1 := strconv.Atoi(match[1])
	y, err2 := strconv.Atoi(match[2])
	vx, err3 := strconv.Atoi(match[3])
	vy, err4 := strconv.Atoi(match[4])

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Fatal("Error parsing robots.")
	}
	return Robot{x, y, vx, vy}
}

type Grid struct {
	robots []Robot
	xshape int
	yshape int
}

// evolve Evolve the grid forward in time.
func (a *Grid) evolve(time int) {
	for i, robot := range a.robots {
		newX := (robot.x + time*robot.vx) % a.xshape
		newY := (robot.y + time*robot.vy) % a.yshape

		if newX < 0 {
			newX += a.xshape
		}
		if newY < 0 {
			newY += a.yshape
		}

		a.robots[i].x = newX
		a.robots[i].y = newY
	}
}

// calculateSafetyFactor Calculate the safety factor for the grid.
func (a *Grid) calculateSafetyFactor() int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, robot := range a.robots {
		if (robot.x < a.xshape/2) {
			if (robot.y < a.yshape/2) {
				q1++
			} else if (robot.y > a.yshape/2) {
				q3++
			}
		} else if (robot.x > a.xshape/2) {
			if (robot.y < a.yshape/2) {
				q2++
			} else if (robot.y > a.yshape/2) {
				q4++
			}
		}
	}
	return q1*q2*q3*q4
}

// print Print the grid; if quads == true, the middle vertical
// and horizontal rows are omitted.
func (a *Grid) print(quads bool) {
	grid := toIntGrid(a)

	for i := range a.yshape {
		if quads && i == a.yshape/2 {
			fmt.Print("\n")
			continue
		}

		for j := range a.xshape {
			if quads && j == a.xshape/2 {
				fmt.Print(" ")
			}
			if grid[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(grid[i][j])
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// toIntGrid Render the sparse grid as a dense grid of integers.
func toIntGrid(a *Grid) [][]int {
	grid := make([][]int, a.yshape)
	for i := range a.yshape {
		grid[i] = make([]int, a.xshape)
	}

	for _, robot := range a.robots {
		grid[robot.y][robot.x] += 1
	}
	return grid
}

type Location struct {
	x int
	y int
}

type Region struct {
	locations map[Location]struct{}
}

func (a *Region) area() int {
	return len(a.locations)
}

// addToSet Add all items to set
func addToSet[T comparable](set map[T]struct{}, items ...T) {
	for _, item := range items {
		set[item] = struct{}{}
	}
}

// joinSets Add all items in the other set to set
func joinSets[T comparable](set map[T]struct{}, other map[T]struct{}) {
	for item := range other {
		set[item] = struct{}{}
	}
}

// getConnected Greedily search for locations connected to the given location.
func getConnected(loc Location, robots map[Location]struct{}, visited map[Location]struct{}) map[Location]struct{} {
	// If we've already visited the current location, just return
	if _, exists := visited[loc]; exists {
		return make(map[Location]struct{})
	}

	visited[loc] = struct{}{}

	// If there's no robot at the current location, just return
	if _, exists := robots[loc]; !exists {
		return make(map[Location]struct{})
	}

	// If there's a robot, check all the neighboring spaces
	connected := make(map[Location]struct{})
	addToSet(connected, loc)

	// Add on the robots on all sides of the current location
	joinSets(connected, getConnected(Location{loc.x+1, loc.y}, robots, visited))
	joinSets(connected, getConnected(Location{loc.x-1, loc.y}, robots, visited))
	joinSets(connected, getConnected(Location{loc.x, loc.y+1}, robots, visited))
	joinSets(connected, getConnected(Location{loc.x, loc.y-1}, robots, visited))
	return connected
}

// findRegions Find all the different regions in the current grid.
func findRegions(grid *Grid) []Region {
	var regions []Region

	// Get a map of all robots
	robots := make(map[Location]struct{})
	for _, robot := range grid.robots {
		robots[Location{robot.x, robot.y}] = struct{}{}
	}

	// Break into regions
	visited := make(map[Location]struct{})
	for robot := range robots {
		if _, exists := visited[robot]; !exists {
			regions = append(regions, Region{getConnected(robot, robots, visited)})
		}
	}
	return regions
}

// maybeChristmas Check if the grid has a grouping of 10 or more robots.
// Could they be in the shape of a christmas tree?
func maybeChristmas(grid *Grid) bool {
	regions := findRegions(grid)
	for _, region := range regions {
		if region.area() > 10 {
			return true
		}
	}
	return false
}

// evolveTilChristmas Evolve the grid until a grouping of robots appears;
// then print the grid and wait for the user to look to see if it's a
// christmas tree. Press return to continue evolving.
func (a *Grid) evolveTilChristmas() int {
	for i := 1;; i++ {
		a.evolve(1)
		if maybeChristmas(a) {
			fmt.Println("i = ", i)
			a.print(false)

			// Wait for the user to press return before continuing
			var input string
			fmt.Scanln(&input)
			if input != "\n" {
				continue
			}
		}
	}
}

func Run() {
	grid := readData()
	grid.evolve(100)
	fmt.Println("[d14.1] safety factor:", grid.calculateSafetyFactor())
	fmt.Println("[d14.2] evolving to christmas:", grid.evolveTilChristmas())
}
