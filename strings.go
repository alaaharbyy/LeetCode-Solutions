package main

import "strings"

func StringSolutions(s1 string, s2 string, question int) any {
	switch question {
	case 1:
		//1768. Merge Strings Alternately
		return mergeAlternately(s1, s2)
	case 2:
		//1071. Greatest Common Divisor of Strings
		return gcdOfStrings(s1, s2)
	case 3:
		//345. Reverse Vowels of a String
		return reverseVowels(s2)
	case 4:
		//151. Reverse Words in a String
		return reverseWords(s2)
	case 5:
		//392. Is Subsequence
		return isSubsequence(s1, s2)
	}
	return "invalid"
}

//392. Is Subsequence
func isSubsequence(s string, t string) bool {
	s1 := []rune(s)
	s2 := []rune(t)
	var s3 []rune

	if len(s1) == 0 {
		return true
	}

	index := 0
	for _, char := range s2 {
		s3 = append(s3, char)
		if index < len(s1) {
			if s1[index] == char {
				strings.Replace(s, string(char), "", -1)
				s1 = []rune(s)
				index++
			}
		}

	}

	if index == len(s1) && string(s3) == t {
		return true
	} else {
		return false
	}

}

//1768. Merge Strings Alternately
func mergeAlternately(word1 string, word2 string) string {
	runes1 := []rune(word1)
	runes2 := []rune(word2)
	var temp []string
	var result string

	if len(runes1) > len(runes2) {
		for i, r := range runes1 {
			temp = append(temp, string(r))
			if i < len(runes2) {
				temp = append(temp, string(runes2[i]))
			}
		}
	} else {
		for i, r := range runes2 {
			if i < len(runes1) {
				temp = append(temp, string(runes1[i]))
			}
			temp = append(temp, string(r))

		}
	}

	result = strings.Join(temp, "")

	return result

}

//1071. Greatest Common Divisor of Strings
func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	return gcdOfStrings(str1[len(str2):], str2)
}

//345. Reverse Vowels of a String
func reverseVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	runes := []rune(s)
	var indices []int

	for i, char := range s {
		for _, v := range vowels {
			if char == v {
				indices = append(indices, i)
			}
		}
	}

	i := 0
	j := len(indices) - 1

	for i < j {
		runes[indices[i]], runes[indices[j]] = runes[indices[j]], runes[indices[i]]
		j--
		i++
	}

	return (string(runes))
}

//151. Reverse Words in a String
func reverseWords(s string) string {
	var distinationList []string
	arr := strings.Split(s, " ")

	for _, char := range arr {
		if len(char) != 0 {
			distinationList = append(distinationList, char)
		}
	}

	j := len(distinationList) - 1
	i := 0
	for i < j {
		distinationList[i], distinationList[j] = distinationList[j], distinationList[i]
		i++
		j--
	}
	return strings.Join(distinationList, " ")
}
