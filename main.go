package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DarthXibalba/Conways-Game-of-Life/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const ConfigFilePath = "config.json"

type Config struct {
	GridFile   string  `json:"gridFile"`
	PixelSize  int     `json:"pixelSize"`
	TickerFreq float64 `json:"tickerFrequency"`
}

type InitialGrid struct {
	Grid [][]int `json:"grid"`
}

func ReadConfig(filePath string) (Config, InitialGrid, error) {
	var cfg Config
	var initGrid InitialGrid
	file, err := os.Open(filePath)
	if err != nil {
		return cfg, initGrid, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return cfg, initGrid, err
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return cfg, initGrid, err
	}

	// Populate grid by reading the specified grid file
	initGrid, err = ReadGridFile(cfg.GridFile)
	return cfg, initGrid, err
}

func ReadGridFile(filePath string) (InitialGrid, error) {
	var initGrid InitialGrid
	file, err := os.Open(filePath)
	if err != nil {
		return initGrid, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return initGrid, err
	}

	err = json.Unmarshal(bytes, &initGrid)
	if err != nil {
		return initGrid, err
	}
	return initGrid, nil
}

func main() {
	cfg, initGrid, err := ReadConfig(ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	thisGame, err := game.NewGame(initGrid.Grid, cfg.TickerFreq)
	if err != nil {
		log.Fatal(err)
	}

	// Window size: Each grid cell is displayed as a nxn pixel square (100 cells * n pixels each)
	fmt.Println("Screen Width:", thisGame.Width())
	fmt.Println("Screen Height:", thisGame.Height())
	ebiten.SetWindowSize(thisGame.Width()*cfg.PixelSize, thisGame.Height()*cfg.PixelSize)
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(thisGame); err != nil {
		log.Fatal(err)
	}
}
