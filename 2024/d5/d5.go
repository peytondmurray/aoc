package d5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readRules(scanner *bufio.Scanner) (map[int]int, error) {
	var result map[int]int
	for scanner.Scan() {
		line := scanner.Text()

		// Separator between rules and updates
		if line == "\n" {
			return result, nil
		}

		split := strings.Split(line, "|")
		v1, err1 := strconv.Atoi(split[0])
		v2, err2 := strconv.Atoi(split[1])

		if err1 != nil || err2 != nil {
			return make(map[int]int), fmt.Errorf(
				"Problem converting this line to integers: ",
				line,
			)
		}
		result[v1] = v2
	}
	return make(map[int]int), fmt.Errorf("Line not found.")
}

func readUpdates(scanner *bufio.Scanner) ([][]int, error) {
	var result [][]int
	for scanner.Scan() {
		line := scanner.Text()

		// EOF
		if line == "\n" {
			continue
		}

		var row []int
		for _, text := range strings.Split(line, ",") {
			value, err := strconv.Atoi(text)
			if err != nil {
				return make([][]int, 0), fmt.Errorf(
					"Problem converting this line to integers: ",
					line,
				)
			}
			row = append(row, value)
		}
		result = append(result, row)
	}
	return result, nil
}

func readData() (map[int]int, [][]int) {
	file, err := os.Open("d5/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules, err1 := readRules(scanner)
	updates, err2 := readUpdates(scanner)

	if err1 != nil || err2 != nil {
		log.Fatal("Problem reading data.")
	}
	return rules, updates
}

func Run() {
	rules, updates := readData()

	fmt.Println("[d4.1] nXmas: ", findAllXmas(data))
	fmt.Println("[d4.2] nCrossed: ", findAllCrossed(data))
}
