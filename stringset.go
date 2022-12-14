package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// ReverseString
// #344 https://leetcode.com/problems/reverse-string/
func ReverseString(s []byte) {
	var min, max = 0, len(s) - 1
	for min < max {
		s[min], s[max] = s[max], s[min]
		min++
		max--
	}
}

// ReverseWords
// #557 https://leetcode.com/problems/reverse-words-in-a-string-iii/
func ReverseWords(s string) string {
	words := []byte(s)
	var min, max, length = 0, 0, len(words)
	for max < length {
		for max < length && words[max] != ' ' {
			max++
		}
		for m := max - 1; min < m; {
			words[min], words[m] = s[m], s[min]
			min++
			m--
		}
		max++
		min = max
	}
	return string(words)
}

// lengthOfLongestSubstring
// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	var max, left, right int
	var length = len(s)
	for i := 0; i < length; i++ {
		for j := left; j < right; j++ {
			if s[j] == s[i] {
				if max < right-left {
					max = right - left
				}
				left = j + 1
				break
			}
		}
		right++
	}
	if max < right-left {
		max = right - left
	}
	return max
}

// checkInclusion
// 567. Permutation in String
// https://leetcode.com/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	var countS2, countS1 uint
	var fit bool
	lengthS1, lengthS2 := len(s1), len(s2)
	for i := 0; i < lengthS2; i++ {
		if lengthS2-i < lengthS1 {
			break
		}

		fit = true
		for j := 0; j < lengthS1; j++ {
			countS2, countS1 = 0, 0
			for jj := 0; jj < lengthS1; jj++ {
				if s2[i+jj] == s1[j] {
					countS2++
				}
				if s1[jj] == s1[j] {
					countS1++
				}
			}
			if countS1 != countS2 {
				fit = false
				break
			}
		}
		if fit {
			return true
		}
	}
	return false
}

// findAnagrams
// 438. Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	lengthS, lengthP := len(s), len(p)
	sMap := make(map[byte]int)
	pMap := make(map[byte]int)
	var result []int
	var leftPart int

	for i := 0; i < lengthP; i++ {
		pMap[p[i]]++
	}

	for i := 0; i < lengthS; i++ {
		leftPart = i - lengthP

		if leftPart >= 0 {
			sMap[s[leftPart]]--

			if sMap[s[leftPart]] == 0 {
				delete(sMap, s[leftPart])
			}
		}

		sMap[s[i]]++

		if len(pMap) == len(sMap) {
			isTheSame := true

			for key, pValue := range pMap {
				if sValue, ok := sMap[key]; ok {
					if sValue != pValue {
						isTheSame = false
						break
					}
				} else {
					isTheSame = false
					break
				}
			}

			if isTheSame {
				result = append(result, i+1-lengthP)
			}
		}
	}

	return result
}

// characterReplacement
// 424. Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	leftSide, maxLength, maxRepeatLetterCount, sLength := 0, 0, 0, len(s)
	sMap := make(map[byte]int)

	for rightSide := 0; rightSide < sLength; rightSide++ {
		rightChar := s[rightSide]
		sMap[rightChar]++

		maxRepeatLetterCount = max(maxRepeatLetterCount, sMap[rightChar])

		if rightSide-leftSide+1-maxRepeatLetterCount > k {
			leftChar := s[leftSide]
			sMap[leftChar]--
			leftSide++
		}

		maxLength = max(maxLength, rightSide-leftSide+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// getHint
// 299. Bulls and Cows
// https://leetcode.com/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	var bulls, caws int
	hash := make(map[int32]int)
	for _, s := range secret {
		hash[s]++
	}
	for i, s := range guess {
		if guess[i] == secret[i] {
			bulls++
			if hash[s] == 0 {
				caws--
			} else {
				hash[s]--
			}
		} else if j, ok := hash[s]; ok && j > 0 {
			hash[s]--
			caws++
		}
	}
	return fmt.Sprintf("%vA%vB", bulls, caws)
}

// backspaceCompare
// 844. Backspace String Compare
// https://leetcode.com/problems/backspace-string-compare/
func backspaceCompare(s string, t string) bool {
	var sIndex, tIndex = len(s) - 1, len(t) - 1
	for tIndex >= 0 || sIndex >= 0 {
		sIndex = backspaceCompareNextSymbol(&s, sIndex)
		tIndex = backspaceCompareNextSymbol(&t, tIndex)

		if (sIndex >= 0 && tIndex < 0) || (sIndex < 0 && tIndex >= 0) || (sIndex >= 0 && tIndex >= 0 && s[sIndex] != t[tIndex]) {
			return false
		}

		sIndex--
		tIndex--
	}

	return sIndex <= 0 && tIndex <= 0
}
func backspaceCompareNextSymbol(s *string, i int) int {
	for i >= 0 && (*s)[i] == '#' {
		jump := 1
		i--
		for i >= 0 && jump > 0 {
			if (*s)[i] == '#' {
				jump++
			} else {
				jump--
			}
			i--
		}
	}
	return i
}

// decodeString
// 394. Decode String
// https://leetcode.com/problems/decode-string/
func decodeString(s string) string {
	left, right, length := 0, 0, len(s)-1
	for right <= length {
		if s[right] == '[' {
			left = right
		}
		if s[right] == ']' {
			leftNum := left - 1
			for leftNum >= 0 && s[leftNum] >= '0' && s[leftNum] <= '9' {
				leftNum--
			}
			num, _ := strconv.Atoi(s[leftNum+1 : left])
			return decodeString(s[0:leftNum+1] + strings.Repeat(s[left+1:right], num) + s[right+1:length+1])
		}
		right++
	}
	return s
}
