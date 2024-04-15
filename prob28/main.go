package main

import "fmt"

func main() {
	// 1 <= haystack.length, needle.length <= 10^4
	// haystack and needle consist of only lowercase English characters.

	// ex1
	haystack := "sadbutsad"
	needle := "sad"
	if strStr(haystack, needle) == 0 {
		fmt.Println("ex1 is ok")
	}

	// ex2
	haystack = "leetcode"
	needle = "leeto"
	if strStr(haystack, needle) == -1 {
		fmt.Println("ex2 is ok")
	}

	// ex3
	haystack = "a"
	needle = "a"
	if strStr(haystack, needle) == 0 {
		fmt.Println("ex3 is ok")
	}

	// ex4
	haystack = "abc"
	needle = "c"
	if strStr(haystack, needle) == 2 {
		fmt.Println("ex4 is ok")
	}
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
