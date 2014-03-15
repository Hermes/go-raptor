package encoder

import (
	"log"
	"os"
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
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	info, err := os.Stat("test")
	if err != nil {
		log.Fatal(err)
	}
	size := info.Size()
	result := Block(file, size)
	blockA := result.blocks[0]
	blockB := result.blocks[1]
	if blockA.symbols[1].dat[1] == blockB.symbols[1].dat[1] {
		t.Error("Something is wrong symbols are matching between blocks")
	}
	if blockA.symbols[1].dat[1] == blockA.symbols[2].dat[1] {
		t.Error("Something is wrong symbols are matching in a block")
	}
}
