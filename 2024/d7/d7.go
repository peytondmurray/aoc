package d7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func power(a int, b int) int {
	if b == 0 {
		return 1
	}

	if b == 1 {
		return a
	}

	result := a
	for i:=1; i<b; i++ {
		result *= a
	}
	return result
}
// fmt.Printf("Output: %20d Calculations: %5d len(inputs): %3d expected calculations: %5d\n", output, len(calculations), len(inputs), power(2, len(inputs) - 1))

func readData() ([]int, [][]int) {
	// file, err := os.Open("d7/input2")
	file, err := os.Open("d7/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineRe := regexp.MustCompile(`(?P<output>\d+): (?P<inputs>.*)`)
	inputsRe := regexp.MustCompile(`(\d+)`)

	var outputs []int
	var inputs [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := lineRe.FindStringSubmatch(line)

		if match[1] != "" && match[2] != "" {
			output, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal("Problem converting this set of numbers to integers: ", match[2])
			}

			var newInputs []int
			for _, intMatch := range inputsRe.FindAllStringSubmatch(match[2], -1) {
				value, err := strconv.Atoi(intMatch[1])
				if err != nil {
					log.Fatal("Problem converting this set of numbers to integers: ", match[2])
				}
				newInputs = append(newInputs, value)
			}
			outputs = append(outputs, output)
			inputs = append(inputs, newInputs)
		} else {
			log.Fatal("Match for this line is invalid: ", line)
		}
	}
	return outputs, inputs
}

func getCalibration(outputs []int, inputs [][]int, allowConcat bool) int {
	result := 0
	for i, output := range outputs {
		if isValid(output, inputs[i], allowConcat) {
			result += output
		}
	}
	return result
}

func isValid(output int, inputs []int, allowConcat bool) bool {
	calculations := calculate(inputs, allowConcat)


	for _, result := range calculations {
		if result == output {
			return true
		}
	}
	return false
}

func concat(a int, b int) int {
	digits := int(math.Log10(float64(b))) + 1
	return (a*power(10, digits)) + b
}

func calculate(inputs []int, allowConcat bool) []int {
	if len(inputs) == 0 {
		log.Fatal("Empty input found.")
	}
	if len(inputs) == 1 {
		return inputs
	}

	var results []int
	last := len(inputs) - 1

	// The innermost calls to calculate are the ones
	// that operate on the beginning of the slice,
	// meaning that the operators associate left to right.
	for _, value := range calculate(inputs[:last], allowConcat) {
		results = append(results, inputs[last]*value)
		results = append(results, inputs[last]+value)

		if allowConcat {
			results = append(results, concat(value, inputs[last]))
		}
	}
	return results
}

func Run() {
	outputs, inputs := readData()
	fmt.Println("[d7.1] calibration result: ", getCalibration(outputs, inputs, false))
	fmt.Println("[d7.2] calibration result with concatenation operator: ", getCalibration(outputs, inputs, true))
}
