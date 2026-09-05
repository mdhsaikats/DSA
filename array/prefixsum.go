package array

func PrefixSum() ([5]int, int) {
	arr := [5]int{2, 4, 1, 5, 3}
	presum := [5]int{}

	presum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		presum[i] = presum[i-1] + arr[i]
	}
	sum := presum[3] - presum[0]

	return presum, sum
}
