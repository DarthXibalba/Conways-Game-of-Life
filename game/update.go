package game

type PixelState int

const (
	Dead PixelState = iota
	Alive
)

type PopulationLimits int

const (
	Underpopulated PopulationLimits = 1
	Reproduction   PopulationLimits = 3
	Overpopulated  PopulationLimits = 4
)

// Up & Down Borders are connected
// Left & Right Borders are connected

// Update advances the game state, it is called with every tick/frame.
func (g *Game) Update() error {
	// Rule 1: Any live cell with 1 or no live neighbors dies (underpopulation)
	// Rule 2: Any live cell with 2-3 live neighbors lives on to the next generation (survival)
	// Rule 3: Any live cell with 4 or more neighbors dies (overpopulation)
	// Rule 4: Any dead cell with exactly three live neighbors becomes a live cell (reproduction)
	gridSearchUpdate(g)
	return nil
}

func gridSearchUpdate(g *Game) error {
	nextGrid := make([][]int, g.Height())
	for i := range nextGrid {
		nextGrid[i] = make([]int, g.Width())
	}

	for y, _ := range g.Grid {
		for x, _ := range g.Grid[y] {
			pixel := &g.Grid[y][x]
			// Get sum of neighbors
			neighbors := getNeighbors(g, y, x)
			numAlive := 0
			for _, coord := range neighbors {
				numAlive += g.Grid[coord[0]][coord[1]]
			}

			nextGrid[y][x] = int(Dead)
			if *pixel == int(Alive) {
				// Survival
				if numAlive > int(Underpopulated) && numAlive < int(Overpopulated) {
					nextGrid[y][x] = int(Alive)
				}
			} else {
				if numAlive == int(Reproduction) {
					nextGrid[y][x] = int(Alive)
				}
			}
		}
	}
	g.Grid = nextGrid
	return nil
}

func getNeighbors(g *Game, y int, x int) [][]int {
	x1 := x - 1
	if x1 < 0 {
		x1 += g.Width()
	}
	x2 := x
	x3 := x + 1
	if x3 >= g.Width() {
		x3 -= g.Width()
	}

	y1 := y - 1
	if y1 < 0 {
		y1 += g.Height()
	}
	y2 := y
	y3 := y + 1
	if y3 >= g.Height() {
		y3 -= g.Height()
	}

	return [][]int{
		[]int{y1, x1},
		[]int{y1, x2},
		[]int{y1, x3},
		[]int{y2, x1},
		//[]int{y2, x2}, exclude itself
		[]int{y2, x3},
		[]int{y3, x1},
		[]int{y3, x2},
		[]int{y3, x3},
	}
}
