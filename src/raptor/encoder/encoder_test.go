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

/*
This is more a sanity test than a proper test, check that the blocks have
different symbols. and that all the symbols aren't the same
*/
func TestEncoder(t *testing.T) {
	result := Block("test")
	if result.blocks[0].symbols[1] == result.blocks[1].symbols[1]{
		t.Error("Something is wrong symbols are matching between blocks")
	}
	if result.blocks[0].symbols[1] == result.blocks[0].symbols[2]{
		t.Error("Something is wrong symbols are matching in a block")
	}
}
