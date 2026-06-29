package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite = mustLoadImage("assets/PNG/Players/Tiles/tile_0000.png")

type Game struct{}

func (g *Game) Update() error {
	return nil
}

// Drawing the sprites on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(PlayerSprite, nil)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWdith int, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello world")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// Sprite drawwing func

func mustLoadImage(name string) *ebiten.Image {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
