package d11

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// readData Read in the data, return an array of stones.
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

// intPow Compute the integer power, val^exp.
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

// countDigits Count the number of digits in a value. Significantly faster
// than int(math.Log10(val)) + 1.
func countDigits(val int) int {
	digits := 1
	for (val / 10 != 0) {
		val /= 10
		digits += 1
	}
	return digits
}

// stepSimple First attempt at a pure iterative approach to computing the stones.
// Doesn't work well past 45 blinks.
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

// stepPreallocate Compute the stones after a single blink. Doesn't really
// work for anything much larger than 45 blinks. Preallocates space for the
// stones by assuming that every stone will split into two. Significantly
// faster (~50% of the run time) of `stepSimple`, but still too slow.
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

// blinkRecursive Initial implementation of the recursive blink.
// Stores the stones in an array, but this is slow.
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

// Cached A key in the cache.
type Cached struct {
	stone int
	blinks int
}

// blinkRecursiveNoArrayCached Compute the number of stones after n
// blinks given the starting stone. No array is used to keep track of
// the current stones, and a cache is used to short-circuit computations that
// we've already done.
func blinkRecursiveNoArrayCached(stone int, n int, cache map[Cached]int) int {
	if n == 0 {
		return 1
	}

	result, exists := cache[Cached{stone, n}]
	if exists {
		return result
	}

	if stone == 0 {
		value := blinkRecursiveNoArrayCached(1, n-1, cache)
		cache[Cached{1, n-1}] = value
		result += value
	} else {
		nDigits := countDigits(stone)
		if nDigits % 2 == 0 {
			factor := intPow(10, nDigits/2)

			// Find the first new stone value by moving the decimal over
			// by half the number of digits; int division drops the right
			// half of the digits
			left := stone / factor
			value := blinkRecursiveNoArrayCached(left, n-1, cache)
			cache[Cached{left, n-1}] = value
			result += value

			// Second new stone value is just the original value of the stone
			// minus the first half of the digits
			right := stone - (left*factor)
			value = blinkRecursiveNoArrayCached(right, n-1, cache)
			cache[Cached{right, n-1}] = value
			result += value
		} else {
			newStone := stone*2024
			value := blinkRecursiveNoArrayCached(newStone, n-1, cache)
			cache[Cached{newStone, n-1}] = value
			result += value
		}
	}
	cache[Cached{stone, n}] = result
	return result
}

// blinkRecursiveNoArray Compute the number of stones after n blinks
// given the starting stone. No array needed to keep track of the current
// set of stones; it's all stored in the stack.
//
// Note this will not converge quick enough for part 2 - you _need_ caching
// to make it fast enough.
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

// lenAfterNBlinks Compute the number of stones after a given number of blinks.
// Prints a progress bar as it goes, but caching makes this so fast it's unnecessary.
//
// For 75 blinks, the cache length is 134k entries, but the number of stones is
// 220357186726677. It saves a huge amount of work, and the computation runs almost
// instantly.
func lenAfterNBlinks(stones []int, blinks int) int {
	cache := make(map[Cached]int)
	result := 0
	for i, stone := range stones {
		result += blinkRecursiveNoArrayCached(stone, blinks, cache)
		fmt.Print("\033[G\033[K")
		fmt.Printf("stone %2d/%2d\n", i, len(stones)-1)
		fmt.Print("\033[A")
	}
	return result
}

func Run() {
	stones := readData()

	fmt.Println("[d11.1] Number of stones after 25 blinks:", lenAfterNBlinks(stones, 25))
	fmt.Println("[d11.2] Number of stones after 75 blinks:", lenAfterNBlinks(stones, 75))
}
