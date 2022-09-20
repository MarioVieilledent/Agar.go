package game

import "github.com/hajimehoshi/ebiten/v2"

const width int = 1280
const height int = 720

const xCenter int = width / 2
const yCenter int = height / 2

const nbFood int = 500

const foodRadius float64 = 8.0

var images map[string]*ebiten.Image

type Game struct {
	cell Cell
	food []Food
}

func InitGame(imgs map[string]*ebiten.Image) (Game, int, int) {
	// Getting images
	images = imgs
	// Creation of game struct
	g := Game{
		NewCell(width, height),
		[]Food{},
	}
	// Spawn food for game start
	for i := 0; i < nbFood; i++ {
		g.food = append(g.food, NewFood(width*2, height*2))
	}
	// Start goroutine that spawn new cells randomly
	go CellSpawner(width, height)
	return g, width, height
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}
