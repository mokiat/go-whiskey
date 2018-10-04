package math

import (
	gomath "math"
)

const Pi = float32(gomath.Pi)

func Abs32(value float32) float32 {
	return float32(gomath.Abs(float64(value)))
}

func Sin32(angle float32) float32 {
	return float32(gomath.Sin(float64(angle)))
}

func Cos32(angle float32) float32 {
	return float32(gomath.Cos(float64(angle)))
}

func Sqrt32(value float32) float32 {
	return float32(gomath.Sqrt(float64(value)))
}

func Signum32(value float32) float32 {
	if gomath.Signbit(float64(value)) {
		return -1.0
	}
	return 1.0
}

func Atan32(value float32) float32 {
	return float32(gomath.Atan(float64(value)))
}
