package encoder

import (
	"bufio"
	"log"
	"math"
	"os"
)

const Al int = 4     //alignment variable, 4 as per recommended by 4.3
const T int = Al * 2 //the symbol size

//As defined by section 4.4.1.1
type Source struct {
	blocks []SourceBlock
}

type SourceBlock struct {
	SBN     int
	symbols []SourceSymbol
}

type SourceSymbol struct {
	ESI int
	dat []byte
}

/*
Block as defined in section 4.4.1.2  will generate the source blocks which later
will be split into source symbols. The object will be partitioned by running the
partition function on ceil(filesize/T) and using the resulting parameters to
determine how to block
*/

func Block(filename string) Source {
	// open up the file
	file, err := os.Open(filename)
	info, _ := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	// Determining how to partition the source
	F := float64(info.Size())
	Kt := int(math.Ceil(F / float64(T)))
	KL, KS, ZL, ZS := Partition(Kt, 5) // TODO find out how to get Z
	blocks := make([]SourceBlock, 0)   //TODO should be Z...
	// Partition the object into the first ZL blocks of KL source symbols of T
	// octets
	for i := 0; i < ZL; i++ {
		symbol := make([]SourceSymbol, 0) //to store the KL source symbols
		for j := 0; j < KL; j++ {
			current := SourceSymbol{
				ESI: j,
				dat: make([]byte, T),
			}
			_, err := reader.Read(current.dat)
			if err != nil {
				log.Fatal(err)
			}
			symbol = append(symbol, current)
		}
		blocks = append(blocks, SourceBlock{
			SBN:     i,
			symbols: symbol,
		})
	}
	// Partition the object into the next ZS blocks of KS source symbols of T
	// octets
	for i := 0; i < ZS; i++ {
		symbol := make([]SourceSymbol, 0) //to store the KS source symbols
		for j := 0; j < KS; j++ {
			current := SourceSymbol{
				ESI: j,
				dat: make([]byte, T),
			}
			_, err := reader.Read(current.dat)
			if err != nil {
				log.Fatal(err)
			}
			symbol = append(symbol, current)
		}
		blocks = append(blocks, SourceBlock{
			SBN:     ZL + i,
			symbols: symbol,
		})
	}
	return Source{
		blocks: blocks,
	}
}
