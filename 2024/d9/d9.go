package d9

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func readData() Disk {
	// text, err := os.ReadFile("d9/input2")
	text, err := os.ReadFile("d9/input")
	if err != nil {
		log.Fatal(err)
	}

	return newDisk(text)
}

func newDisk(input []byte) Disk {
	var ids []int
	var sizes []int
	var freeSpaces []int

	id := 0
	for i, value := range input {
		if value == '\n' {
			break
		}

		if i % 2 == 0 {
			ids = append(ids, id)
			size, err := strconv.Atoi(string(value))
			if err != nil {
				log.Fatal("Problem converting this value to integer: ", value)
			}
			sizes = append(sizes, size)

			id += 1
		} else {
			freeSpace, err := strconv.Atoi(string(value))
			if err != nil {
				log.Fatal("Problem converting this value to integer: ", value)
			}
			freeSpaces = append(freeSpaces, freeSpace)
		}
	}

	// There's no free space after the last element
	freeSpaces = append(freeSpaces, 0)

	return Disk{input, ids, sizes, freeSpaces}
}

type Disk struct {
	diskMap []byte
	id []int
	size []int
	freeSpace []int
}

// Convert the disk to its string representation
func (a *Disk) render() []string {
	var result []string

	for i, id := range a.id {
		for range a.size[i] {
			result = append(result, strconv.Itoa(id))
		}

		for range a.freeSpace[i] {
			result = append(result, ".")
		}
	}
	return result
}

// Do a simple defrag, moving pieces of files as far left as possible
func (a *Disk) defrag() []string {
	rendered := a.render()

	i := 0
	j := len(rendered) - 1

	for i != j {
		if rendered[i] != "." {
			i++
			continue
		}

		if rendered[j] == "." {
			j--
			continue
		}

		rendered[i], rendered[j] = rendered[j], rendered[i]
		i++
	}

	return rendered
}

// Index into the file IDs, and get the corresponding index in the rendered disk
func (a *Disk) indexOfRendered(idx int) int {
	start := 0
	for i := 0; i<idx; i++ {
		start += a.size[i] + a.freeSpace[i]
	}
	return start
}

func (a *Disk) defragWholeFiles() []string {
	rendered := a.render()

	// Iterate backwards through the ids
	for idx := len(a.id) - 1; idx >= 0; idx-- {
		start := a.indexOfRendered(idx)
		size := a.size[idx]

		// Find the next free space, starting from the beginning of the render
		for i := 0; i<start; i++ {
			if rendered[i] != "." {
				continue
			}

			// Once a free space has been found, find the end of it
			j := i
			for j < start && rendered[j] == "." {
				j++
			}
			space := j - i

			// If the space required by the file at index `start` and of length `size`
			// fits, swap it and move on to the next file
			if space >= size {
				swap(rendered[i:i+size], rendered[start:start+size])
				break
			}

			// If not, continue searching for another open space. We can skip ahead
			// by the size of the free space here.
			i += space
		}
	}
	return rendered
}

// swap Swap the contents of two slices
func swap[T comparable](vals1 []T, vals2 []T) {
	for i := range vals1 {
		vals1[i], vals2[i] = vals2[i], vals1[i]
	}
}

// Compute the checksum, which is the sum of the index*value
// of each value in values
func checksum(values []string) int {
	result := 0
	for i, value := range values {
		if value != "." {
			intval, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("Problem converting this value to integer: ", value)
			}
			result += i*intval
		}
	}
	return result
}

func Run() {
	disk := readData()

	fmt.Println("[d9.1] checksum: ", checksum(disk.defrag()))
	fmt.Println("[d9.2] checksum: ", checksum(disk.defragWholeFiles()))
}

func main() {
	Run()
}
