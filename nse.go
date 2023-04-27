package objfunc

import (
	"log"
	"math"
)

// NSE : the Nash-Sutcliffe efficiency measure
// optimal maximization to 1.0
func NSE(o, s []float64) float64 {
	return NSEpow(o, s, 2.)
}

// NSEpow : the Nash-Sutcliffe efficiency measure
// p: power
// optimal maximization to 1.0
func NSEpow(o, s []float64, p float64) float64 {
	var n, d float64
	om, _ := Meansd(o)
	if len(o) != len(s) {
		log.Fatalf("NSEpow array lengths not matching %d %d\n", len(o), len(s))
	}
	for i := range o {
		if !math.IsNaN(s[i]) && !math.IsNaN(o[i]) {
			n += math.Pow(s[i]-o[i], p)
			d += math.Pow(o[i]-om, p)
		}
	}
	return 1. - n/d
}

// NNSE rescales the NSE form [-Inf,1] to [-1,-1], where 0 has the same meaning as NSE=0
// modifed from: Nossent and Bauwens, 2012: Application of a normalized Nash-Sutcliffe efficiency to improve the accuracy of the Sobol’ sensitivity analysis of a hydrological mode. Geophysical Research Abstracts, Vol. 14.
func NNSE(o, s []float64) float64 {
	// return 1. / (2. - NSEpow(o, s, 2.))  // original form: [0,-1], where .5 has the same meaning as NSE=0
	return 2./(2.-NSEpow(o, s, 2.)) - 1. // modified form
}

func NSEsmooth(o, s []float64, window int) float64 {
	nn, fn := len(o)-window, float64(window)
	oo, ss := make([]float64, nn), make([]float64, nn)
	for i := 0; i < nn; i++ {
		for j := 0; j < window; j++ {
			oo[i] += o[i+j]
			ss[i] += s[i+j]
		}
		oo[i] /= fn
		ss[i] /= fn
	}
	return NSEpow(oo, ss, 2.)
}
