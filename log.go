package objfunc

import (
	"math"

	"github.com/maseology/mmaths"
)

func logTransform(o, s []float64) (_, _ []float64) {
	var x []int
	ol, sl := make([]float64, len(o)), make([]float64, len(s))
	for i := range o {
		if o[i] <= 0. || s[i] <= 0. || math.IsNaN(o[i]) || math.IsNaN(s[i]) {
			x = append(x, i)
		} else {
			ol[i] = math.Log10(o[i])
			sl[i] = math.Log10(s[i])
		}
	}
	mmaths.Rev(x)
	for _, i := range x {
		ol = append(ol[:i], ol[i+1:]...)
		sl = append(sl[:i], sl[i+1:]...)
	}
	return ol[:len(ol)], sl[:len(sl)]
}
