package lib

import (
	"math"
	"strconv"
	"strings"
)

func convertToAscii(a rune) int {
	return int(a)
}

func convertToBinary(a int) string {
	var result [8]byte
	for i := 7; i >= 0; i-- {
		result[i] = '0' + byte(a%2)
		a /= 2
	}
	return string(result[:])
}
func cekIs6Bit(input string) string {
	for len(input) < 6 {
		input += "0"
	}
	return input
}

func splitInto6Bits(input string) string {
	var result string
	for i := 0; i < len(input); i += 6 {
		end := i + 6
		if end > len(input) {
			end = len(input)
		}
		if len(result) == 0 {
			result += cekIs6Bit(input[i:end])
		} else {
			result += " " + cekIs6Bit(input[i:end])
		}
	}
	return result
}

func convertToChar(input string) int {
	dec := 0
	for i, c := range input {
		z, _ := strconv.Atoi(string(c))
		dec += z * int(math.Pow(2, float64(len(input)-i-1)))
	}
	return dec
}

func padding(word string) string {
	if len(word)%3 == 2 {
		return "="
	} else if len(word)%3 == 1 {
		return "=="
	} else {
		return ""
	}

}

func Encode(word string) string {
	if word != "" {
		binaryString := ""

		for _, c := range word {
			binaryString += convertToBinary(convertToAscii(c))
		}
		binary6Bit := splitInto6Bits(binaryString)
		arrayBinary6Bit := strings.Split(binary6Bit, " ")
		finalBase64 := ""
		for _, b := range arrayBinary6Bit {
			finalBase64 += Map[convertToChar(b)]
		}
		return finalBase64 + padding(word)
	}
	return ""
}
