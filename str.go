package randstr

import (
	"fmt"
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const n = 32

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Get returns a pseudorandom string of length 32 from [a-zA-Z0-9]
func Get(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("%d is not a valid length", n)
	}

	result := make([]rune, n)

	for i := 0; i < n; i++ {
		result[i] = runes[rand.Intn((len(runes)))]
	}

	return string(result), nil
}
