package game

import "math"

type Position struct {
	X float64
	Y float64
}

func GetAngle(cpx, cpy int) float64 {
	a := cpx - xCenter
	b := cpy - yCenter
	theta := math.Atan(float64(-b) / float64(a))

	// Optional?
	if theta < 0.0 {
		theta += math.Pi
	}
	// Optional?
	if b > 0.0 {
		theta += math.Pi
	}

	return theta
}
