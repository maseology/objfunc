package objfunc

import "math"

func Standardize(x []float64) (y []float64, mean, sd float64) {
	y = make([]float64, len(x))
	mean, sd = Meansd(x)
	for i, v := range x {
		if math.IsNaN(v) {
			panic("Standardize error, must use clean data (no NaNs)")
		}
		y[i] = (v - mean) / sd
	}
	return
}

// rescales to (0,1)
func Unitize(x []float64) (y []float64) {
	y = make([]float64, len(x))
	mn, mx := minmax(x)
	for i, v := range x {
		if math.IsNaN(v) {
			panic("Unitize error, must use clean data (no NaNs)")
		}
		y[i] = (v - mn) / (mx - mn)
	}
	return
}

func RescaleLim(x []float64, lb, ub float64) (y []float64) {
	y = make([]float64, len(x))
	for i, v := range Unitize(x) {
		y[i] = v*(ub-lb) + lb
	}
	return
}
