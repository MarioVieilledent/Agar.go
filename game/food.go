package game

import "math/rand"

type Food struct {
	X float64
	Y float64
}

func NewFood(width, height int) Food {
	return Food{
		randFloats(0.0, float64(width)),
		randFloats(0.0, float64(height)),
	}
}

func randFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
