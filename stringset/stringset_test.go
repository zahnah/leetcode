package stringset

import (
	"testing"
)

type ReverseStringTest struct {
	arg1     []byte
	expected string
}

type ReverseWordsTest struct {
	arg1     string
	expected string
}

var ReverseStringTests = []ReverseStringTest{
	{arg1: []byte{'h', 'e', 'l', 'l', 'o'}, expected: "olleh"},
}

var ReverseWordsTests = []ReverseWordsTest{
	{arg1: "Let's take LeetCode contest", expected: "s'teL ekat edoCteeL tsetnoc"},
}

// ReverseString
func TestReverseString(t *testing.T) {
	for _, test := range ReverseStringTests {
		ReverseString(test.arg1)
		if string(test.arg1) != test.expected {
			t.Errorf("Output %q not equal to expected %q", string(test.arg1), test.expected)
		}
	}
}

// ReverseString
func TestReverseWords(t *testing.T) {
	for _, test := range ReverseWordsTests {

		if result := ReverseWords(test.arg1); result != test.expected {
			t.Errorf("Output %q not equal to expected %q", result, test.expected)
		}
	}
}
