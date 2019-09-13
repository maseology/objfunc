package objfunc

import "math"

// RMSE is the root mean squared errors
func RMSE(o, s []float64) float64 {
	c, n := 0., 0.
	for i := range o {
		if !math.IsNaN(s[i]) && !math.IsNaN(o[i]) {
			n += math.Pow(s[i]-o[i], 2.)
			c++
		}
	}
	return math.Sqrt(n / c)
}
