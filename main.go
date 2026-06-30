package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite = mustLoadImage("assets/PNG/Players/Tiles/tile_0000.png")

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {
	return nil
}

// Drawing the sprites on the screen
func (g *Game) Draw(screen *ebiten.Image) {

	//animating the rotate
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWdith int, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello world")

	g := &Game{
		playerPosition: Vector{X: 100, Y: 100},
	}

	if err := ebiten.RunGame(g); err != nil {
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
