package objfunc

import "math"

// Krause is a weighted r² efficiency criteria
// ref: Krause, P., D.P. Boyle, F. Base, 2005. Comparison of different efficiency criteria for hydrological model assessment. Advances in Geoscience 5: 89-97.
func Krause(o, s []float64) float64 {
	r, sxx, _, sxy := r(o, s)
	b := math.Abs(sxy / sxx) // pg 367 in Walpole Meyers Myers
	if math.IsNaN(r) {
		return math.NaN()
	}
	if b <= 1. {
		return b * r * r
	}
	return r * r / b
}

// Krausei "overloads" Krause
func Krausei(oi, si []interface{}) float64 {
	o, s := make([]float64, len(oi)), make([]float64, len(si))
	for i := range oi {
		o[i] = oi[i].(float64)
		s[i] = si[i].(float64)
	}
	return Krause(o, s)
}
