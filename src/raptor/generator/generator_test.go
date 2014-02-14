package generator

import (
	"testing"
)

func TestRand(t *testing.T) {
	if Rand(1, 1, 539210) != 526998 {
		t.Errorf("526998 != %d", Rand(1, 1, 539210))
	}

	if Rand(1, 2, 539210) != 265977 {
		t.Errorf("265977 != %d", Rand(1, 2, 539210))
	}
}

func TestDeg(t *testing.T) {
	if Deg(0) != 1 {
		t.Errorf("1 != %d", Deg(0))
	}
	if Deg(937888) != 10 {
		t.Errorf("10 != %d", Deg(937888))
	}
	if Deg(1048575) != 30 {
		t.Errorf("30 != %d", Deg(1048575))
	}
}
