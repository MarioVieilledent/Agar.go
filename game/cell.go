package game

import (
	"math"
	"time"
)

const DefaultSize float64 = 50.0          // Diameter of cell on start
const SizeOfImage float64 = 500           // Diameter of image of cell
const Speed float64 = 2.0                 // Speed of cell
const MaxDiameter float64 = 500.0         // Maximum diameter of cell
const MaxRadius float64 = MaxDiameter / 2 // Maximum radius of cell
const DelaySpawnNewCells int = 1000       // Delay to wait before new cell is created

type Cell struct {
	Diameter float64
	X        float64
	Y        float64
}

func NewCell(width, height int) Cell {
	offSet := DefaultSize / 2
	return Cell{DefaultSize,
		float64(width)/2.0 - offSet,
		float64(height)/2.0 - offSet,
	}
}

func (c Cell) GetScale() float64 {
	return c.Diameter / SizeOfImage
}

func (c *Cell) Move(theta float64) {
	c.X += Speed * math.Cos(theta)
	c.Y += Speed * -math.Sin(theta)
}

func (c *Cell) Eat() bool {
	if c.Diameter < MaxDiameter {
		c.Diameter++
		return true
	} else {
		return false
	}
}

func CellSpawner(width, height int) {
	for {
		time.Sleep(time.Duration(DelaySpawnNewCells) * time.Millisecond)
		NewCell(width, height)
	}
}
