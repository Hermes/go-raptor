package partition

import (
		"math"
		"fmt"
)

/*
Partition() takes a pair of integers I, J as input and returns four integers as
output as defined by section 5.3.1.2 of the raptor specification
*/
func Partition(I int, J int) (int, int, int, int){
		IJ := float64(I) / float64(J)
		IL := math.Ceil(IJ)
		IS := math.Floor(IJ)
		JL := float64(I) - IS*float64(J)
		JS := float64(J) - JL
		return int(IL), int(IS), int(JL), int(JS)
}
