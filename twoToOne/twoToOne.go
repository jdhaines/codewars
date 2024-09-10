package twoToOne

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	// Hash Map to store the unique string of characters
	result := ""
	hashMap := make(map[string]bool)

	for _, char := range s1 {
		if !hashMap[string(char)] {
			hashMap[string(char)] = true
			result = fmt.Sprintf("%s%s", result, string(char))
		}
	}
	for _, char := range s2 {
		if !hashMap[string(char)] {
			hashMap[string(char)] = true
			result = fmt.Sprintf("%s%s", result, string(char))
		}
	}

	// Sort the result string
	return sortString(result)
}

func sortString(uglyString string) string {
	prettyString := strings.Split(uglyString, "")
	sort.Strings(prettyString)
	return strings.Join(prettyString, "")
}

func wishIWasSmart(s1 string, s2 string) string {
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s1+s2, char) {
			result += char
		}
	}
	return result
}

func wishIWasSmart2(s1 string, s2 string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	res := ""

	for _, a := range alphabet {
		if strings.ContainsRune(s1, a) || strings.ContainsRune(s2, a) {
			res += string(a)
		}
	}

	return res
}
