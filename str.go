package randstr

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVXYZ123456789")

const n = 32

// Get returns a pseudorandom string of length 32
func Get() string {
	result := make([]rune, n)

	for i := 0; i < n; i++ {
		result[i] = runes[rand.Intn(len(runes))]
	}

	return string(result)
}
