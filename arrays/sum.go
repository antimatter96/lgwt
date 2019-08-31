package main

func Sum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}

func SumAll(arr ...[]int) (sums []int) {
	sums = make([]int, len(arr))
	for i, val := range arr {
		sums[i] += Sum(val)
	}
	return sums
}

func SumAll2(arr ...[]int) []int {
	var sums []int
	for _, val := range arr {
		sums = append(sums, Sum(val))
	}
	return sums
}
