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
