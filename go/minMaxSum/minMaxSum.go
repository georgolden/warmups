package minMaxSum

func MinMaxSum(arr []int32) (uint32, uint32) {
	length := len(arr)
	first := uint32(arr[0])
	var minSum, maxSum, min, max uint32
	min = first
	max = first
	for i := 1; i < length; i++ {
		el := uint32(arr[i])
		if el < min {
			maxSum += min
			min = el
			minSum += el
		} else if el > max {
			minSum += max
			max = el
			maxSum += el
		} else {
			minSum += el
			maxSum += el
		}
	}
	return minSum, maxSum
}
