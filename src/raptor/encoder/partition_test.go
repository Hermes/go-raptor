package encoder

import (
	"testing"
)

func TestPartition(t *testing.T) {
	IL, IS, JL, JS := Partition(10, 20)
	if IL != 1 {
		t.Error("IL calculated improperly")
	}
	if IS != 0 {
		t.Error("IS calculated improperly")
	}
	if JL != 10 {
		t.Error("JL calculated improperly")
	}
	if JS != 10 {
		t.Error("JS calculated improperly")
	}
}
