package generator

import (
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
linear search such that f[d-1] <= v <= f[d] TODO take min(d, W-2) also maybe
binary search?.
*/
func Deg(v uint) int {
	for i := 1; i <= 30; i++ {
		if f[i-1] <= v && v < f[i] {
			return i
		}
	}
	return int(f[30])
}
