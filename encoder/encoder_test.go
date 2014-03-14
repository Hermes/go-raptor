package encoder

import (
	"testing"
)

func TestPartition(t *testing.T) {
	I := 16
	J := 42
	IL, IS, JL, JS := Partition(I, J)
	if IL != 1 {
		t.Errorf("IL calculated improperly %d instead of 1", IL)
	}
	if IS != 0 {
		t.Errorf("IS calculated improperly %d instead of 0", IL)
	}
	if JL != 16 {
		t.Errorf("JL calculated improperly %d instead of 40", JL)
	}
	if JS != 26 {
		t.Errorf("JS calculated improperly %d instead of 26", JS)
	}
}

/*
This is more a sanity test than a proper test, check that the blocks have
different symbols. and that all the symbols aren't the same
*/
func TestEncoder(t *testing.T) {
	result := Block("test")
	blockA := result.blocks[0]
	blockB := result.blocks[1]
	if blockA.symbols[1].dat[1] == blockB.symbols[1].dat[1] {
		t.Error("Something is wrong symbols are matching between blocks")
	}
	if blockA.symbols[1].dat[1] == blockA.symbols[2].dat[1] {
		t.Error("Something is wrong symbols are matching in a block")
	}
}
