package leetcode

import (
	"fmt"
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

type FindAnagramsTest struct {
	arg1     string
	arg2     string
	expected []int
}

var FindAnagramsTests = []FindAnagramsTest{
	{arg1: "cbaebabacd", arg2: "abc", expected: []int{0, 6}},
	{arg1: "abab", arg2: "ab", expected: []int{0, 1, 2}},
}

func TestFindAnagrams(t *testing.T) {
	for i, test := range FindAnagramsTests {
		if result := findAnagrams(test.arg1, test.arg2); fmt.Sprint(result) != fmt.Sprint(test.expected) {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}

type CharacterReplacementTest struct {
	arg1     string
	arg2     int
	expected int
}

var CharacterReplacementTests = []CharacterReplacementTest{
	{arg1: "ABAB", arg2: 2, expected: 4},
	{arg1: "AABABBA", arg2: 1, expected: 4},
	{arg1: "ABBB", arg2: 1, expected: 4},
	{arg1: "BAAAB", arg2: 2, expected: 5},
}

func TestCharacterReplacement(t *testing.T) {
	for i, test := range CharacterReplacementTests {
		if result := characterReplacement(test.arg1, test.arg2); result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}

type GetHintTest struct {
	arg1     string
	arg2     string
	expected string
}

var GetHintTests = []GetHintTest{
	{arg1: "1807", arg2: "7810", expected: "1A3B"},
	{arg1: "1123", arg2: "0111", expected: "1A1B"},
	{arg1: "1122", arg2: "1222", expected: "3A0B"},
}

func TestGetHint(t *testing.T) {
	for i, test := range GetHintTests {
		if result := getHint(test.arg1, test.arg2); result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}

type BackspaceCompareTest struct {
	arg1     string
	arg2     string
	expected bool
}

var BackspaceCompareTests = []BackspaceCompareTest{
	{arg1: "ab#c", arg2: "ad#c", expected: true},
	{arg1: "ab##", arg2: "c#d#", expected: true},
	{arg1: "a#c", arg2: "b", expected: false},
	{arg1: "a##c", arg2: "#a#c", expected: true},
	{arg1: "bxj##tw", arg2: "bxo#j##tw", expected: true},
	{arg1: "xywrrmp", arg2: "xywrrmu#p", expected: true},
	{arg1: "bbbextm", arg2: "bbb#extm", expected: false},
	{arg1: "rheyggodcclgstf", arg2: "#rheyggodcclgstf", expected: true},
}

func TestBackspaceCompare(t *testing.T) {
	for i, test := range BackspaceCompareTests {
		if result := backspaceCompare(test.arg1, test.arg2); result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}

type DecodeStringTest struct {
	arg1     string
	expected string
}

var DecodeStringTests = []DecodeStringTest{
	{arg1: "3[a]2[bc]", expected: "aaabcbc"},
	{arg1: "3[a2[c]]", expected: "accaccacc"},
	{arg1: "2[abc]3[cd]ef", expected: "abcabccdcdcdef"},
}

func TestDecodeStringTests(t *testing.T) {
	for i, test := range DecodeStringTests {
		if result := decodeString(test.arg1); result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}
