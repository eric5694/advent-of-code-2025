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

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var digits [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineDigits []int
		for _, char := range line {
			digit := int(char - '0')
			lineDigits = append(lineDigits, digit)
		}
		digits = append(digits, lineDigits)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Read %d lines\n", len(digits))

	totalJoltage := 0
	for _, line := range digits {
		joltage := getLargestJoltage2(line)
		fmt.Printf("Line: %v, Joltage: %d\n", line, joltage)
		totalJoltage += joltage
	}
	fmt.Printf("Total Joltage: %d\n", totalJoltage)
}

func getLargestJoltage(digits []int) (joltage int) {
	largestLeft := digits[0]
	largestRight := digits[1]
	for i := 2; i < len(digits); i++ {
		d := digits[i]
		if largestLeft < largestRight {
			largestLeft = largestRight
			largestRight = d
			continue
		}
		if largestRight < d {
			largestRight = d
			continue
		}
	}

	joltage = largestLeft*10 + largestRight
	return
}

func getLargestJoltage2(digits []int) (joltage int) {
	joltageArray := make([]int, 12)
	index := 0
	for i := 0; i < len(joltageArray); i++ {
		for j := index; j <= len(digits)-12+i; j++ {
			if digits[j] > joltageArray[i] {
				joltageArray[i] = digits[j]
				index = j + 1
			}
		}
	}

	return digitsToInt(joltageArray)
}

func digitsToInt(digits []int) int {
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	return result
}
