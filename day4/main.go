package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	grid, err := parseInput(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	accessible := false
	numAccessible := 0
	numRemoved := 0
	removedRoll := true
	for removedRoll {
		removedRoll = false
		numAccessible = 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] {
					accessible = isAccessible(grid, i, j)
					if accessible {
						fmt.Print("Y ")
						grid[i][j] = false
						numAccessible++
					} else {
						fmt.Print("N ")
					}
				} else {
					fmt.Print("N ")
				}
			}
			fmt.Printf("\n")
		}
		if numAccessible > 0 {
			removedRoll = true
		}
		numRemoved += numAccessible
		fmt.Printf("\n")
		fmt.Printf("Accessible: %d\n", numAccessible)
	}
	fmt.Printf("Removed: %d\n", numRemoved)
}

func parseInput(filename string) ([][]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]bool
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, char := range line {
			row[i] = char == '@'
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func isAccessible(grid [][]bool, x int, y int) bool {
	rolls := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if hasRoll(grid, i, j) {
				rolls++
			}
			if rolls >= 4 {
				return false
			}
		}
	}
	return true
}

func hasRoll(grid [][]bool, x int, y int) bool {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}
	return grid[x][y]
}
