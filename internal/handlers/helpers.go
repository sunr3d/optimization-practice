package httphandlers

import "math"

func roundToTwo(val float64) float64 {
	return math.Round(val*100) / 100
}
