package randstr

import "testing"

func TestRandString(t *testing.T) {
	n := 10
	result, err := Get(n)

	if err != nil {
		t.Errorf("Correct input %d gave error %v", n, err)
	}

	if len(result) != n {
		t.Error("Result string had wrong length")
	}

	n = 0
	_, err = Get(n)
	if err == nil {
		t.Errorf("Input %d didn't produce error", n)
	}

	n = -1
	_, err = Get(n)
	if err == nil {
		t.Errorf("Input %d didn't produce error", n)
	}
}
