package objfunc

import "math"

// PeakError returns the error in peak as a fraction (see HEC-HMS manual)
func PeakError(o, s []float64) float64 {
	so, ss := -9999., -9999.
	for i := range o {
		if math.IsNaN(o[i]) || math.IsNaN(s[i]) {
			continue
		}
		if o[i] > so {
			so = o[i]
		}
		if s[i] > ss {
			ss = s[i]
		}
	}
	return (ss - so) / so
}
