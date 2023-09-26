package objfunc

import "math"

// PkRMSE is the Peak-weighted root mean square error objective function (see HEC-HMS manual)
// US Army Corps of Engineers, USACE (1998). HEC-1 flood hydrograph package user's manual. Hydrologic Engineering Center, Davis, CA.
func PkRMSE(o, s []float64) float64 {
	c, n := 0., 0.
	om, _ := Meansd(o)
	for i := range o {
		if !math.IsNaN(s[i]) && !math.IsNaN(o[i]) {
			n += math.Pow(s[i]-o[i], 2.) * (o[i] + om) / 2. / om
			c++
		}
	}
	return math.Sqrt(n / c)
}
