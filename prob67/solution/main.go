package solution

func AddBinary(a string, b string) string {
	result := ""
	carry := 0
	sum := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		sum = carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = string(sum%2+'0') + result
		carry = sum / 2
	}
	if carry != 0 {
		result = "1" + result
	}
	return result
}
