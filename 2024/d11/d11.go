package d11

import (
	"fmt"
	"log"
	"os"
	// "runtime/pprof"
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
		// newStones = stepPreallocate(newStones)
		newStones = stepSimple(newStones)
		fmt.Print("\033[G\033[K")
		fmt.Printf("%2d/%2d\n", i, n)
		fmt.Print("\033[A")
	}
	return newStones
}

func blinkRecursive(stones []int, n int) int {
	if n == 0 {
		return len(stones)
	}

	result := 0
	for _, stone := range stones {
		if stone == 0 {
			result += blinkRecursive([]int{1}, n-1)
		} else {
			nDigits := countDigits(stone)
			if nDigits % 2 == 0 {
				factor := intPow(10, nDigits/2)

				// Find the first new stone value by moving the decimal over
				// by half the number of digits; int division drops the right
				// half of the digits
				left := stone / factor
				result += blinkRecursive([]int{left}, n-1)

				// Second new stone value is just the original value of the stone
				// minus the first half of the digits
				result += blinkRecursive([]int{stone - (left*factor)}, n-1)
			} else {
				result += blinkRecursive([]int{stone*2024}, n-1)
			}
		}
	}
	return result
}

func blinkRecursiveNoArray(stone int, n int) int {
	if n == 0 {
		return 1
	}

	result := 0
	if stone == 0 {
		result += blinkRecursiveNoArray(1, n-1)
	} else {
		nDigits := countDigits(stone)
		if nDigits % 2 == 0 {
			factor := intPow(10, nDigits/2)

			// Find the first new stone value by moving the decimal over
			// by half the number of digits; int division drops the right
			// half of the digits
			left := stone / factor
			result += blinkRecursiveNoArray(left, n-1)

			// Second new stone value is just the original value of the stone
			// minus the first half of the digits
			result += blinkRecursiveNoArray(stone - (left*factor), n-1)
		} else {
			result += blinkRecursiveNoArray(stone*2024, n-1)
		}
	}
	return result
}

func Run() {
	stones := readData()

	// f, err := os.Create("cpu_step_recursive_no_array.prof")
	// if err != nil {
	// 	log.Fatal("Couldn't create CPU profile: ", err)
	// }
	// defer f.Close()
	//
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	log.Fatal("Couldn't start CPU profile: ", err)
	// }
	// defer pprof.StopCPUProfile()

	// stones = blink(stones, 25)
	// fmt.Println("[d11.1] Number of stones after 25 blinks: ", len(stones))
	// stones = blink(stones, 47)


	// result := 0
	// for _, stone := range stones {
	// 	result += len(blink([]int{stone}, 40))
	// }

	// fmt.Println("[d11.2] Number of stones after 75 blinks: ", result)
	// fmt.Println("[d11.2] Number of stones after 75 blinks: ", len(stones))

	// fmt.Println("[d11.2] Number of stones after 6 blinks: ", blinkRecursive(stones, 6))
	// fmt.Println("[d11.2] Number of stones after 25 blinks: ", blinkRecursive(stones, 25))

	result := 0
	for i, stone := range stones {
		result += blinkRecursiveNoArray(stone, 47)
		fmt.Print("\033[G\033[K")
		fmt.Printf("stone %2d/%2d\n", i, len(stones))
		fmt.Print("\033[A")
	}
	fmt.Println("[d11.2] Number of stones after 40 blinks: ", result)
	// fmt.Println("[d11.2] Number of stones after 25 blinks: ", blinkRecursive(stones, 25))
}
