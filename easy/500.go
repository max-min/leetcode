package easy

import (
	"strings"
)

func FindWords(words []string) []string {

	arrs := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

	var ret []string
	for _, word := range words {
		strings.ToLower(word)

		var times int
		for _, arr := range arrs {
			if strings.ContainsAny(arr, word) {
				times++
			}
		}
		if times < 2 {
			ret = append(ret, word)
		}
	}

	return ret

}
