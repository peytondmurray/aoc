package d13

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	dx int
	dy int
}

type ClawMachine struct {
	A Button
	B Button
	prize Location
}

func (a *ClawMachine) findMinCost(
	target Location,
	factor int,
	nab Nab,
	cost int,
	cache map[Location]int,
) (int, Nab) {
	loc := nab.loc(a.A, a.B)

	if loc.x > factor*target.x || loc.y > factor*target.y {
		return 0, 0
	}
	return
}

type Location struct {
	x int
	y int
}

func newButton(text string) Button {
	re := regexp.MustCompile(`Button [AB]: X\+(?P<dx>\d+), Y\+(?P<dy>\d+)`)
	match := re.FindStringSubmatch(text)

	dx, err1 := strconv.Atoi(match[1])
	dy, err2 := strconv.Atoi(match[2])

	if err1 != nil || err2 != nil {
		log.Fatal("Error parsing button.")
	}
	return Button{dx, dy}
}

func parsePrize(text string) Location {
	re := regexp.MustCompile(`Prize: X=(?P<x>\d+), Y=(?P<y>\d+)`)
	match := re.FindStringSubmatch(text)

	x, err1 := strconv.Atoi(match[1])
	y, err2 := strconv.Atoi(match[2])

	if err1 != nil || err2 != nil {
		log.Fatal("Error parsing button.")
	}
	return Location{x, y}

}

func readData() []ClawMachine {
	text, err := os.ReadFile("d13/input")
	// text, err := os.ReadFile("d13/input2")
	if err != nil {
		log.Fatal(err)
	}

	var result []ClawMachine

	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(lines) - 1; i+=4 {
		result = append(
			result,
			ClawMachine{newButton(lines[i]), newButton(lines[i+1]), parsePrize(lines[i+2])},
		)
	}
	return result
}

func solve(machines []ClawMachine) int {
	result := 0
	for _, machine := range machines {
		result += solveMachine(machine)
	}
	return result
}


func solveMachine(machine ClawMachine) int {
	minCost := 0
	for i := range 100 {
		for j := range 100 {
			if (
				(i*machine.A.dx + j*machine.B.dx) == machine.prize.x &&
				(i*machine.A.dy + j*machine.B.dy) == machine.prize.y) {

				// Cost: A - 3 tokens, B - 1 token
				cost := i*3 + j
				if minCost == 0 || cost < minCost {
					minCost = cost
				}
			}
		}
	}
	return minCost
}

func solveMachineP2(machine ClawMachine, factor int) int {
	cost, nab := findMinCost(
		machine.prize,
		machine.A,
		machine.B,
		factor,
		Nab{0, 0},
		0,
		make(map[Location]int),
	)
	return cost
}

// func findMinCost(
// 	target Location,
// 	A Button,
// 	B Button,
// 	factor int,
// 	nab Nab,
// 	cost int,
// 	cache map[Location]int,
// ) (int, Nab) {
// 	loc := nab.loc(A, B)
// 	if loc.x > factor*target.x || loc.y > factor*target.y {
// 		return 0, Nab{}
// 	}
// 	if loc.x == factor*target.x && loc.y == factor*target.y {
// 		return cost, nab
// 	}
//
// 	// Minimum cost has already been cached
// 	if cachedCost, exists := cache[nab]; exists {
// 		return cost + cachedCost, nab
// 	}
//
// 	ca := findMinCost(target, A, B, factor, Location{loc.x+A.dx, loc.y+A.dy}, cost + 3, cache)
// 	cb := findMinCost(target, A, B, factor, Location{loc.x+B.dx, loc.y+B.dy}, cost + 1, cache)
//
// 	if ca > 0 {
// 		if cb > 0 {
// 			cost := min(ca, cb)
// 			cache[loc] = cost
// 			return cost
// 		}
// 		cache[loc] = ca
// 		return ca
// 	}
//
// 	if cb > 0 {
// 		cache[loc] = cb
// 		return cb
// 	}
// 	cache[loc] = 0
// 	return 0
// }

type Nab struct {
	na int
	nb int
}

func (a *Nab) cost() int {
	return a.na*3 + a.nb
}

func (a *Nab) loc(A Button, B Button) Location {
	return Location{a.na*A.dx + a.nb*B.dx, a.na*A.dy + a.nb*B.dy}
}

func solveP2(machines []ClawMachine) int {
	result := 0
	for _, machine := range machines {
		result += solveMachineP2(machine, 10000000000000)
	}
	return result
}

func Run() {

	data := readData()
	fmt.Println("[d13.1] tokens:", solve(data))
	fmt.Println("[d13.2] tokens:", solveP2(data))
}
