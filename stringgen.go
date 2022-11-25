package stringgen

import (
	"math"
)

func GenerateAllStringsUpToLength(chars []rune, maxLength float64, c chan string) {
	GenerateAllStringsUpToLengthWithStartString(chars, maxLength, "", c)
}

func GenerateAllStringsUpToLengthWithStartString(chars []rune, maxLength float64, start string, c chan string) {
	var charsLen = len(chars)

	var i float64
	var max float64 = 0
	for i = 1; i <= maxLength; i++ {
		max += math.Pow(float64(charsLen), i)
	}

	var startInt = 0
	if start != "" {
		startInt = GetStringNumber(chars, start)
	}

	for v := startInt; v < int(max); v++ {
		tmpV := v
		var res string

		for {
			if tmpV < charsLen {
				res = res + string(chars[tmpV])
				break
			} else {
				remainder := tmpV % charsLen
				toDivide := tmpV - remainder
				res = res + string(chars[remainder])
				tmpV = (toDivide / charsLen) - 1
			}
		}

		c <- res
	}
}

// GetStringNumber takes a string and a list of characters and returns the number that would've generated that string using the GenerateAllStringsUpToLength function
func GetStringNumber(chars []rune, str string) int {
	var charsLen = len(chars)

	var res int
	var i float64
	for _, c := range str {
		var charIndex int
		for j, char := range chars {
			if c == char {
				charIndex = j
				break
			}
		}

		if i == 0 {
			res = charIndex
		} else {
			res += (1 + charIndex) * int(math.Pow(float64(charsLen), i))
		}

		i++
	}

	return res
}
