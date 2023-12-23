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
	Grid      [][]int `json:"grid"`
	PixelSize int     `json:"pixelSize"`
}

func ReadConfig(filePath string) (Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	return config, err
}

func main() {
	cfg, err := ReadConfig(ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	thisGame := game.Game{
		Grid: cfg.Grid,
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
