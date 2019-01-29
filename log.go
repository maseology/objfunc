package objfunc

import (
	"math"

	"github.com/maseology/mmaths"
)

func logTransform(o, s []float64) (_, _ []float64) {
	var x []int
	for i := range o {
		if o[i] <= 0. || s[i] <= 0. || math.IsNaN(o[i]) || math.IsNaN(s[i]) {
			x = append(x, i)
		} else {
			o[i] = math.Log10(o[i])
			s[i] = math.Log10(s[i])
		}
	}
	mmaths.Rev(x)
	for _, i := range x {
		s = append(o[:i], o[i+1:]...)
		s = append(s[:i], s[i+1:]...)
	}
	return o[:len(o)], s[:len(s)]
}
