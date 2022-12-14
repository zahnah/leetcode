package leetcode

import (
	"testing"
)

type ReverseStringTest struct {
	arg1     []byte
	expected string
}

var ReverseStringTests = []ReverseStringTest{
	{arg1: []byte{'h', 'e', 'l', 'l', 'o'}, expected: "olleh"},
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

type ReverseWordsTest struct {
	arg1     string
	expected string
}

var ReverseWordsTests = []ReverseWordsTest{
	{arg1: "Let's take LeetCode contest", expected: "s'teL ekat edoCteeL tsetnoc"},
}

// ReverseString
func TestReverseWords(t *testing.T) {
	for _, test := range ReverseWordsTests {
		if result := ReverseWords(test.arg1); result != test.expected {
			t.Errorf("Output %q not equal to expected %q", result, test.expected)
		}
	}
}

type LengthOfLongestSubstringTest struct {
	arg1     string
	expected int
}

var LengthOfLongestSubstringTests = []LengthOfLongestSubstringTest{
	{arg1: "abcabcbb", expected: 3},
	{arg1: "bbbbb", expected: 1},
	{arg1: "pwwkew", expected: 3},
	{arg1: " ", expected: 1},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for i, test := range LengthOfLongestSubstringTests {
		if result := lengthOfLongestSubstring(test.arg1); result != test.expected {
			t.Errorf("%d: Output %d not equal to expected %d", i, result, test.expected)
		}
	}
}

type CheckInclusionTest struct {
	arg1     string
	arg2     string
	expected bool
}

var CheckInclusionTests = []CheckInclusionTest{
	{arg1: "ab", arg2: "eidbaooo", expected: true},
	{arg1: "ab", arg2: "eidboaoo", expected: false},
	{arg1: "ab", arg2: "a", expected: false},
	{arg1: "ab", arg2: "ab", expected: true},
	{arg1: "abc", arg2: "bbbca", expected: true},
}

func TestCheckInclusion(t *testing.T) {
	for i, test := range CheckInclusionTests {
		if result := checkInclusion(test.arg1, test.arg2); result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}
