package d3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readData() string {
	result, err := os.ReadFile("d3/input")
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

// enableDoDont toggles whether the `do()` and `don't()` statements matter
func extractValues(str string, enableDoDont bool) [][]int {
	var result [][]int

	enabled := true
	re := regexp.MustCompile(`(mul\((?P<v1>\d+),(?P<v2>\d+)\)|do\(\)|don't\(\))`)
	for _, match := range re.FindAllStringSubmatch(str, -1) {

		if match[1] == "don't()" {
			enabled = false
		} else if match[1] == "do()" {
			enabled = true
		} else {
			if !enableDoDont || enabled {
				v1, err1 := strconv.Atoi(match[2])
				v2, err2 := strconv.Atoi(match[3])

				if err1 != nil || err2 != nil {
					log.Fatal("Problem converting this match to integers: ", match)
				}

				result = append(result, []int{v1, v2})
			}
		}
	}
	return result
}

func sumPairs(pairs [][]int) int {
	result := 0
	for _, pair := range pairs {
		result += pair[0] * pair[1]
	}
	return result
}

func Run() {
	str := readData()

	fmt.Println("[d3.1] sum of multiplications: ", sumPairs(extractValues(str, false)))
	fmt.Println("[d3.2] sum of multiplications with do() and don't(): ", sumPairs(extractValues(str, true)))
}
