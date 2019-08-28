package main

func Sum(arr [5]int) int {
	var sum int
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}
	return sum
}
