package objfunc

import "math"

// NSE : the Nash-Sutcliffe efficiency measure
// optimal maximization to 1.0
func NSE(o, s []float64) float64 {
	return NSEpow(o, s, 2.)
}

// NSEi "overloads" KGE
func NSEi(oi, si []interface{}) float64 {
	o, s := make([]float64, len(oi)), make([]float64, len(si))
	for i := range oi {
		o[i] = oi[i].(float64)
		s[i] = si[i].(float64)
	}
	return NSEpow(o, s, 2.)
}

// NSEpow : the Nash-Sutcliffe efficiency measure
// p: power
// optimal maximization to 1.0
func NSEpow(o, s []float64, p float64) float64 {
	var n, d float64
	om, _ := meansd(o)
	for i := range o {
		if !math.IsNaN(s[i]) && !math.IsNaN(o[i]) {
			n += math.Pow(s[i]-o[i], p)
			d += math.Pow(o[i]-om, p)
		}
	}
	return 1. - n/d
}
