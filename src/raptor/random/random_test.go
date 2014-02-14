package random

import (
	"testing"
)

func TestRand(t *testing.T) {
	if Rand(1, 1, 539210) != 526998 {
		t.Errorf("526998 != %d", Rand(1, 1, 539210))
	}

	if Rand(1, 2, 539210) != 526998 {
		t.Errorf("526998 != %d", Rand(1, 2, 539210))
	}
}
