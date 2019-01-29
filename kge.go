package objfunc

import (
	"math"
)

// KGE the Kling-Gupta efficiency measure
// ref: Gupta, H.V., H. Kling, K.K. Yilmaz, G.F. Martinez, 2009. Decomposition of the mean squared error and NSE performance criteria: Implications for improving hydrological modelling. Journal of Hydrology 377: 80-91.
// uniformly scaled between correlation, variability and bias
// optimal maximization to 1.0
func KGE(o, s []float64) float64 {
	return KGEscaled(o, s, 1., 1., 1.)
}

// LogKGE KGE of log transformed data
func LogKGE(o, s []float64) float64 {
	o, s = logTransform(o, s)
	return KGEscaled(o, s, 1., 1., 1.)
}

// KGEscaled : the Kling-Gupta efficiency measure
// sr: correlation scale factor,
// sa: variability scale factor, and
// sb: bias scale factor
// optimal maximization to 1.0
func KGEscaled(o, s []float64, sr, sa, sb float64) float64 {
	ms, ss := meansd(s)
	mo, so := meansd(o)
	r := rm(o, s, mo, ms)
	if math.IsNaN(r) {
		return math.NaN()
	}

	g1 := math.Pow(sr*(r-1.), 2.)
	g2 := math.Pow(sa*(ss/so-1.), 2.)
	g3 := math.Pow(sb*(ms/mo-1.), 2.)

	return 1. - math.Sqrt(g1+g2+g3)
}
