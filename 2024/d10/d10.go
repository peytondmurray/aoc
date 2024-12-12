package d10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readData() [][]int {
	// file, err := os.Open("d10/input2")
	file, err := os.Open("d10/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		row := make([]int, len(text))
		for i, char := range text {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal("Problem converting this value to integer: ", string(char))
			}
			row[i] = value
		}
		result = append(result, row)
	}
	return result
}

type Node struct {
	children []Node
	location *Location
	value int
}

type Location struct {
	x int
	y int
}

// findZeros Find the locations of every distinct zero on the trailmap.
func findZeros(trailmap [][]int) []Location {
	var zeros []Location
	for i, row := range trailmap {
		for j, value := range row {
			if value == 0 {
				zeros = append(zeros, Location{j, i})
			}
		}
	}
	return zeros
}

// neighbors Get a mapping between neighbors to the top, bottom, left, and right
// of another location (if they exist) and their topographic values
func neighbors(trailmap [][]int, location *Location) map[Location]int {
	sizey := len(trailmap)
	sizex := len(trailmap[0])
	result := make(map[Location]int)

	right := Location{location.x + 1, location.y}
	left := Location{location.x - 1, location.y}
	up := Location{location.x, location.y - 1}
	down := Location{location.x, location.y + 1}

	if right.x < sizex {
		result[right] = trailmap[right.y][right.x]
	}
	if left.x >= 0 {
		result[left] = trailmap[left.y][left.x]
	}
	if up.y >= 0 {
		result[up] = trailmap[up.y][up.x]
	}
	if down.y < sizey {
		result[down] = trailmap[down.y][down.x]
	}
	return result
}

// newNodeTree Generate a new node tree from the given location.
func newNodeTree(trailmap [][]int, location *Location) Node {
	value := trailmap[location.y][location.x]

	// Find all neighbors who have a height that increases by 1 from
	// the current location
	var children []Node
	for neighbor, neighborValue := range neighbors(trailmap, location) {
		if neighborValue == value + 1 {
			children = append(children, newNodeTree(trailmap, &neighbor))
		}
	}

	return Node{
		children, location, trailmap[location.y][location.x],
	}
}

// newTree Generate a new (root) tree from the trailmap.
func newTree(trailmap [][]int) Node {
	zeros := findZeros(trailmap)

	var roots []Node
	for _, zero := range zeros {
		roots = append(roots, newNodeTree(trailmap, &zero))
	}

	// Generate a dummy root node containing all the other root nodes
	return Node{roots, nil, -1}
}

// getTrailheadScore Return the number of nine-height positions from each trailhead (value == 0).
func getTrailheadScore(node *Node, distinctTrails map[Location]struct{}) int {
	loc := node.location

	// If this is the dummy root node, don't check for its presence in distinct trails,
	// just add up the score of its children
	if loc != nil {
		if _, exists := distinctTrails[*loc]; exists {
			return 0
		}

		if node.value == 9 {
			distinctTrails[*loc] = struct{}{}
			return 1
		}
	}


	values := 0
	for _, child := range node.children {
		values += getTrailheadScore(&child, distinctTrails)
	}
	return values
}

// getRating Get the number of distinct trails accessible from each zero-valued location.
func getRating(node *Node) int {
	loc := node.location

	// If this is the dummy root node, don't check for its presence in distinct trails,
	// just add up the score of its children
	if loc != nil {
		if node.value == 9 {
			return 1
		}
	}

	values := 0
	for _, child := range node.children {
		values += getRating(&child)
	}
	return values

}

// getRootTrailheadScore Get the trailhead score from the root node.
func getRootTrailheadScore(root *Node) int {
	score := 0
	for _, child := range root.children {
		score += getTrailheadScore(&child, map[Location]struct{}{})
	}
	return score
}

// printNode Print some information about a node.
func (a *Node) printNode() {
	if a.location != nil {
		fmt.Println("Location: ", a.location.x, a.location.y, "Value :", a.value)
	}

	for _, child := range a.children {
		fmt.Println(child.location.x, child.location.y)
	}
}

// printTree Print a tree. Child node stats are prefixed by `indent`.
func (a *Node) printTree(indent string) {
	if a.location == nil {
		fmt.Println("---Tree---")
		for _, child := range a.children {
			child.printTree("")
		}
	} else {
		fmt.Printf("%s(%d, %d): %d\n", indent, a.location.x, a.location.y, a.value)
		for _, child := range a.children {
			child.printTree(fmt.Sprintf("%s%s", indent, "  "))
		}
	}
}

func Run() {
	trailmap := readData()

	tree := newTree(trailmap)

	fmt.Println("[d10.1] trailhead score: ", getRootTrailheadScore(&tree))
	fmt.Println("[d10.2] trailhead rating: ", getRating(&tree))
}
