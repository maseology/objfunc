package objfunc

import "math"

// NSE : the Nash-Sutcliffe effeciency measure
// optimial maximization to 1.0
func NSE(o []float64, s []float64) float64 {
	return NSEpow(o, s, 2.)
}

// NSEpow : the Nash-Sutcliffe effeciency measure
// p: power
// optimial maximization to 1.0
func NSEpow(o []float64, s []float64, p float64) float64 {
	var n, d float64
	om, _ := meansd(o)
	for i := range o {
		if s[i] != -9999. && o[i] != -9999. {
			n += math.Pow(s[i]-o[i], p)
			d += math.Pow(o[i]-om, p)
		}
	}
	return 1. - n/d
}
