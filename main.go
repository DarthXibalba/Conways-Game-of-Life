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
	GridFile string `json:"gridFile"`
}

type GridConfig struct {
	Grid        [][]int `json:"grid"`
	PixelSize   int     `json:"pixelSize"`
	RefreshRate float64 `json:"refreshRate"`
}

func ReadConfig(filePath string) (Config, GridConfig, error) {
	var cfg Config
	var gridConfig GridConfig
	file, err := os.Open(filePath)
	if err != nil {
		return cfg, gridConfig, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return cfg, gridConfig, err
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return cfg, gridConfig, err
	}

	// Populate grid by reading the specified grid file
	gridConfig, err = ReadGridConfig(cfg.GridFile)
	return cfg, gridConfig, err
}

func ReadGridConfig(filePath string) (GridConfig, error) {
	var gridConfig GridConfig
	file, err := os.Open(filePath)
	if err != nil {
		return gridConfig, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return gridConfig, err
	}

	err = json.Unmarshal(bytes, &gridConfig)
	if err != nil {
		return gridConfig, err
	}
	return gridConfig, nil
}

func main() {
	_, gridConfig, err := ReadConfig(ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	thisGame, err := game.NewGame(gridConfig.Grid, gridConfig.RefreshRate)
	if err != nil {
		log.Fatal(err)
	}

	// Window size: Each grid cell is displayed as a nxn pixel square (100 cells * n pixels each)
	fmt.Println("Screen Width:", thisGame.Width())
	fmt.Println("Screen Height:", thisGame.Height())
	ebiten.SetWindowSize(thisGame.Width()*gridConfig.PixelSize, thisGame.Height()*gridConfig.PixelSize)
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(thisGame); err != nil {
		log.Fatal(err)
	}
}
