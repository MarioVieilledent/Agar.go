package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw all food
	for _, f := range g.food {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(f.X-foodRadius-g.cell.X+float64(xCenter),
			f.Y-foodRadius-g.cell.Y+float64(yCenter))
		screen.DrawImage(images["food"], op)
	}

	// Draw player
	op := &ebiten.DrawImageOptions{}
	scl := g.cell.GetScale()
	offSet := g.cell.Diameter / 2
	op.GeoM.Scale(scl, scl)
	op.GeoM.Translate(float64(xCenter)-offSet, float64(yCenter)-offSet)
	screen.DrawImage(images["cell"], op)
}
