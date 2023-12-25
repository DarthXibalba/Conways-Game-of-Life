package game

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var ErrorInvalidDimensions = errors.New("game dimensions cannot be zero")

// Implement ebiten.Game interface
type Game struct {
	Grid   [][]int
	width  int
	height int
}

func NewGame(grid [][]int) (*Game, error) {
	gWidth := len(grid[0])
	gHeight := len(grid)
	if gWidth == 0 || gHeight == 0 {
		return nil, ErrorInvalidDimensions
	}

	g := Game{
		Grid:   grid,
		width:  gWidth,
		height: gHeight,
	}
	return &g, nil
}

func (g *Game) Height() int {
	return g.height
}

func (g *Game) Width() int {
	return g.width
}

// Update advances the game state, it is called with every tick/frame.
func (g *Game) Update() error {
	// TODO: Write update game logic
	return nil
}

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

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(_ int, _ int) (int, int) {
	// Ignore outside width & height because we want to use a fixed sized grid
	// Grid size, each cell is a pixel
	return g.width, g.height
}
