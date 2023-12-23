package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Implement ebiten.Game interface
type Game struct{}

// Update advances the game state, it is called with every tick/frame.
func (g *Game) Update() error {
	// TODO: Write update game logic
	return nil
}

// Draw will draw the game screen, called on every frame
func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: Write draw on screen logic
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Grid size: 100x100 (each cell is a pixel)
	return 100, 100
}

func main() {
	game := &Game{}
	// Window size: Each grid cell is displayed as a 5x5 pixel square
	ebiten.SetWindowSize(500, 500) // 100 cells * 5 pixels each
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
