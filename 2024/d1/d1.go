package d1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readData() ([]int, []int) {
	file, err := os.Open("d1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Can't use a Reader because Reader.Comma is of type rune
	// Need to read line by line...?
	scanner := bufio.NewScanner(file)
	var c1 []int
	var c2 []int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		val1, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}

		val2, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		c1 = append(c1, val1)
		c2 = append(c2, val2)
	}
	return c1, c2
}

func difference(c1 []int, c2 []int) int {
	var result int
	for i := range(len(c1)) {
		diff := c1[i] - c2[i]

		if diff < 0 {
			diff = -diff
		}
		result += diff
	}

	return result
}

func similarity(c1 []int, c2 []int) int {
	var result int

	// Get the occurrences of each value in c2
	c2Counts := make(map[int]int)
	for _, value := range(c2) {
		c2Counts[value] += 1
	}

	// Similarity is (value)*(occurrences of value in c2)
	// for each value in c1
	for _, value := range(c1) {
		result += c2Counts[value] * value
	}

	return result
}

func Run() {
	c1, c2 := readData()

	sort.Ints(c1)
	sort.Ints(c2)

	diff := difference(c1, c2)
	similarity := similarity(c1, c2)

	fmt.Println("[d1.1] Difference: ", diff)
	fmt.Println("[d1.2] Similarity: ", similarity)
}
