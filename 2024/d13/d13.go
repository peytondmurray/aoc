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

func isIntegral(val float32) bool {
	return val == float32(int(val))
}

// solveMachine For a given machine, the problem can be written as
// a system of linear equations:
//
// [ Ax Bx ][ Na ]  = [ Px ]
// [ Ay By ][ Nb ]  = [ Py ]
//
// Where {Na, Nb} are the number of button presses of A and B,
// {Px, Py} are the prize location (x, y), and {Ax, Ay, Bx, By}
// are how much the A and B buttons move the claw on each press.
//
// As long as the movements of the claw are linearly independent,
// there's always a unique {Na, Nb} that can be found that can
// be made to get to {Px, Py}. The solutions are those for which
// there are _integer_ numbers of presses of A and B.
//
// To solve this, just invert the 2x2 matrix and multiply by
// {Px, Py}. Inversion will fail if the claw vectors are not
// linearly independent.
func solveMachine(machine ClawMachine, factor int) int {

	px := machine.prize.x + factor
	py := machine.prize.y + factor

	ax := machine.A.dx
	ay := machine.A.dy

	bx := machine.B.dx
	by := machine.B.dy

	na := int(float64(px*by - py*bx)/float64(ax*by - ay*bx))
	nb := int(float64(py*ax - px*ay)/float64(ax*by - ay*bx))

	// The truncation above only ever reduces the number of presses
	// of each button; we therefore only need to increase each button
	// press by 1 to see if it gets us the prize.
	if (na+1)*ax + nb*bx == px && (na+1)*ay + nb*by == py {
		return 3*(na+1) + nb
	}
	if na*ax + (nb+1)*bx == px && na*ay + (nb+1)*by == py {
		return 3*na + nb+1
	}
	if (na+1)*ax + (nb+1)*bx == px && (na+1)*ay + (nb+1)*by == py {
		return 3*(na+1) + nb+1
	}
	if na*ax + nb*bx == px && na*ay + nb*by == py {
		// If na and nb are already integral, no need for adjustment
		return 3*na + nb
	}

	// No solution exists
	return 0
}


func solve(machines []ClawMachine, factor int) int {
	result := 0
	for _, machine := range machines {
		result += solveMachine(machine, factor)
	}
	return result
}

func Run() {
	data := readData()
	fmt.Println("[d13.1] tokens:", solve(data, 0))
	fmt.Println("[d13.2] tokens:", solve(data, 10000000000000))
}
