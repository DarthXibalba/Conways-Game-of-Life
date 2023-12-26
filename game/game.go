package game

import (
	"errors"
	"time"
)

var ErrorInvalidDimensions = errors.New("game dimensions cannot be zero")

// Implement ebiten.Game interface
type Game struct {
	Grid         [][]int
	width        int
	height       int
	ticker       *time.Ticker
	tickerFreq   float64
	tickerPeriod time.Duration
	lastUpdate   time.Time
}

func NewGame(grid [][]int, tickerFreq float64) (*Game, error) {
	gWidth := len(grid[0])
	gHeight := len(grid)
	if gWidth == 0 || gHeight == 0 {
		return nil, ErrorInvalidDimensions
	}

	period := time.Millisecond * time.Duration(1000.0/tickerFreq)
	g := Game{
		Grid:         grid,
		width:        gWidth,
		height:       gHeight,
		ticker:       time.NewTicker(period),
		tickerFreq:   tickerFreq,
		tickerPeriod: period,
		lastUpdate:   time.Now(),
	}
	return &g, nil
}

func (g *Game) Height() int {
	return g.height
}

func (g *Game) Width() int {
	return g.width
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(_ int, _ int) (int, int) {
	// Ignore outside width & height because we want to use a fixed sized grid
	// Grid size, each cell is a pixel
	return g.width, g.height
}
