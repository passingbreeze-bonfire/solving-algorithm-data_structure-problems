package solution

func MakeEqual(words []string) bool {
    answer := make(map[string]int)
	for _, word := range words {
		for _, char := range word {
			answer[string(char)]++
		}
	}
	var prev int = 0
	for _, count := range answer {
		if prev != 0 && prev != count {
			return false
		}
		prev = count
	}
    if len(answer) > 0 && prev == 1 {
        return false
    }
	return true
}
