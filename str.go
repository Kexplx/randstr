package randstr

import (
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const n = 32

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Get returns a pseudorandom string of max length n = 500
// Letter Range [a-zA-Z0-9]
func Get(n int) string {
	if n <= 0 || n > 500 {
		return ""
	}

	result := make([]rune, n)

	for i := 0; i < n; i++ {
		result[i] = runes[rand.Intn((len(runes)))]
	}

	return string(result)
}
