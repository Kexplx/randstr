package randstr

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const n = 32

// Get returns a pseudorandom string of length 32 from [a-zA-Z0-9]
func Get() string {
	result := make([]rune, n)

	for i := 0; i < n; i++ {
		result[i] = runes[rand.Intn(len(runes))]
	}

	return string(result)
}
