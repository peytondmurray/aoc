package d11

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

func readData() []int {
	text, err := os.ReadFile("d11/input2")
	// text, err := os.ReadFile("d11/input")
	if err != nil {
		log.Fatal(err)
	}

	fields := strings.Fields(string(text))
	values := make([]int, len(fields))
	for i, field := range fields {
		value, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(err)
		}
		values[i] = value
	}

	return values
}

func intPow(val int, exp int) int {
	if exp == 0 {
		return 1
	}
	result := val
	for range exp - 1 {
		result *= val
	}
	return result
}

func countDigits(val int) int {
	digits := 1
	for (val / 10 != 0) {
		val /= 10
		digits += 1
	}
	return digits
}

func stepSimple(stones []int) []int {
	// 1. if the stone of value 0, it is replaced by a stone with value 1
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else {
			nDigits := countDigits(stone)
			if nDigits % 2 == 0 {
				factor := intPow(10, nDigits/2)

				// Find the first new stone value by moving the decimal over
				// by half the number of digits; int division drops the right
				// half of the digits
				left := stone / factor

				// Second new stone value is just the original value of the stone
				// minus the first half of the digits
				right := stone - (left*factor)

				newStones = append(newStones, left)
				newStones = append(newStones, right)
			} else {
				newStones = append(newStones, 2024*stone)
			}
		}
	}
	return newStones
}

func stepPreallocate(stones []int) []int {
	// 1. if the stone of value 0, it is replaced by a stone with value 1
	// var newStones []int

	nStones := 0
	newStones := make([]int, 2*len(stones))
	for _, stone := range stones {
		if stone == 0 {
			newStones[nStones] = 1
			nStones++
		} else {
			nDigits := countDigits(stone)
			if nDigits % 2 == 0 {
				factor := intPow(10, nDigits/2)

				// Find the first new stone value by moving the decimal over
				// by half the number of digits; int division drops the right
				// half of the digits
				left := stone / factor

				// Second new stone value is just the original value of the stone
				// minus the first half of the digits
				right := stone - (left*factor)

				newStones[nStones] = left
				nStones++
				newStones[nStones] = right
				nStones++
			} else {
				newStones[nStones] = stone*2024
				nStones++
			}
		}
	}
	return newStones[:nStones]
}

func blink(stones []int, n int) []int {
	newStones := stones
	for i := range n {
		newStones = stepPreallocate(newStones)
		// newStones = stepSimple(newStones)
		fmt.Print("\033[G\033[K")
		fmt.Printf("%2d/%2d\n", i, n)
		fmt.Print("\033[A")
	}
	return newStones
}


func Run() {
	stones := readData()

	f, err := os.Create("cpu_step_preallocate.prof")
	if err != nil {
		log.Fatal("Couldn't create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Couldn't start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// stones = blink(stones, 25)
	// fmt.Println("[d11.1] Number of stones after 25 blinks: ", len(stones))
	// stones = blink(stones, 47)


	result := 0
	for _, stone := range stones {
		result += len(blink([]int{stone}, 47))
	}

	fmt.Println("[d11.2] Number of stones after 75 blinks: ", result)
	// fmt.Println("[d11.2] Number of stones after 75 blinks: ", len(stones))
}
