package d11

import (
	"fmt"
	"log"
	// "math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

func readData() []int {
	// text, err := os.ReadFile("d11/input2")
	text, err := os.ReadFile("d11/input")
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

func step(stones []int) []int {
	// 1. if the stone of value 0, it is replaced by a stone with value 1
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else {

			str := strconv.Itoa(stone)
			if len(str) % 2 == 0 {

				left, err := strconv.Atoi(str[:len(str)/2])
				if err != nil {
					log.Fatal("Invalid stone")
				}

				right, err := strconv.Atoi(str[len(str)/2:])
				if err != nil {
					log.Fatal("Invalid stone")
				}
				newStones = append(newStones, left)
				newStones = append(newStones, right)

			} else {
				newStones = append(newStones, stone*2024)
			}

			// nDigits := (int(math.Log10(float64(stone))) + 1)
			// if nDigits % 2 == 0 {
			// 	factor := intPow(10, nDigits/2)
			//
			// 	// Find the first new stone value by moving the decimal over
			// 	// by half the number of digits; int division drops the right
			// 	// half of the digits
			// 	left := stone / factor
			//
			// 	// Second new stone value is just the original value of the stone
			// 	// minus the first half of the digits
			// 	right := stone - (left*factor)
			//
			// 	newStones = append(newStones, left)
			// 	newStones = append(newStones, right)
			// } else {
			// 	newStones = append(newStones, stone*2024)
			// }
		}
	}
	return newStones
}

func blink(stones []int, n int) []int {
	newStones := stones
	for i := range n {
		newStones = step(newStones)
		fmt.Print("\033[G\033[K")
		fmt.Printf("%2d/%2d\n", i, n)
		fmt.Print("\033[A")
	}
	return newStones
}

func Run() {
	stones := readData()

	f, err := os.Create("cpu_str.prof")
	if err != nil {
		log.Fatal("Couldn't create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Couldn't start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// fmt.Println("[d11.1] Number of stones after 25 blinks: ", len(blink(stones, 25)))
	fmt.Println("[d11.2] Number of stones after 25 blinks: ", len(blink(stones, 40)))
	// fmt.Println("[d11.2] Number of stones after 25 blinks: ", len(blink(stones, 75)))
}
