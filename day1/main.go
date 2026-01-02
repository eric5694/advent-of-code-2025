package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	Direction rune
	Clicks    int
}

const (
	MAX = 99
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input-file>")
		os.Exit(1)
	}

	instructions, err := readInput(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	curr := 50
	zeroes := 0
	newZeroes := 0
	fmt.Printf("Position: %d\n", curr)
	for _, inst := range instructions {
		curr, newZeroes = turn(inst, curr)
		if curr == 0 {
			zeroes++
		}
		zeroes += newZeroes
		fmt.Printf("Direction: %c, Clicks: %d, Position: %d, NewZeroes: %d, Zeroes: %d\n", inst.Direction, inst.Clicks, curr, newZeroes, zeroes)
	}
	//fmt.Printf("Zeroes: %d\n", zeroes)
}

func readInput(filename string) ([]Instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		direction := rune(line[0])
		if direction != 'L' && direction != 'R' {
			return nil, fmt.Errorf("invalid direction: %c", direction)
		}

		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", line[1:])
		}

		instructions = append(instructions, Instruction{
			Direction: direction,
			Clicks:    clicks,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}

func turn(instr Instruction, start int) (value int, zeroes int) {
	value = start
	zeroes = 0
	clicks := instr.Clicks

	for clicks > MAX {
		clicks -= MAX + 1
		zeroes++
	}

	if instr.Direction == 'L' {
		value -= clicks
	} else {
		value += clicks
	}

	for value < 0 {
		if start != 0 {
			zeroes++
		}
		value += MAX + 1
	}

	for value > MAX {
		value -= MAX + 1
		if value != 0 {
			zeroes++
		}
	}

	return
}
