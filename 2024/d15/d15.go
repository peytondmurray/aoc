package d15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid struct {
	arr [][]string
}

func readGrid(scanner *bufio.Scanner) [][]string {
	var arr [][]string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			return arr
		}

		arr = append(arr, strings.Split(text, ""))
	}
	return arr
}

func readSteps(scanner *bufio.Scanner) []string {
	var steps []string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			return steps
		}
		steps = append(steps, strings.Split(text, "")...)
	}
	return steps
}

func readData() (Grid, []string) {
	file, err := os.Open("d15/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return Grid{readGrid(scanner)}, readSteps(scanner)
}
