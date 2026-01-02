package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	line := strings.TrimSpace(string(data))
	parts := strings.Split(line, ",")

	var ranges []Range
	for _, part := range parts {
		part = strings.TrimSpace(part)
		nums := strings.Split(part, "-")
		if len(nums) != 2 {
			fmt.Printf("Invalid format: %s\n", part)
			continue
		}

		start, err := strconv.Atoi(strings.TrimSpace(nums[0]))
		if err != nil {
			fmt.Printf("Invalid start number in %s: %v\n", part, err)
			continue
		}

		end, err := strconv.Atoi(strings.TrimSpace(nums[1]))
		if err != nil {
			fmt.Printf("Invalid end number in %s: %v\n", part, err)
			continue
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	sum := 0
	for _, r := range ranges {
		fmt.Printf("Start: %d, End: %d\n", r.Start, r.End)
		invalidIds := getInvalidIds2(r.Start, r.End)
		fmt.Printf("Invalid ids: %v\n", invalidIds)
		for _, id := range invalidIds {
			sum += id
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func getInvalidIds(start int, end int) (invalidIds []int) {
	for i := start; i <= end; i++ {
		numStr := strconv.Itoa(i)
		numLen := len(numStr)
		if numLen%2 == 1 {
			continue
		}

		halfNum := numStr[:numLen/2]

		if numStr == halfNum+halfNum {
			invalidIds = append(invalidIds, i)
		}
	}
	return
}

func getInvalidIds2(start int, end int) (invalidIds []int) {
	for i := start; i <= end; i++ {
		numStr := strconv.Itoa(i)
		numLen := len(numStr)
		for j := numLen / 2; j > 0; j-- {
			if numLen%j != 0 {
				continue
			}
			window := numStr[:j]
			testNum := strings.Repeat(window, numLen/j)
			if testNum == numStr && !slices.Contains(invalidIds, i) {
				invalidIds = append(invalidIds, i)
			}
		}
	}
	return
}
