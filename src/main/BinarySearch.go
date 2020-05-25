package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 10, 20, 47, 59, 63, 75, 88, 99, 107, 120, 133, 155, 162, 176, 188, 199, 200, 210, 222}
	fmt.Println(binarySearch(arr, 47))
}

func binarySearch(arr []int, num int) int {
	windowStart := 0
	windowEnd := len(arr) - 1
	for {
		mid := int(math.Ceil(float64((windowEnd + windowStart) / 2)))
		if arr[mid] == num {
			return mid
		} else if num < arr[mid] {
			windowEnd = mid
		} else {
			windowStart = mid
		}
	}

	return -1
}
