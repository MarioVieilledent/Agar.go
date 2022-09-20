package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	// Eat food
	radius := g.cell.Diameter
	for i, f := range g.food {
		if f.X < g.cell.X+MaxRadius &&
			f.X > g.cell.X-MaxRadius &&
			f.Y < g.cell.Y+MaxRadius &&
			f.Y > g.cell.Y-MaxRadius {
			dX := (f.X - g.cell.X)
			dY := (f.Y - g.cell.Y)
			dist := math.Sqrt(dX*dX + dY*dY)
			if dist < radius/2 {
				if g.cell.Eat() {
					g.food[i] = g.food[len(g.food)-1]
					g.food = g.food[:len(g.food)-1]
				}
			}
		}
	}
	// Move cell
	cpx, cpy := ebiten.CursorPosition()
	theta := GetAngle(cpx, cpy)
	g.cell.Move(theta)
	return nil
}
