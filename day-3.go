package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "fmt"
)

func readDay3Data(fileName string) ([]int64, func() error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File not loaded")
	}

	var binaryNumbers []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		binaryString := scanner.Text()
		if err != nil {
			log.Fatal("Failed to convert string depth to int depth")
		}

    binary, err := strconv.ParseInt(binaryString, 2, 64)
    if err != nil {
        fmt.Println(err)
        log.Fatal("Failed to convert binary number")
    }


		binaryNumbers = append(binaryNumbers, binary)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return binaryNumbers, file.Close
}

type ComparisonStruct struct {
  oneBit int
  zeroBit int
}

func day3part1() {
	binaryList, closeFunction := readData("./day-3-input.txt")
	defer closeFunction()

  var lineBits []ComparisonStruct

  firstLine := strconv.Itoa(binaryList[0])
  for i := 0; i < len(firstLine); i++ {
    var bit ComparisonStruct
    bit.zeroBit = 0
    bit.oneBit = 0
    lineBits = append(lineBits, bit)
  }

  for idx, binary := range binaryList {
    s := strconv.Itoa(binary)
    fmt.Println(s)

    for _, numberRune := range s {
      number := int(numberRune)

      if number == 0 {
        lineBits[idx].zeroBit = lineBits[idx].zeroBit + 1
      } else if number == 1 {
        lineBits[idx].oneBit = lineBits[idx].oneBit + 1
      }
    }
  }

  fmt.Println(lineBits)

  commonBitString := ""
  for _, bit := range lineBits {
    isZero := bit.zeroBit > bit.oneBit

    if isZero {
      commonBitString = commonBitString + "0"
    } else {
      commonBitString = commonBitString + "1"
    }
  }

  binary, err := strconv.ParseInt(commonBitString, 2, 64)
  if err != nil {
    log.Fatal("Error converting to binary number")
  }

  fmt.Println("Most common binaries:", binary)

}

func day3part2() {}
