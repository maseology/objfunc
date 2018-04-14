package objfunc

import (
	"math"
)

func meansd(d []float64) (float64, float64) {
	c, s, ss := 0, 0., 0.
	for i := 0; i < len(d); i++ {
		if d[i] != -9999. {
			s += d[i]
			c++
		}
	}
	s /= float64(c)
	for i := 0; i < len(d); i++ {
		if d[i] != -9999. {
			ss += math.Pow(d[i]-s, 2.)
		}
	}
	return s, math.Sqrt(ss / float64(c-1))
}

func r(x, y []float64) float64 {
	// see page 367 & 396 of Walpole Meyers Myers
	if len(x) != len(y) {
		panic("Coefficient of determination error: unequal array lengths")
	}
	c, mx, my := 0, 0., 0.
	for i := 0; i < len(x); i++ {
		if x[i] != -9999. && y[i] != -9999. {
			mx += x[i]
			my += y[i]
			c++
		}
	}
	mx /= float64(c)
	my /= float64(c)

	sxx, syy, sxy := 0., 0., 0.
	for i := 0; i < len(x); i++ {
		if x[i] != -9999. && y[i] != -9999. {
			sxx += math.Pow(x[i]-mx, 2.)     // variance in x
			syy += math.Pow(y[i]-my, 2.)     // variance in y
			sxy += (x[i] - mx) * (y[i] - my) // covariance
		}
	}

	if sxx > 0. && syy > 0. {
		return math.Copysign(sxy/math.Sqrt(sxx*syy), sxy)
	}
	return -9999.
}
