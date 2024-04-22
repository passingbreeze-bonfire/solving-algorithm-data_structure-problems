package solution

func PlusOne(digits []int) []int {
	digitLength := len(digits)
	for i := digitLength - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
