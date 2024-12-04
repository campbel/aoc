package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid, err := load()
	if err != nil {
		panic(err)
	}

	safeCount := 0
	for _, row := range grid {
		if isSafeOrSubvariant(row) {
			safeCount++
		}
	}
	println(safeCount)
}

func load() ([][]int, error) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		grid[i] = make([]int, len(parts))
		for j, part := range parts {
			part = strings.TrimSpace(part)
			grid[i][j], err = strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
		}
	}
	return grid, nil
}

func isSafeOrSubvariant(row []int) bool {
	if isSafe(row) {
		fmt.Println("safe", row)
		return true
	}
	for i := range row {
		newRow := removeIndex(row, i)
		if isSafe(newRow) {
			fmt.Println("safe subvariant", row)
			return true
		}
	}
	fmt.Println("unsafe", row)
	return false
}

func removeIndex(row []int, i int) []int {
	newRow := make([]int, len(row))
	copy(newRow, row)
	return append(newRow[:i], newRow[i+1:]...)
}

func isSafe(row []int) bool {
	increasing := false
	decreasing := false
	for i := range row {
		if i >= len(row)-1 {
			continue
		}
		diff := row[i] - row[i+1]
		if diff < 0 {
			diff = -diff
			decreasing = true
		} else {
			increasing = true
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return (increasing || decreasing) && (increasing != decreasing)
}
