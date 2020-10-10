package lib

import (
	"strings"
)

var charset = [62]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// Base62Encode : Base62 Encode
func Base62Encode(n int) string {
	if n == 0 {
		return "0"
	}

	var result string

	for n > 0 {
		result = charset[n%62] + result
		n /= 62
	}

	return result
}

// MapTrim : Trim String Slice
func MapTrim(vs []string) []string {
	for i, v := range vs {
		vs[i] = strings.Trim(v, " ")
	}

	return vs
}
