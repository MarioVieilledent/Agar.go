package main

import "math"

const defaultSize float64 = 50.0
const sizeOfImage float64 = 500
const speed float64 = 2.0
const maxDiameter float64 = 500.0
const maxRaius float64 = maxDiameter / 2

type Cell struct {
	Diameter float64
	X        float64
	Y        float64
}

func newCell(width, height int) Cell {
	offSet := defaultSize / 2
	return Cell{defaultSize,
		float64(width)/2.0 - offSet,
		float64(height)/2.0 - offSet,
	}
}

func (c Cell) getScale() float64 {
	return c.Diameter / sizeOfImage
}

func (c *Cell) move(theta float64) {
	c.X += speed * math.Cos(theta)
	c.Y += speed * -math.Sin(theta)
}

func (c *Cell) eat() bool {
	if c.Diameter < maxDiameter {
		c.Diameter++
		return true
	} else {
		return false
	}
}
