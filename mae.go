package objfunc

import "math"

// MAE is the mean absolute errors
func MAE(o, s []float64) float64 {
	c, n := 0., 0.
	for i := range o {
		if !math.IsNaN(s[i]) && !math.IsNaN(o[i]) {
			n += math.Abs(s[i] - o[i])
			c++
		}
	}
	return n / c
}
