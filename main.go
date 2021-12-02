package main

func main() {
	day1part2()
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
