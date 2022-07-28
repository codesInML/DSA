package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

// time complexity: O(a + b)
// space complexity: O(a + b)
func buildString(str string) string {
	builtStringArray := []string{}

	for _, s := range strings.Split(str, "") {
		if s != "#" {
			builtStringArray = append(builtStringArray, s)
		} else {
			if len(builtStringArray) > 0 {
				builtStringArray = builtStringArray[:len(builtStringArray)-1]
			}
		}
	}
	return strings.Join(builtStringArray, "")
}

func stringsEqual(S, T string) bool {
	finalS, finalT := buildString(S), buildString(T)
	return finalS == finalT
}

func lengthOfLongestSubstring(s string) (longest int) {
	if len(s) <= 1 {
		return len(s)
	}

	// BRUTE FORCE SOLUTION (time complexity: O(N^2), space complexity: O(N))

	// for i := 0; i < len(s); i++ {
	// 	counter := 1
	// 	strTrack := map[string]int{string(s[i]): 1}
	// 	fmt.Println(strTrack)
	// 	for j := i + 1; j < len(s); j++ {
	// 		_, ok := strTrack[string(s[j])]
	// 		if !ok {
	// 			strTrack[string(s[j])] = 1
	// 			counter++
	// 			fmt.Println("counter", counter)
	// 			longest = int(math.Max(float64(counter), float64(longest)))
	// 		} else {
	// 			break
	// 		}
	// 		fmt.Println(strTrack)
	// 	}
	// }

	// Optimal solution with sliding window technique
	// (time complexity: O(N), space complexity: O(N))
	left := 0
	seenChars := map[string]int{}

	for right := 0; right < len(s); right++ {
		currentChar := s[right]
		prevSeenChar, ok := seenChars[string(currentChar)]

		if ok {
			if left <= prevSeenChar {
				left = prevSeenChar + 1
			}
		}
		seenChars[string(currentChar)] = right
		longest = int(math.Max(float64(right-left+1), float64(longest)))
	}
	return
}

func isValidPalindrome(s string) bool {
	s = strings.ToLower(regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(s, ""))
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlmostPalindrome(s string) bool {
	s = strings.ToLower(regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(s, ""))
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] != s[right] {
			return isValidPalindrome(s[:left]+s[left+1:]) || isValidPalindrome(s[:right]+s[right+1:])
		}
		left++
		right--
	}

	return true
}

func main() {
	// fmt.Println(stringsEqual("ab##", "c#d#"))
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
	// fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
	// fmt.Println(isValidPalindrome("A man, A plan, A Canal: Panama"))
	// fmt.Println(isValidPalindrome("Race car"))
	// fmt.Println(isValidPalindrome("aldlfakjdfia"))
	fmt.Println(isAlmostPalindrome("RaceACar"))
	fmt.Println(isAlmostPalindrome("abccdba"))
	fmt.Println(isAlmostPalindrome("abcdefdba"))
	fmt.Println(isAlmostPalindrome(""))
	fmt.Println(isAlmostPalindrome("a"))
	fmt.Println(isAlmostPalindrome("ab"))
}
