package random

import "math/rand"

var (
	// Numbers 数字
	Numbers = "0123456789"
	// UpperLetters 大写字母
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// LowerLetters 小写字母
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	// Letters 字母
	Letters = UpperLetters + LowerLetters
)

func Int(min, max int) int {
	return min + rand.Intn(max-min)
}

func Number(length int) string {
	return String(length, Numbers)
}

func LowerLetter(length int) string {
	return String(length, LowerLetters)
}

func UpperLetter(length int) string {
	return String(length, UpperLetters)
}

func Letter(length int) string {
	return String(length, Letters)
}

func String(length int, chars string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[Int(0, len(chars))]
	}
	return string(b)
}
