// The "Find All Anagrams in a String" problem involves searching for all occurrences of anagrams of a given
//non-empty string p within a larger string s. An anagram is a word or phrase formed by rearranging the letters
// of another, such as "cinema" and "iceman." In this problem, you are required to find all starting indices of
// anagrams of p in s.

// Problem Statement:
// Given a string s and a non-empty string p, find all the starting indices of anagrams of p in s.

// Example:
// plaintext
// Copy code
// Input: s = "cbaebabacd", p = "abc"
// Output: [0, 6]
// Explanation:
// The substring with starting index 0 is "cba," which is an anagram of "abc."
// The substring with starting index 6 is "bac," which is an anagram of "abc

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("-----Find all anagrams in string-----")
	s := "cbaebabacdbacb"
	p := "cba"
	p = sortString(p)
	var indices []int
	var subStr []string

	for i := 0; i <= len(s)-3; i++ {
		subStr = append(subStr, s[i:i+3])
		subStr[i] = sortString(subStr[i])
		if subStr[i] == p {
			indices = append(indices, i)
		}
	}
	fmt.Println("All possible anagram starting indices : ", indices)
}

func sortString(s string) string {
	//abcdef
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	char := strings.Split(s, "")
	sort.Strings(char)
	s = strings.Join(char, "")
	return s
}
