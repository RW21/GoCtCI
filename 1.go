package main

import "strings"

func isUniqueChar(s string) bool {
	var check [128] bool
	for i := range s {
		if check[s[i]] {
			return false
		}
		check[s[i]] = true
	}
	return true
}

func isPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var check_a [128] int

	for i := range a {
		check_a[a[i]]++
	}

	for i := range b {
		check_a[b[i]]--
		if check_a[b[i]] < 0 {
			return false
		}
	}
	return true
}

func URLify(chararray string) string {
	var s strings.Builder

	for _, c := range chararray {
		if c == ' ' {
			s.WriteString("%20")
		} else {
			s.WriteRune(c)
		}
	}
	return s.String()
}

func isPalindromePermutation(s string) bool {
	hashmap := make(map[rune]bool)
	for _, c := range s {
		if c != ' ' {
			hashmap[c] = !hashmap[c]
		}

	}
	// count hashmap
	oddCount := 0
	if len(s)%2 == 0 {
		for _, v := range hashmap {
			if v {
				oddCount++
			}
		}
		if oddCount == 1 {
			return true
		}
	} else {
		for _, v := range hashmap {
			if v {
				oddCount++
			}
		}
		if oddCount == 0 {
			return true
		}
	}

	return false
}


func main() {

}
