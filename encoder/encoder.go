package encoder

import (
	"bufio"
	"github.com/wmak/go-raptor/constants"
	"github.com/wmak/go-raptor/generator"
	"log"
	"math"
	"os"
)

//As defined by section 4.4.1.1
type Source struct {
	blocks []SourceBlock
}

type SourceBlock struct {
	SBN int
	sub [][]SourceSymbol
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
TODO Convert filename parameter to Reader
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
	Kt := int(math.Ceil(F / float64(constants.T)))

	// Determining the value of Z
	SS := 4 // TODO findout how to get SS
	N_max := constants.Al * SS
	Z := int(math.Ceil(float64(Kt) / float64(generator.KL(uint32(N_max)))))

	// Determining the value of N
	N := N_max
	for i := 1; i <= N_max; i++ {
		if math.Ceil(float64(Kt)/float64(Z)) <= float64(generator.KL(uint32(i))) {
			N = i
			break
		}
	}

	KL, KS, ZL, ZS := Partition(Kt, Z)
	TL, TS, NL, NS := Partition(constants.T/constants.Al, N)
	blocks := make([]SourceBlock, 0)
	sub := make([][]SourceSymbol, 0)

	// Partition the object into the first ZL blocks of KL source symbols of T
	// octets
	for i := 0; i < ZL; i++ {
		symbol := make([]SourceSymbol, 0)
		for j := 0; j < KL; j++ {
			current := SourceSymbol{
				ESI: j,
				dat: make([]byte, constants.T), // create a buffer of size T
			}
			_, err := reader.Read(current.dat) // fill the buffer from the file
			if err != nil {
				log.Fatal(err) // report error on failure
			}
			symbol = append(symbol, current)
		}
		blocks = append(blocks, SourceBlock{
			SBN: i,
			symbols: symbol,
			sub: sub,
		})
	}

	// Partition the object into the next ZS blocks of KS source symbols of T
	// octets
	for i := 0; i < ZS; i++ {
		symbol := make([]SourceSymbol, 0) //to store the KS source symbols
		for j := 0; j < KS; j++ {
			current := SourceSymbol{
				ESI: j,
				dat: make([]byte, constants.T),
			}
			_, err := reader.Read(current.dat)
			if err != nil {
				log.Fatal(err)
			}
			symbol = append(symbol, current)
		}
		blocks = append(blocks, SourceBlock{
			SBN: ZL + i,
			symbols: symbol,
			sub: sub,
		})
	}

	// If Kt*T > F then for encoding purposes the last symbol of the last source
	// block must be padded with KT*T-F zero octets
	if Kt*constants.T > int(F) {
		padding := blocks[ZL+ZS-1].symbols
		current := SourceSymbol{
			ESI: KS + KL,
			dat: make([]byte, Kt*constants.T-int(F)),
		}
		padding = append(padding, current)
	}

	// Divide each source block with K source symbols into N sub-blocks each
	// consisting of k contiguous sub symbols of size TL*AJ octets
	for i := 0; i < Z; i++ {
		block := blocks[i]
		for j := 0; j < NL; j++{
			newsub := make([]SourceSymbol, 0)
			block.sub = append(block.sub, newsub)
		}
		for j := 0; j < NS; j++{
		}
	}
	source := Source{
		blocks: blocks,
	}
	return source
}

/*
each source block with K source symbols MUST be divided into N = NL + NS 
contiguous sub-blocks, the first NL sub-blocks each consisting of K contiguous 
sub-symbols of size of TL*Al octets and the remaining NS sub-blocks each 
consisting of K contiguous sub-symbols of size of TS*Al octets. The symbol 
alignment parameter Al ensures that sub-symbols are always a multiple of Al
 octets.

*/
