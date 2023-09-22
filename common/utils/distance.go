package utils

import "math"

func EDistance(v1, v2 []float32) float32 {
	subV := subtract(v1, v2)
	squV := square(subV)
	sumV := sum(squV)
	ret := math.Sqrt(float64(sumV))
	return float32(ret)
}

func subtract(v1 []float32, v2 []float32) []float32 {
	var res []float32
	var n int = len(v2)
	for i := 0; i < n; i++ {
		res = append(res, v1[i]-v2[i])
	}
	return res
}

func square(v []float32) []float32 {
	var res []float32
	var n int = len(v)
	for i := 0; i < n; i++ {
		res = append(res, v[i]*v[i])
	}
	return res
}

func sum(slice []float32) float32 {
	var result float32 = 0
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result
}
