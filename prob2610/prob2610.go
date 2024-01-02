package main

import (
	"fmt"
	"solving-algorithm-data_structure-problems/prob2610/solution"
)

func main() {
	fmt.Println("Here is problem 2610 solution.")
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(solution.FindMatrix(nums1))

	nums2 := []int{1, 3, 4, 1, 2, 3, 1}
	fmt.Println(solution.FindMatrix(nums2))
}
