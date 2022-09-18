package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const width int = 1280
const height int = 720

const xCenter int = width / 2
const yCenter int = height / 2

const nbFood int = 500

const foodRadius float64 = 8.0

var cell *ebiten.Image
var food *ebiten.Image

func init() {
	var err error
	cell, _, err = ebitenutil.NewImageFromFile("cell.png")
	if err != nil {
		log.Fatal(err)
	}
	food, _, err = ebitenutil.NewImageFromFile("food.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	cell Cell
	food []Food
}

func initGame() Game {
	g := Game{
		newCell(width, height),
		[]Food{},
	}
	for i := 0; i < nbFood; i++ {
		g.food = append(g.food, newFood(width*2, height*2))
	}
	return g
}

func (g *Game) Update() error {
	// Eat food
	radius := g.cell.Diameter
	for i, f := range g.food {
		if f.X < g.cell.X+maxRaius &&
			f.X > g.cell.X-maxRaius &&
			f.Y < g.cell.Y+maxRaius &&
			f.Y > g.cell.Y-maxRaius {
			dX := (f.X - g.cell.X)
			dY := (f.Y - g.cell.Y)
			dist := math.Sqrt(dX*dX + dY*dY)
			if dist < radius/2 {
				if g.cell.eat() {
					g.food[i] = g.food[len(g.food)-1]
					g.food = g.food[:len(g.food)-1]
				}
			}
		}
	}
	// Move cell
	cpx, cpy := ebiten.CursorPosition()
	theta := getAngle(cpx, cpy)
	g.cell.move(theta)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw all food
	for _, f := range g.food {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(f.X-foodRadius-g.cell.X+float64(xCenter),
			f.Y-foodRadius-g.cell.Y+float64(yCenter))
		screen.DrawImage(food, op)
	}

	// Draw player
	op := &ebiten.DrawImageOptions{}
	scl := g.cell.getScale()
	offSet := g.cell.Diameter / 2
	op.GeoM.Scale(scl, scl)
	op.GeoM.Translate(float64(xCenter)-offSet, float64(yCenter)-offSet)
	screen.DrawImage(cell, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := initGame()

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Agar.go")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func getAngle(cpx, cpy int) float64 {
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

func remove(s []Food, i int) []Food {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
