package strings

// ReverseString
// https://leetcode.com/problems/reverse-string/
func ReverseString(s []byte) {
	var min, max = 0, len(s) - 1
	for min < max {
		s[min], s[max] = s[max], s[min]
		min++
		max--
	}
}
