package solution

func FindMatrix(nums []int) [][]int {
	freq := make(map[int]int)  // nums 배열에 들어있는 숫자들의 반복 횟수를 메모한다.
	answer := make([][]int, 0) // 정답이 저장될 2차원 배열

	for _, num := range nums { // 배열 순회
		if freq[num] >= len(answer) { // 숫자 반복 횟수가 answer의 길이보다 크거나 같으면
			answer = append(answer, []int{}) // 빈 배열을 하나 추가해준다.
		}
		answer[freq[num]] = append(answer[freq[num]], num) // 정답 배열 속 1차원 배열에 숫자를 추가한다.
		freq[num]++                                        // 반복 횟수 추가
	}

	return answer
}
