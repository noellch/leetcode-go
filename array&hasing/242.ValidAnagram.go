/* Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once. */

package main

func main() {

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	arr := [26]int{}

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, char := range arr {
		if char != 0 {
			return false
		}
	}

	return true

}
