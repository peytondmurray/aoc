package d5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// readRules Read the rules from the input. The first line is assumed
// to be the first rule.
func readRules(scanner *bufio.Scanner) (map[int][]int, error) {
	result := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "|")
		if len(split) != 2 {
			return result, nil
		}

		v1, err1 := strconv.Atoi(split[0])
		v2, err2 := strconv.Atoi(split[1])

		if err1 != nil || err2 != nil {
			return make(map[int][]int), fmt.Errorf(
				"Problem converting this line to integers: %v",
				line,
			)
		}
		_, ok := result[v1]
		if ok {
			result[v1] = append(result[v1], v2)
		} else {
			result[v1] = []int{v2}
		}
	}
	return make(map[int][]int), fmt.Errorf("Line not found.")
}

// readUpdates Read the updates section of the input using the
// given scanner; assume that the next line to be scanned is the first
// update.
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
					"Problem converting this line to integers: %v",
					line,
				)
			}
			row = append(row, value)
		}
		result = append(result, row)
	}
	return result, nil
}

// readData Read the raw data and parse it into a list of rules and updates.
//
// The rules are a mapping from a value (page number) to a slice of page
// numbers that should come after it.
//
// The updates are a slice of slices containing page numbers.
func readData() (map[int][]int, [][]int) {
	file, err := os.Open("d5/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules, err := readRules(scanner)
	if err != nil {
		log.Fatal("Problem reading rules.")
	}

	updates, err := readUpdates(scanner)
	if err != nil {
		log.Fatal("Problem reading updates.", err)
	}
	return rules, updates
}

// contains Return true if the value is found in the items.
func contains(items []int, val int) bool {
	for _, item := range items {
		if item == val {
			return true
		}
	}
	return false
}

// Check if the update is ordered.
//
// An update is ordered if any value to the left of any other value
// doesn't have a corresponding entry in the rules map.
func isOrdered(rules map[int][]int, update []int) bool {
	for i := range len(update) - 1 {
		leftVal := update[i]
		for j := i + 1; j < len(update) - 1; j++ {
			rightVals, ok := rules[update[j]]
			if ok && contains(rightVals, leftVal) {
				return false
			}
		}
	}
	return true
}


// Return the indices of the ordered updates
func findOrdered(rules map[int][]int, updates [][]int) []int {
	var result []int
	for i, update := range updates {
		if isOrdered(rules, update) {
			result = append(result, i)
		}
	}
	return result
}

// sumAllMiddles Sum the middle values of every element in a slice
func sumAllMiddles(updates [][]int) int {
	result := 0
	for _, update := range updates {
		result += update[len(update)/2]
	}
	return result
}


// sumMiddles Sum the middle values of every specified index
func sumMiddles(updates [][]int, indicesToSum []int) int {
	result := 0
	for _, i := range indicesToSum {
		update := updates[i]
		result += update[len(update)/2]
	}
	return result
}

// Object which pairs a value (page number) and the values (page numbers) that
// should come after it
type ValueRule struct {
	Value int
	ToRight []int
}

// Define an interface for go's sort.Sort so that we don't have to roll our own
// sorting function
// https://pkg.go.dev/sort#Interface
type ByRule []ValueRule
func (a ByRule) Len() int {
	return len(a)
}
func (a ByRule) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByRule) Less(i int, j int) bool {
	return contains(a[i].ToRight, a[j].Value)
}

// fixDisorderedUpdate Fix the ordering of a disordered update slice
func fixDisorderedUpdate(rules map[int][]int, update []int) []int {
	var result []int

	// Pair up each value with the values that are greater than it
	var valueRules []ValueRule
	for _, val := range update {
		valueRules = append(valueRules, ValueRule{val, rules[val]})
	}

	// Use go's sort.Sort to reorder valueRules before retrieving the actual values
	sort.Sort(ByRule(valueRules))
	for _, vr := range valueRules {
		result = append(result, vr.Value)
	}

	return result
}

// fixDisordered Fixes and returns all disordered updates
func fixDisordered(rules map[int][]int, updates [][]int, ordered []int) [][]int {
	var result [][]int

	for i, update := range updates {
		if !contains(ordered, i) {
			result = append(result, fixDisorderedUpdate(rules, update))
		}
	}
	return result
}

func Run() {
	rules, updates := readData()

	ordered := findOrdered(rules, updates)
	fixed := fixDisordered(rules, updates, ordered)

	fmt.Println("[d5.1] sum of middle values of ordered updates: ", sumMiddles(updates, ordered))
	fmt.Println("[d5.2] sum of middle values of fixed updates: ", sumAllMiddles(fixed))
}
