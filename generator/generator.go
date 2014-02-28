package generator

import (
	"github.com/wmak/go-raptor/constants"
	"math"
)

/*
Random number generator as defined by section 5.3.5.1 and will use the arrays
in constant to determine it's value
*/
func Rand(y uint, i uint, m uint) uint32 {
	pow28 := math.Pow(2, 8)
	pow216 := math.Pow(2, 16)
	pow224 := math.Pow(2, 24)
	x0 := int(math.Mod(float64(y+i), pow28))
	x1 := int(math.Mod(math.Floor(float64(y)/pow28)+float64(i), pow28))
	x2 := int(math.Mod(math.Floor(float64(y)/pow216)+float64(i), pow28))
	x3 := int(math.Mod(math.Floor(float64(y)/pow224)+float64(i), pow28))
	return uint32(math.Mod(float64(V0[x0]^V1[x1]^V2[x2]^V3[x3]), float64(m)))
}

/*
Degree generator as defined by section 5.3.5.2 will find the index of f using
linear search such that constants.F[d-1] <= v <= constants.F[d].
TODO take min(d, W-2) also maybe binary search?.
*/
func Deg(v uint) int {
	for i := 1; i <= 30; i++ {
		if constants.F[i-1] <= v && v < constants.F[i] {
			return i
		}
	}
	return int(constants.F[30])
}

/*
Function KL as defined by sections 4.3 will find the maximum value of K' such
that K' <= constants.WS/(constants.Al*(ceil(constants.T/(constants.Al*n)))).
TODO write tests, confirm validity
*/
func KL(N_max uint32) int {
	var max uint32 = 0
	var r float64 = 0

	for n := 1; n <= int(N_max); n++ {
		r = float64(constants.WS) / (float64(constants.Al) * math.Ceil(float64(constants.T)/float64(constants.Al*n)))
		for k := 0; k < 477; k++ { // 477 is number of elements in K'
			if float64(constants.K[k]) <= r && uint32(constants.K[k]) > max {
				max = uint32(constants.K[k])
			}
		}
	}
	return int(max)
}
