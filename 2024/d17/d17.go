package d17

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	A int
	B int
	C int
}

func (a *Machine) combo(n int) int {
	if n >= 0 && n <= 3 {
		return n
	} else if n == 4 {
		return a.A
	} else if n == 5 {
		return a.B
	} else if n == 6 {
		return a.C
	} else if n == 7 {
		log.Fatal("Invalid literal operand")
	}
	return -1
}

func (a *Machine) exec(prog Program) {
	for i := 0; i < len(prog); {
		op := prog[i]
		n := prog[i+1]

		if op == 0 {
			a.A /= intPow(2, a.combo(n))
		} else if op == 1 {
			a.B ^= n
		} else if op == 2 {
			a.B = a.combo(n) % 8
		} else if op == 3 {
			if a.A != 0 {
				i = n
				continue
			}
		} else if op == 4 {
			a.B ^= a.C
		} else if op == 5 {
			fmt.Print(a.combo(n), ",")
		} else if op == 6 {
			a.B = a.A / intPow(2, a.combo(n))
		} else if op == 7 {
			a.C = a.A / intPow(2, a.combo(n))
		}
		i += 2
	}
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

type Program []int

func readData() (Machine, Program) {
	file, err := os.Open("d17/input2")
	// file, err := os.Open("d17/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a int
	var b int
	var c int
	var program Program

	reA := regexp.MustCompile(`Register A: (?P<value>\d+)`)
	reB := regexp.MustCompile(`Register B: (?P<value>\d+)`)
	reC := regexp.MustCompile(`Register C: (?P<value>\d+)`)
	reProgram := regexp.MustCompile(`(?P<program>(\d+,\d+(,?))+)`)
	for scanner.Scan() {
		text := scanner.Text()
		matchA := reA.FindStringSubmatch(text)
		if matchA != nil {
			val, err := strconv.Atoi(matchA[1])
			if err != nil {
				log.Fatal("Error parsing register A")
			}
			a = val
		}


		matchB := reB.FindStringSubmatch(text)
		if matchB != nil {
			val, err := strconv.Atoi(matchB[1])
			if err != nil {
				log.Fatal("Error parsing register B")
			}
			b = val
		}

		matchC := reC.FindStringSubmatch(text)
		if matchC != nil {
			val, err := strconv.Atoi(matchC[1])
			if err != nil {
				log.Fatal("Error parsing register C")
			}
			c = val
		}

		matchProgram := reProgram.FindStringSubmatch(text)
		if matchProgram != nil {
			for _, val := range strings.Split(matchProgram[1], ",") {
				val, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal("Error parsing integer in program")
				}
				program = append(program, val)
			}
		}
	}
	return Machine{a, b, c}, program
}

func Run() {
	machine, program := readData()

	fmt.Println("[d17.1] output:")
	machine.exec(program)
	// fmt.Println("[d17.2] optimal seats:", len(paths))
}
