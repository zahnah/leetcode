package stringset

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
