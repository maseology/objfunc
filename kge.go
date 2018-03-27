package objfunc

import "math"

// KGE : the Kling-Gupta effeciency measure
// uniformly scaled between correlation, variablity and bias
// optimial maximization to 1.0
func KGE(o []float64, s []float64) float64 {
	return KGEscaled(o, s, 1., 1., 1.)
}

// KGEscaled : the Kling-Gupta effeciency measure
// sr: correlation scale factor,
// sa: variablity scale factor, and
// sb: bias scale factor
// optimial maximization to 1.0
func KGEscaled(o []float64, s []float64, sr, sa, sb float64) float64 {
	ms, ss := meansd(s)
	mo, so := meansd(o)
	r := r(o, s)
	g1 := math.Pow(sr*(r-1.), 2.)
	g2 := math.Pow(sa*(ss/so-1.), 2.)
	g3 := math.Pow(sr*(ms/mo-1.), 2.)
	return 1. - math.Sqrt(g1+g2+g3)
}
