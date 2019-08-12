package randstr

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVXYZ123456789")

const lenght = 32

// Get returns a pseudorandom string of length 32
func Get() string {
	result := make([]rune, lenght)

	for index := 0; index < lenght; index++ {
		result[index] = runes[rand.Intn(len(runes))]
	}

	return string(result)
}
