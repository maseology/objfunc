package objfunc

import "math"

// Bias returns the mass balance as a fraction
func Bias(o, s []float64) float64 {
	so, ss := 0., 0.
	for i := range o {
		if math.IsNaN(o[i]) || math.IsNaN(s[i]) {
			continue
		}
		so += o[i]
		ss += s[i]
	}
	return (ss - so) / so
}

// Biasi "overloads" KGE
func Biasi(oi, si []interface{}) float64 {
	o, s := make([]float64, len(oi)), make([]float64, len(si))
	for i := range oi {
		o[i] = oi[i].(float64)
		s[i] = si[i].(float64)
	}
	return Bias(o, s)
}
