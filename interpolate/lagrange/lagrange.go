package lagrange

import (
	"fmt"

	"github.com/DzananGanic/numericalgo/interpolate"
)

// Lagrange provides the basic functionality for lagrange interpolation.
// Given X and Y float64 slices, it can estimate the value of the function at the desired point.
type Lagrange struct {
	interpolate.Base
}

// New returns the new Lagrange object.
func New() *Lagrange {
	lg := &Lagrange{}
	return lg
}

func (lg *Lagrange) Interpolate(val float64) float64 {
	var est float64

	for i := 0; i < len(lg.X); i++ {
		prod := lg.Y[i]
		for j := 0; j < len(lg.X); j++ {
			if i != j {
				prod = prod * (val - lg.X[j]) / (lg.X[i] - lg.X[j])
			}
		}
		est += prod
	}

	return est
}

func (lg *Lagrange) Validate(val float64) error {

	// TODO: Check case where lg.x[i]-lg.x[j] is 0

	if val < lg.XYPairs[0].X {
		return fmt.Errorf("Value to interpolate is too small and not in range")
	}

	if val > lg.XYPairs[len(lg.XYPairs)-1].X {
		return fmt.Errorf("Value to interpolate is too large and not in range")
	}

	return nil
}
