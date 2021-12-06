package main

func main() {
  // day1part1()
  // day1part2()
	// day2part1()
  // day2part2()
  day3part1()
  // day3part2()
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
