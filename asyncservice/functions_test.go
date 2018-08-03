package asyncservice

import (
	"testing"
	"math"
)

//Test round() implementation
func TestRound(t * testing.T) {
	array := []float64{1.005, 1.005001, 1.0049999}

	for _, value := range array {
		res := Round2(value)
		if res != math.Round(value*100)/100 {
			t.Errorf("Incorrect rounding of %f, result is %f",value,res)
		}
	}
}
