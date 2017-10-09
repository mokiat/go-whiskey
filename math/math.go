package math

import (
	intmath "math"
)

const Pi = float32(intmath.Pi)

func Abs32(value float32) float32 {
	return float32(intmath.Abs(float64(value)))
}

func Sin32(angle float32) float32 {
	return float32(intmath.Sin(float64(angle)))
}

func Cos32(angle float32) float32 {
	return float32(intmath.Cos(float64(angle)))
}

func Sqrt32(value float32) float32 {
	return float32(intmath.Sqrt(float64(value)))
}

func Signum32(value float32) float32 {
	if intmath.Signbit(float64(value)) {
		return -1.0
	}
	return 1.0
}

func Atan32(value float32) float32 {
	return float32(intmath.Atan(float64(value)))
}
