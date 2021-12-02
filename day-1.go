package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readData(fileName string) ([]int, func() error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File not loaded")
	}

	var depths []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failted to convert string depth to int depth")
		}
		depths = append(depths, depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return depths, file.Close
}

func day1part2() {
	depths, closeFunction := readData("./day-1-input.txt")
	defer closeFunction()

	var result = 0
	var previousSum = 0
	for idx := range depths {
		if idx > 1 {
			tempSum := sum(depths[idx-2 : idx+1])

			if previousSum != 0 && tempSum > previousSum {
				result = result + 1
			}
			previousSum = tempSum
		}
	}
	fmt.Println("Day 1 - Part 2, result:", result)
}

func day1part1() {
	depths, closeFunction := readData("./day-1-input.txt")
	defer closeFunction()

	var result = 0
	var previousDepth = 0
	for _, depth := range depths {
		if previousDepth != 0 && depth > previousDepth {
			result = result + 1
		}
		previousDepth = depth
	}
	fmt.Println("Day 1 - Part 1, result:", result)
}
