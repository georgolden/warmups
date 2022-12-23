package plusMinus

/*
Given an array of ints, positive, negative or 0
Return ratios of postive/length, negative/length, 0-value/length
Input: [10, 30, 0, -23, -45]
Return: 0.400000, 0.400000, 0.200000
*/
func PlusMinus(arr []int) []float64 {
	length := float64(len(arr))
	var positives, negatives, zeros float64 = 0, 0, 0
	for _, v := range arr {
		if v > 0 {
			positives++
		} else if v < 0 {
			negatives++
		} else {
			zeros++
		}
	}
	return []float64{positives / length, negatives / length, zeros / length}
}
