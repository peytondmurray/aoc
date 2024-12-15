package d11

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type Stone struct {
	value int
	left *Stone
	right *Stone
}

// llStep Iterate over the stones to the right of the current stone,
// mutating them
func (a *Stone) llStep() {
	// Keep track of the next element in the linked list
	// that needs to be stepped, before we potentially modify the current
	// element
	stone := a
	for stone != nil {
		if stone.value == 0 {
			stone.value = 1
			stone = stone.right
		} else {
			nDigits := countDigits(stone.value)
			if nDigits % 2 == 0 {
				factor := intPow(10, nDigits/2)

				// This stone's value is found by moving the decimal over
				// by half the number of digits; int division drops the right
				// half of the digits
				newValue := stone.value / factor

				// New stone value is just the original value of the stone
				// minus the first half of the digits
				newStone := &Stone{stone.value - (newValue*factor), stone, stone.right}

				stone.value = newValue
				stone.right = newStone
				stone = newStone.right
			} else {
				stone.value *= 2024
				stone = stone.right
			}
		}
	}
}

// len Count all stones to the right of a given stone
func (a *Stone) len() int {
	result := 1
	for stone := a; stone.right != nil; stone = stone.right {
		result++
	}
	return result
}

func (a *Stone) print() {
	stone := a
	var values []int
	for stone != nil {
		values = append(values, stone.value)
		stone = stone.right
	}
	fmt.Println(values)
}

func makeNewStones(values []int) Stone {
	leftStone := &Stone{values[0], nil, nil}
	first := leftStone
	for i := 1; i<len(values); i++ {
		stone := Stone{values[i], leftStone, nil}
		leftStone.right = &stone
		leftStone = &stone
	}
	return *first
}

func blinkStones(stones *Stone, n int) {
	for i := range n {
		stones.llStep()
		fmt.Print("\033[G\033[K")
		fmt.Printf("%2d/%2d\n", i, n)
		fmt.Print("\033[A")
	}
}

func RunLL() {
	stones := makeNewStones(readData())

	f, err := os.Create("cpu_linked_list.prof")
	if err != nil {
		log.Fatal("Couldn't create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Couldn't start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	blinkStones(&stones, 25)
	fmt.Println("[d11.1] Number of stones after 25 blinks: ", stones.len())
	// blinkStones(&stones, 50)
	// fmt.Println("[d11.2] Number of stones after 75 blinks: ", stones.len())
}
