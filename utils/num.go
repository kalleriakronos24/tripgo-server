package utils

import "math"

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func ToFixed32(num float32, precision int) float64 {
	var newNum = float64(num)
	output := math.Pow(10, float64(precision))
	return float64(round(newNum*output)) / output
}

func ConvertMYRtoSGD(num float32) float64 {
	return ToFixed32((num * 0.3), 0)
}

func ConvertMYRtoIDR(num float32) float64 {
	return ToFixed32((num * 3700), 0)
}
