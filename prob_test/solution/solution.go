package solution

import "sort"

func FindContentChildren(g []int, s []int) int { // 첫번째 통과한 정답
	answer := 0
	sort.Ints(g)
	sort.Ints(s)
	for _, size := range s {
		for i, greed := range g {
			if size >= greed {
				answer++
				g = append(g[:i], g[i+1:]...)
				break
			}
		}
	}
	return answer
}

func FindContentChildrenTwoPointer(g []int, s []int) int { // Two Pointers
	answer := 0
	cookieSizeIndex := 0
	sort.Ints(g)
	sort.Ints(s)
	for cookieSizeIndex < len(s) && answer < len(g) {
		if s[cookieSizeIndex] >= g[answer] {
			answer++
		}
		cookieSizeIndex++
	}
	return answer
}
