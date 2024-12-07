package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	var m runtime.MemStats

	// Read file
	start := time.Now()
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	runtime.ReadMemStats(&m)
	fmt.Println("Read file in", elapsed)
	fmt.Printf("Heap Alloc: %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Total Alloc: %v MB\n", m.TotalAlloc/1024/1024)

	// Split into sections
	start = time.Now()
	sections := strings.SplitN(string(data), "\n\n", 2)
	elapsed = time.Since(start)
	runtime.ReadMemStats(&m)
	fmt.Println("Split into sections in", elapsed)
	fmt.Printf("Heap Alloc: %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Total Alloc: %v MB\n", m.TotalAlloc/1024/1024)

	// Parse points
	start = time.Now()
	points := make(map[string]struct{})
	for _, line := range strings.SplitN(strings.TrimSpace(sections[0]), "\n", -1) {
		points[line] = struct{}{}
	}
	elapsed = time.Since(start)
	runtime.ReadMemStats(&m)
	fmt.Println("Parsed points in", elapsed)
	fmt.Printf("Heap Alloc: %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Total Alloc: %v MB\n", m.TotalAlloc/1024/1024)

	// Count valid lines
	start = time.Now()
	count := 0
LINE_LOOP:
	for _, line := range strings.SplitN(strings.TrimSpace(sections[1]), "\n", -1) {
		parts := strings.Split(line, ",")
		for i := 0; i < len(parts)-1; i++ {
			for j := i + 1; j < len(parts); j++ {
				if _, ok := points[parts[j]+"|"+parts[i]]; ok {
					slices.SortFunc(parts, func(a, b string) int {
						if _, ok := points[a+"|"+b]; ok {
							return -1
						}
						return 1
					})
					m, _ := strconv.Atoi(parts[len(parts)/2])
					count += m
					continue LINE_LOOP
				}
			}
		}
	}
	elapsed = time.Since(start)
	runtime.ReadMemStats(&m)
	fmt.Println("Counted valid lines in", elapsed)
	fmt.Printf("Heap Alloc: %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Total Alloc: %v MB\n", m.TotalAlloc/1024/1024)

	println(count)
}
