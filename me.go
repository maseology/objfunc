package objfunc

import "math"

// ME is the mean error
func ME(o, s []float64) float64 {
	c, n := 0., 0.
	for i := range o {
		if !math.IsNaN(s[i]) && !math.IsNaN(o[i]) {
			n += s[i] - o[i]
			c++
		}
	}
	return n / c
}
