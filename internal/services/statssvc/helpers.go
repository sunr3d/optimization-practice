package statssvc

import "sort"

// func getSum(data []float64) float64 {
// 	res := 0.0

// 	for _, val := range data {
// 		res += val
// 	}

// 	return res
// }

// func getMinMax(data []float64) (float64, float64) {
// 	min, max := data[0], data[0]

// 	for _, val := range data {
// 		if val < min {
// 			min = val
// 		}
// 		if val > max {
// 			max = val
// 		}
// 	}

// 	return min, max
// }

func getSumMinMax(data []float64) (float64, float64, float64) {
	sum := 0.0
	min, max := data[0], data[0]

	for _, val := range data {
		sum += val
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return sum, min, max
}

func getMedian(data []float64) float64 {
	n := len(data)
	mid := n / 2

	sort.Float64s(data)

	if n%2 == 0 {
		return (data[mid-1] + data[mid]) / 2
	}

	return data[mid]
}

func getVariance(data []float64, mean float64) float64 {
	res := 0.0

	for _, val := range data {
		diff := val - mean
		res += diff * diff
	}

	return res / float64(len(data))
}
