package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw will draw the game screen, called on every frame
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen (revert to black)
	screen.Clear()

	for y := range g.Grid {
		for x := range g.Grid[y] {
			if g.Grid[y][x] == 1 {
				screen.Set(x, y, color.White)
			}
		}
	}
}
