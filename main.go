package main

import (
	"AgarGo/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images map[string]*ebiten.Image = map[string]*ebiten.Image{}

func init() {
	var err error
	cell, _, err := ebitenutil.NewImageFromFile("cell.png")
	if err != nil {
		log.Fatal(err)
	} else {
		images["cell"] = cell
	}
	food, _, err := ebitenutil.NewImageFromFile("food.png")
	if err != nil {
		log.Fatal(err)
	} else {
		images["food"] = food
	}
}

func main() {
	game, width, height := game.InitGame(images)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Agar.go")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
