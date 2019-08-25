package randstr

import (
	"math"
	"testing"
)

func TestRandString(t *testing.T) {
	n := math.MaxInt64
	if len(Get(n)) != 0 {
		t.Errorf("Wrong length: n = math.MaxInt64")
	}

	n = 1
	if len(Get(n)) != n {
		t.Errorf("Wrong length: n = 1")
	}

	n = 0
	if len(Get(n)) != 0 {
		t.Errorf("Wrong length: n = 0")
	}

	n = math.MinInt64
	if len(Get(n)) != 0 {
		t.Errorf("Wrong length: n = math.MinInt64")
	}
}
