package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

type SubmarineInstruction struct {
  Type string
  Value int
}

func readDay2Data(fileName string) ([]SubmarineInstruction, func() error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File not loaded")
	}

	var instructions []SubmarineInstruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructionString := scanner.Text()
		if err != nil {
			log.Fatal("Failted to convert string depth to int depth")
		}

    split := strings.Split(instructionString, " ")
    var ins SubmarineInstruction
    insValue, err := strconv.Atoi(split[1]) 
    if err != nil {
      log.Fatal("Failted to convert instruction value to int")
    }

    ins.Type = split[0]
    ins.Value = insValue
		instructions = append(instructions, ins)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions, file.Close
}

func day2part1() {
  instructions, closeFile := readDay2Data("./day-2-input.txt")
  defer closeFile()

  depth := 0
  horizontalPosition := 0

  for _, instruction := range instructions {
    switch instruction.Type {
    case "forward":
      horizontalPosition = horizontalPosition + instruction.Value
    case "down":
      depth = depth + instruction.Value
    case "up":
      depth = depth - instruction.Value
    }
  }

  fmt.Println("Depth is:", depth)
  fmt.Println("horizontalPosition is:", horizontalPosition)
  fmt.Println("Result is:", depth*horizontalPosition)
}

func day2part2() {
  instructions, closeFile := readDay2Data("./day-2-input.txt")
  defer closeFile()

  aim := 0
  depth := 0
  horizontalPosition := 0

  for _, instruction := range instructions {
    switch instruction.Type {
    case "forward":
      horizontalPosition = horizontalPosition + instruction.Value
      depth = depth + (aim * instruction.Value)
    case "down":
      aim = aim + instruction.Value
    case "up":
      aim = aim - instruction.Value
    }
  }

  fmt.Println("Depth is:", depth)
  fmt.Println("horizontalPosition is:", horizontalPosition)
  fmt.Println("Result is:", depth*horizontalPosition)

}
