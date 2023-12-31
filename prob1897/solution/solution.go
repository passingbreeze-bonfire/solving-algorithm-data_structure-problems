package solution

func MakeEqual(words []string) bool {
	freqLen := len(words) // 반복되는 글자의 수는 words의 길이와 같다.
	answer := make(map[string]int)

	for _, word := range words {
		for _, char := range word {
			answer[string(char)]++
		}
	}

	for _, count := range answer {
		if count % freqLen != 0 {
			return false
		}
	}
	return true
}
