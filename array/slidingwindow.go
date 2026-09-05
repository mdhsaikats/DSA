package array

import (
	"fmt"
)

func FixedWindow() {
	arr := []int{2, 5, 3, 7, 1}
	windowSize := 3

	windowSum := arr[0] + arr[1] + arr[2]
	fmt.Println("First window sum:", windowSum)

	for right := windowSize; right < len(arr); right++ {
		left := right - windowSize
		windowSum = windowSum + arr[right] - arr[left]
		fmt.Println("Next window sum:", windowSum)
	}
}

func VariableFixed() {
	arr := []int{2, 5, 3, 7, 1}
	target := 7

	shortestLength := 999999
	windowSum := 0
	left := 0

	for right := 0; right < len(arr); right++ {
		windowSum += arr[right]

		for windowSum >= target {
			currentLength := right - left + 1
			if currentLength < shortestLength {
				shortestLength = currentLength
			}
			windowSum -= arr[left]
			left++
		}
	}
	fmt.Println("The shortest window length is:", shortestLength)
}
