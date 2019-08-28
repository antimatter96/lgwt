package main

func Sum(arr [5]int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
