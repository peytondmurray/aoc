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
	// text, err := os.ReadFile("d14/input")
	text, err := os.ReadFile("d14/input2")
	if err != nil {
		log.Fatal(err)
	}

	var robots []Robot
	for _, line := range strings.Split(string(text), "\n") {
		if len(line) > 0 {
			robots = append(robots, newRobot(line))
		}
	}
	return Grid{robots, 11, 7}
	// return Grid{robots, 101, 103}
}

type Robot struct {
	x int
	y int
	vx int
	vy int
}

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

func (a *Grid) evolve(time int) {
	for i := range len(a.robots) {
		a.robots[i].x = (a.robots[i].x + time*a.robots[i].vx) % a.xshape
		a.robots[i].y = (a.robots[i].y + time*a.robots[i].vy) % a.yshape
	}
}

func (a *Grid) calculateSafetyFactor() int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, robot := range a.robots {
		if (robot.x < a.xshape/2) {
			if (robot.y < a.yshape/2) {
				q1++
			} else if (robot.y > a.yshape/2 + 1) {
				q3++
			}
		} else if (robot.x > a.xshape/2 + 1) {
			if (robot.y < a.yshape/2) {
				q2++
			} else if (robot.y > a.yshape/2 + 1) {
				q4++
			}
		}
	}
	return q1*q2*q3*q4
}

func (a *Grid) print(quads bool) {
	grid := make([][]int, a.yshape)
	for i := range a.yshape {
		grid[i] = make([]int, a.xshape)
	}

	for _, robot := range a.robots {
		grid[robot.y][robot.x] += 1
	}

	for i := range a.yshape {
		if quads && i == a.yshape/2 {
			fmt.Print("\n")
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


func Run() {
	grid := readData()
	grid.print(true)

	grid.evolve(100)

	grid.print(true)
	fmt.Println("[d13.1] safety score:", grid.calculateSafetyFactor())
	// fmt.Println("[d13.2] tokens:", solve(data, 10000000000000))
}
