package utils

import (
	"crypto/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func RandStringBytes() string {
	return time.Now().Format(time.RFC3339)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringGenerator(n int) string {
	ll := len(letterBytes)
	b := make([]byte, n)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < n; i++ {
		b[i] = letterBytes[int(b[i])%ll]
	}
	return string(b)
}

func RemoveAllSymbolsWithSpaces(str string) string {
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, " ")
	return str
}

func IntegerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

const otpChars = "1234567890"

func GenerateNumber(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
